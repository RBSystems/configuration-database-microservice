package accessors

import (
	"database/sql"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

func (accessorGroup *AccessorGroup) GetDevicePowerStates() ([]structs.DevicePowerState, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM DevicePowerStates")
	if err != nil {
		return []structs.DevicePowerState{}, err
	}

	devicepowerstates, err := exctractDevicePowerStateData(rows)
	if err != nil {
		return []structs.DevicePowerState{}, err
	}
	defer rows.Close()

	return devicepowerstates, nil
}

func (accessorGroup *AccessorGroup) AddDevicePowerState(dps structs.DevicePowerState) (structs.DevicePowerState, error) {
	response, err := accessorGroup.Database.Exec("INSERT INTO DevicePowerStates (devicePowerStateID, deviceID, powerStateID) VALUES(?,?,?)", dps.ID, dps.DeviceID, dps.PowerStateID)
	if err != nil {
		return structs.DevicePowerState{}, err
	}

	id, err := response.LastInsertId()
	dps.ID = int(id)

	return dps, nil
}

func exctractDevicePowerStateData(rows *sql.Rows) ([]structs.DevicePowerState, error) {

	var devicepowerstates []structs.DevicePowerState
	var devicepowerstate structs.DevicePowerState
	var id *int
	var dID *int
	var pID *int

	for rows.Next() {
		err := rows.Scan(&id, &dID, &pID)
		if err != nil {
			return []structs.DevicePowerState{}, err
		}

		if id != nil {
			devicepowerstate.ID = *id
		}
		if dID != nil {
			devicepowerstate.DeviceID = *dID
		}
		if pID != nil {
			devicepowerstate.PowerStateID = *pID
		}

		devicepowerstates = append(devicepowerstates, devicepowerstate)
	}

	err := rows.Err()
	if err != nil {
		return []structs.DevicePowerState{}, err
	}

	return devicepowerstates, nil
}
