package accessors

import (
	"database/sql"
)

type DeviceCommand struct {
	ID           int          `json:"id,omitempty"`
	Device       Device       `json:"device"`
	Command      Command      `json:"command"`
	Microservice Microservice `json:"microservice"`
	Endpoint     Endpoint     `json:"endpoint"`
	Enabled      int          `json:"enabled"`
}

func (accessorGroup *AccessorGroup) GetDeviceCommands() ([]DeviceCommand, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceCommands")
	if err != nil {
		return []DeviceCommand{}, err
	}

	devicecommands, err := extractDeviceCommands(rows)
	if err != nil {
		return []DeviceCommand{}, err
	}
	defer rows.Close()

	return devicecommands, nil
}

func (accessorGroup *AccessorGroup) AddDeviceCommand(devicecommand DeviceCommand) (DeviceCommand, error) {
	// devicecommand.ID needs to be changed to devicecommand.Command.ID, but Command doesn't have that field yet
	result, err := accessorGroup.Database.Exec("Insert into DeviceCommands (devicecommandID, deviceID, commandID, microserviceID, endpointID, enabled) VALUES(?,?,?,?,?,?)", devicecommand.ID, devicecommand.Device.ID, devicecommand.ID, devicecommand.Microservice.ID, devicecommand.Endpoint.ID, devicecommand.Enabled)
	if err != nil {
		return DeviceCommand{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return DeviceCommand{}, err
	}

	devicecommand.ID = int(id)
	return devicecommand, nil
}

func extractDeviceCommands(rows *sql.Rows) ([]DeviceCommand, error) {
	var devicecommands []DeviceCommand
	var devicecommand DeviceCommand
	var id *int
	var dID *int
	var cID *int
	var mID *int
	var eID *int
	var enabled *int

	for rows.Next() {
		err := rows.Scan(&id, &dID, &cID, &mID, &eID, &enabled)
		if err != nil {
			return []DeviceCommand{}, err
		}
		if id != nil {
			devicecommand.ID = *id
		}
		if dID != nil {
			devicecommand.Device.ID = *dID
		}
		if cID != nil {
			devicecommand.ID = *cID // also needs to be changed to devicecommand.Command.ID
		}
		if mID != nil {
			devicecommand.Microservice.ID = *mID
		}
		if eID != nil {
			devicecommand.Endpoint.ID = *eID
		}
		if enabled != nil {
			devicecommand.Enabled = *enabled
		}

		devicecommands = append(devicecommands, devicecommand)
	}
	return devicecommands, nil
}
