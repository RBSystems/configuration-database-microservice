package accessors

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Device struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Input       bool      `json:"input"`
	Output      bool      `json:"output"`
	Building    Building  `json:"building"`
	Room        Room      `json:"room"`
	Type        string    `json:"type"`
	Power       string    `json:"power"`
	Blanked     *bool     `json:"blanked,omitempty"`
	Volume      *int      `json:"volume,omitempty"`
	Muted       *bool     `json:"muted,omitempty"`
	PowerStates []string  `json:"powerstates,omitempty"`
	Responding  bool      `json:"responding"`
	Ports       []Port    `json:"ports,omitempty"`
	Commands    []Command `json:"commands,omitempty"`
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

/* GetDevicesByQuery is a function that abstracts some of the execution and extraction
of data from the database when we're looking for responses based on the COMPLETE device struct.
The function MAY have the WHERE clause passed in to limit the devices found.
The function MAY have any JOIN clauses necessary to the WEHRE Clause not included in
the base query.
JOIN statements in the base query:
JOIN Rooms on Devices.roomID = Rooms.RoomID
JOIN Buildings on Rooms.buildingID = Buildings.buildingID
JOIN DeviceTypes on Devices.typeID = DeviceTypes.deviceTypeID
If empty string is passed in no WHERE clause will be appended, and thus all devices
will be returned.

Flow	->	Find all devices based on the clause passed in
			->	For each device found find the Ports
			->	For each device found find the Commands

Examples of valid parameters.
Example 1:
`JOIN deviceRole on deviceRole.deviceID = Devices.deviceID
JOIN DeviceRoleDefinition on DeviceRole.deviceRoleDefinitionID = DeviceRoleDefinition.deviceRoleDefinitionID
WHERE DeviceRoleDefinition.name LIKE 'AudioIn'`
Example 2:
`WHERE Devices.RoomID = 1`
*/
func (accessorGroup *AccessorGroup) GetDevicesByQuery(query string, parameters ...interface{}) ([]Device, error) {
	baseQuery := `SELECT DISTINCT Devices.deviceID,
  	Devices.Name as deviceName,
  	Devices.address as deviceAddress,
  	Devices.input,
  	Devices.output,
  	Devices.Responding,
  	Rooms.roomID,
  	Rooms.name as roomName,
  	Rooms.description as roomDescription,
  	Buildings.buildingID,
  	Buildings.name as buildingName,
  	Buildings.shortName as buildingShortname,
  	Buildings.description as buildingDescription,
  	DeviceTypes.name as deviceType,
  	PowerStates.name as power
  	FROM Devices
  	JOIN Rooms on Rooms.roomID = Devices.roomID
  	JOIN Buildings on Buildings.buildingID = Devices.buildingID
  	JOIN DeviceTypes on Devices.typeID = DeviceTypes.deviceTypeID
  	JOIN PowerStates on PowerStates.powerStateID = Devices.powerID
    JOIN DeviceRole on DeviceRole.deviceID = Devices.deviceID
    JOIN DeviceRoleDefinition on DeviceRole.deviceRoleDefinitionID = DeviceRoleDefinition.deviceRoleDefinitionID`

	allDevices := []Device{}

	rows, err := accessorGroup.Database.Query(baseQuery+" "+query, parameters...)
	if err != nil {
		return []Device{}, err
	}

	defer rows.Close()

	for rows.Next() {
		device := Device{}

		err := rows.Scan(&device.ID,
			&device.Name,
			&device.Address,
			&device.Input,
			&device.Output,
			&device.Responding,
			&device.Room.ID,
			&device.Room.Name,
			&device.Room.Description,
			&device.Building.ID,
			&device.Building.Name,
			&device.Building.Shortname,
			&device.Building.Description,
			&device.Type,
			&device.Power)
		if err != nil {
			return []Device{}, err
		}

		device.Commands, err = accessorGroup.GetDeviceCommandsByBuildingAndRoomAndName(device.Building.Shortname, device.Room.Name, device.Name)
		if err != nil {
			return []Device{}, err
		}

		device.Ports, err = accessorGroup.GetDevicePortsByBuildingAndRoomAndName(device.Building.Shortname, device.Room.Name, device.Name)
		if err != nil {
			return []Device{}, err
		}

		device.PowerStates, err = accessorGroup.GetPowerStatesByDeviceID(device.ID)
		if err != nil {
			return []Device{}, err
		}

		allDevices = append(allDevices, device)
	}

	return allDevices, nil
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

func (AccessorGroup *AccessorGroup) GetPowerStatesByDeviceID(deviceID int) ([]string, error) {
	query := `SELECT PowerStates.name FROM PowerStates
	JOIN DevicePowerStates on DevicePowerStates.powerStateID = PowerStates.powerStateID
	Where DevicePowerStates.deviceID = ?`

	toReturn := []string{}
	rows, err := AccessorGroup.Database.Query(query, deviceID)
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var value string

		err := rows.Scan(&value)
		if err != nil {
			return []string{}, err
		}
		toReturn = append(toReturn, value)
	}
	return toReturn, nil
}

