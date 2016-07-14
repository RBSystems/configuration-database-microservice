package accessors

import "errors"

type Device struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Room     int    `json:"room"`
	Protocol string `json:"protocol"`
}

// GetAllDevices returns a list of devices from the database
func (accessorGroup *AccessorGroup) GetAllDevices() ([]Device, error) {
	allDevices := []Device{}

	rows, err := accessorGroup.Database.Query("SELECT devices.id, devices.name, devices.address, rooms.name, devices.protocol FROM devices JOIN rooms ON devices.room=room.ID")
	if err != nil {
		return []Device{}, err
	}

	defer rows.Close()

	for rows.Next() {
		device := Device{}

		err := rows.Scan(&device.ID, &device.Name, &device.Address, &device.Room, &device.Protocol)
		if err != nil {
			return []Device{}, err
		}

		allDevices = append(allDevices, device)
	}

	err = rows.Err()
	if err != nil {
		return []Device{}, err
	}

	return allDevices, nil
}

func (accessorGroup *AccessorGroup) GetDeviceByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) (Device, error) {
	room, err := accessorGroup.GetRoomByBuildingAndName(buildingShortname, roomName)
	if err != nil {
		return Device{}, errors.New("Could not find a room named \"" + roomName + "\" in a building named \"" + buildingShortname + "\"")
	}

	device := &Device{}
	err = accessorGroup.Database.QueryRow("SELECT * FROM devices WHERE building=? AND room=?", room.Building, room.ID).Scan(&device.ID, &device.Name, &device.Address, &device.Room, &device.Protocol)
	if err != nil {
		return Device{}, err
	}

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
