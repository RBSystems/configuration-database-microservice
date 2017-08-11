package accessors

import "github.com/byuoitav/configuration-database-microservice/structs"

func (accessorGroup *AccessorGroup) AddDeviceCommand(dc structs.DeviceCommand) (structs.DeviceCommand, error) {
	// devicecommand.ID needs to be changed to devicecommand.Command.ID, but Command doesn't have that field yet
	result, err := accessorGroup.Database.Exec("Insert into DeviceCommands (deviceCommandID, deviceID, commandID, microserviceID, endpointID, enabled) VALUES(?,?,?,?,?,?)", dc.ID, dc.DeviceID, dc.CommandID, dc.MicroserviceID, dc.EndpointID, dc.Enabled)

	if err != nil {
		return structs.DeviceCommand{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.DeviceCommand{}, err
	}

	dc.ID = int(id)
	return dc, nil
}
