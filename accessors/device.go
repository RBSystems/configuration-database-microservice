package accessors

import (
	"errors"
	"log"
)

type Device struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Input      bool      `json:"input"`
	Output     bool      `json:"output"`
	Building   Building  `json:"building"`
	Room       Room      `json:"room"`
	Type       int       `json:"type"`
	Power      int       `json:"power"`
	Responding bool      `json:"responding"`
	Ports      []Port    `json:"ports,omitempty"`
	Commands   []Command `json:"commands,omitempty"`
}

type Port struct {
	Source      string `json:"source"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
}

type Command struct {
	Name         string   `json:"name"`
	Endpoint     Endpoint `json:"endpoint"`
	Microservice string   `json:"microservice"`
}

type Endpoint struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type DeviceRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Building string `json:"building"`
	Room     string `json:"room"`
}

// GetAllDevices returns a list of devices from the database
func (accessorGroup *AccessorGroup) GetAllDevices() ([]Device, error) {
	allBuildings := []Building{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Devices")
	if err != nil {
		return []Device{}, err
	}

	for rows.Next() {
		building := Building{}

		err := rows.Scan(&building.ID, &building.Name, &building.Shortname)
		if err != nil {
			return []Device{}, err
		}

		allBuildings = append(allBuildings, building)
	}

	allRooms := []RoomRequest{}

	rows, err = accessorGroup.Database.Query("SELECT rooms.id, rooms.name, rooms.vlan FROM rooms")
	if err != nil {
		return []Device{}, err
	}

	defer rows.Close()

	for rows.Next() {
		room := RoomRequest{}

		err := rows.Scan(&room.ID, &room.Name, &room.VLAN)
		if err != nil {
			return []Device{}, err
		}

		allRooms = append(allRooms, room)
	}

	err = rows.Err()
	if err != nil {
		return []Device{}, err
	}

	allDevices := []Device{}

	rows, err = accessorGroup.Database.Query("SELECT * FROM devices")
	if err != nil {
		return []Device{}, err
	}

	defer rows.Close()

	for rows.Next() {
		device := Device{}

		err := rows.Scan(&device.ID, &device.Name, &device.Address, &device.Room.ID, &device.Building.ID)
		if err != nil {
			return []Device{}, err
		}

		for i := 0; i < len(allBuildings); i++ {
			if allBuildings[i].ID == device.Building.ID {
				device.Building = allBuildings[i]
				break
			}
		}

		for i := 0; i < len(allRooms); i++ {
			if allRooms[i].ID == device.Room.ID {
				// device.Room = allRooms[i]
				break
			}
		}

		allDevices = append(allDevices, device)
	}

	err = rows.Err()
	if err != nil {
		return []Device{}, err
	}

	return allDevices, nil
}

func (accessorGroup *AccessorGroup) GetDevicesByBuildingAndRoom(buildingShortname string, roomName string) ([]Device, error) {
	room, err := accessorGroup.GetRoomByBuildingAndName(buildingShortname, roomName)
	if err != nil {
		return []Device{}, errors.New("Could not find a room named \"" + roomName + "\" in a building named \"" + buildingShortname + "\"")
	}

	allDevices := []Device{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Devices WHERE buildingID=? AND roomID=?", room.Building.ID, room.ID)
	if err != nil {
		return []Device{}, err
	}

	for rows.Next() {
		device := Device{}

		err := rows.Scan(&device.ID, &device.Name, &device.Address, &device.Input, &device.Output, &device.Building.ID, &device.Room.ID, &device.Type, &device.Power, &device.Responding)
		if err != nil {
			return []Device{}, err
		}

		allDevices = append(allDevices, device)
	}

	return allDevices, nil
}

func (accessorGroup *AccessorGroup) GetDeviceCommandsByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) ([]Command, error) {
	allCommands := []Command{}

	rows, err := accessorGroup.Database.Query(`SELECT Commands.name as commandName, Endpoints.name as endpointName, Endpoints.path as endpointPath, Microservices.address as microserviceAddress
    FROM Devices
    JOIN DeviceCommands on Devices.deviceID = DeviceCommands.deviceID JOIN Commands on DeviceCommands.commandID = Commands.commandID JOIN Endpoints on DeviceCommands.endpointID = Endpoints.endpointID JOIN Microservices ON DeviceCommands.microserviceID = Microservices.microserviceID
    JOIN Rooms ON Rooms.roomID=Devices.roomID
    JOIN Buildings ON Rooms.buildingID=Buildings.buildingID
    WHERE Rooms.name=? AND Buildings.shortName=? AND Devices.name=?`, roomName, buildingShortname, deviceName)
	if err != nil {
		return []Command{}, err
	}

	for rows.Next() {
		command := Command{}

		err := rows.Scan(&command.Name, &command.Endpoint.Name, &command.Endpoint.Path, &command.Microservice)
		if err != nil {
			return []Command{}, err
		}

		allCommands = append(allCommands, command)
	}

	return allCommands, nil
}

func (accessorGroup *AccessorGroup) GetDevicePortsByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) ([]Port, error) {
	allPorts := []Port{}

	rows, err := accessorGroup.Database.Query(`SELECT srcDevice.Name as sourceName, Ports.Port as portName, destDevice.Name as DestinationDevice FROM Ports
    JOIN PortConfiguration ON Ports.PortID = PortConfiguration.PortID
    JOIN Devices as srcDevice on srcDevice.DeviceID = PortConfiguration.sourceDeviceID
    JOIN Devices as destDevice on destDevice.DeviceID = PortConfiguration.destinationDeviceID
    JOIN Rooms ON Rooms.roomID=destDevice.roomID
    JOIN Buildings ON Rooms.buildingID=Buildings.buildingID
    WHERE Rooms.name=? AND Buildings.shortName=? AND destDevice.name=?`, roomName, buildingShortname, deviceName)
	if err != nil {
		log.Print(err)
		return []Port{}, err
	}

	for rows.Next() {
		port := Port{}

		err := rows.Scan(&port.Source, &port.Name, &port.Destination)
		if err != nil {
			log.Print(err)
			return []Port{}, err
		}

		allPorts = append(allPorts, port)
	}

	return allPorts, nil
}

func (accessorGroup *AccessorGroup) GetDeviceByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) (Device, error) {
	room, err := accessorGroup.GetRoomByBuildingAndName(buildingShortname, roomName)
	if err != nil {
		return Device{}, errors.New("Could not find a room named \"" + roomName + "\" in a building named \"" + buildingShortname + "\"")
	}

	device := &Device{}
	err = accessorGroup.Database.QueryRow("SELECT * FROM Devices WHERE buildingID=? AND roomID=? AND name=?", room.Building.ID, room.ID, deviceName).Scan(&device.ID, &device.Name, &device.Address, &device.Input, &device.Output, &device.Building.ID, &device.Room.ID, &device.Type, &device.Power, &device.Responding)
	if err != nil {
		return Device{}, err
	}

	commands, err := accessorGroup.GetDeviceCommandsByBuildingAndRoomAndName(buildingShortname, roomName, deviceName)
	if err != nil {
		return Device{}, errors.New("Could not find a device named \"" + deviceName + "\" in a room named \"" + roomName + "\" in a building named \"" + buildingShortname + "\"")
	}

	device.Commands = commands

	ports, err := accessorGroup.GetDevicePortsByBuildingAndRoomAndName(buildingShortname, roomName, deviceName)
	if err != nil {
		return Device{}, errors.New("Poots Could not find a device named \"" + deviceName + "\" in a room named \"" + roomName + "\" in a building named \"" + buildingShortname + "\"")
	}

	device.Ports = ports

	return *device, nil
}

// MakeDevice adds a device to the database
func (accessorGroup *AccessorGroup) MakeDevice(name string, address string, buildingShortname string, roomName string, protocol string) (Device, error) {
	room, err := accessorGroup.GetRoomByBuildingAndName(buildingShortname, roomName)
	if err != nil {
		return Device{}, errors.New("Could not find a room named \"" + roomName + "\" in a building named \"" + buildingShortname + "\"")
	}

	_, err = accessorGroup.Database.Exec("INSERT INTO devices (name, address, room, protocol) VALUES (?, ?, ?, ?)", name, address, room.ID, protocol)
	if err != nil {
		return Device{}, err
	}

	device, err := accessorGroup.GetDeviceByBuildingAndRoomAndName(buildingShortname, roomName, name)
	if err != nil {
		return Device{}, err
	}

	return device, nil
}