func (accessorGroup *AccessorGroup) GetDevicesByBuildingAndRoomAndRole(buildingShortname string, roomName string, roleName string) ([]Device, error) {
	devices, err := accessorGroup.GetDevicesByQuery(`WHERE Rooms.name LIKE ? AND Buildings.shortname LIKE ? AND DeviceRoleDefinition.name LIKE ?`,
		roomName, buildingShortname, roleName)

	if err != nil {
		return []Device{}, err
	}
	switch strings.ToLower(roleName) {
	case "audioout":
		log.Printf("AudioOutDetected")
		devices, err = accessorGroup.GetAudioInformationForDevices(devices)
		if err != nil {
			return []Device{}, err
		}
		break
	case "videoout":
		devices, err = accessorGroup.GetDisplayInformationForDevices(devices)
		if err != nil {
			return []Device{}, err
		}
	}
	return devices, nil

}

func (accessorGroup *AccessorGroup) GetDevicesByRoleAndType(deviceRole string, deviceType string) ([]Device, error) {
	return accessorGroup.GetDevicesByQuery(`WHERE DeviceRoleDefinition.name LIKE ? AND DeviceTypes.name LIKE ?`, deviceRole, deviceType)
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

	rows, err := accessorGroup.Database.Query(`SELECT srcDevice.Name as sourceName, Ports.name as portName, destDevice.Name as DestinationDevice FROM Ports
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

	devices, err := accessorGroup.GetAudioInformationForDevices([]Device{*device})
	if err != nil {
		return Device{}, errors.New("Error in attempt to get AudioInformation")
	}

	if len(devices) == 1 {
		return devices[0], nil
	}

	return *device, nil
}

func (AccessorGroup *AccessorGroup) GetAudioInformationForDevices(devices []Device) ([]Device, error) {
	for indx := 0; indx < len(devices); indx++ {
		query := "SELECT muted, volume FROM AudioDevices where deviceID = ?"

		rows, err := AccessorGroup.Database.Query(query, devices[indx].ID)
		if err != nil {
			return []Device{}, err
		}
		log.Printf("Found some items.\n")
		for rows.Next() {
			err = rows.Scan(&devices[indx].Muted, &devices[indx].Volume)
			if err != nil {
				return []Device{}, err
			}
		}
	}
	return devices, nil
}

func (AccessorGroup *AccessorGroup) GetDisplayInformationForDevices(devices []Device) ([]Device, error) {
	for indx := 0; indx < len(devices); indx++ {
		query := "SELECT blanked FROM Displays where deviceID = ?"

		rows, err := AccessorGroup.Database.Query(query, devices[indx].ID)
		if err != nil {
			return []Device{}, err
		}
		log.Printf("Found some items.\n")
		for rows.Next() {
			err = rows.Scan(&devices[indx].Blanked)
			if err != nil {
				return []Device{}, err
			}
		}
	}
	return devices, nil
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

func (accessorGroup *AccessorGroup) PutDeviceAttributeByDeviceAndRoomAndBuilding(building string, room string, device string, attribute string, attributeValue string) (Device, error) {
	switch strings.ToLower(attribute) {
	case "volume":
		statement := `update AudioDevices SET volume = ? WHERE deviceID =
			(Select deviceID from Devices
				JOIN Rooms on Rooms.roomID = Devices.roomID
				JOIN Buildings on Buildings.buildingID = Rooms.buildingID
				WHERE Devices.name LIKE ? AND Rooms.name LIKE ? AND Buildings.shortName LIKE ?)`
		val, err := strconv.Atoi(attributeValue)
		if err != nil {
			return Device{}, err
		}

		_, err = accessorGroup.Database.Exec(statement, val, device, room, building)
		if err != nil {
			return Device{}, err
		}
		break

	case "muted":
		statement := `udpate AudioDevices SET muted = ? WHERE deviceID =
			(Select deviceID from Devices
				JOIN Rooms on Rooms.RoomID = Devices.RoomID
				JOIN Buildings on Buildings.RoomID = Buildings.RoomID
				WHERE Devices.name LIKE ? AND Rooms.name LIKE ? AND Buildings.shortName LIKE ?)`
		_, err := accessorGroup.Database.Exec(statement, attributeValue, device, room, building)
		if err != nil {
			return Device{}, err
		}
		break
	}

	dev, err := accessorGroup.GetDeviceByBuildingAndRoomAndName(building, room, device)
	return dev, err
}
