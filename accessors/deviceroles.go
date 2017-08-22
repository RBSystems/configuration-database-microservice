package accessors

import (
	"database/sql"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

func (accessorGroup *AccessorGroup) GetDeviceRoles() ([]structs.DeviceRole, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceRole")
	if err != nil {
		return []structs.DeviceRole{}, err
	}

	deviceroles, err := exctractDeviceRoleData(rows)
	if err != nil {
		return []structs.DeviceRole{}, err
	}
	defer rows.Close()

	return deviceroles, nil
}

func (accessorGroup *AccessorGroup) AddDeviceRole(dr structs.DeviceRole) (structs.DeviceRole, error) {
	response, err := accessorGroup.Database.Exec("INSERT INTO DeviceRole (deviceRoleID, deviceID, deviceRoleDefinitionID) VALUES(?,?,?)", dr.ID, dr.DeviceID, dr.DeviceRoleDefinitionID)
	if err != nil {
		return structs.DeviceRole{}, err
	}

	id, err := response.LastInsertId()
	dr.ID = int(id)

	return dr, nil
}

func exctractDeviceRoleData(rows *sql.Rows) ([]structs.DeviceRole, error) {
	var deviceroles []structs.DeviceRole
	var devicerole structs.DeviceRole
	var id *int
	var dID *int
	var rID *int

	for rows.Next() {
		err := rows.Scan(&id, &dID, &rID)
		if err != nil {
			return []structs.DeviceRole{}, err
		}

		if id != nil {
			devicerole.ID = *id
		}
		if dID != nil {
			devicerole.DeviceID = *dID
		}
		if rID != nil {
			devicerole.DeviceRoleDefinitionID = *rID
		}
		deviceroles = append(deviceroles, devicerole)
	}

	err := rows.Err()
	if err != nil {
		return []structs.DeviceRole{}, err
	}

	return deviceroles, nil
}
