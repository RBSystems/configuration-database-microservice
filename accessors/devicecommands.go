package accessors

type DeviceCommand struct {
	ID           int          `json:"id,omitempty"`
	Device       Device       `json:"device"`
	Command      RawCommand   `json:"command"`
	Microservice Microservice `json:"microservice"`
	Endpoint     Endpoint     `json:"endpoint"`
	Enabled      int          `json:"enabled"`
}

func (accessorGroup *AccessorGroup) AddDeviceCommand(dc DeviceCommand) (DeviceCommand, error) {
	// devicecommand.ID needs to be changed to devicecommand.Command.ID, but Command doesn't have that field yet
	result, err := accessorGroup.Database.Exec("Insert into DeviceCommands (deviceCommandID, deviceID, commandID, microserviceID, endpointID, enabled) VALUES(?,?,?,?,?,?)", dc.ID, dc.Device.ID, dc.Command.ID, dc.Microservice.ID, dc.Endpoint.ID, dc.Enabled)

	if err != nil {
		return DeviceCommand{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return DeviceCommand{}, err
	}

	dc.ID = int(id)
	return dc, nil
}
