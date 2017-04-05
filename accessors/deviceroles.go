package accessors

import "database/sql"

type DeviceRole struct {
	ID     int           `json:"id,omitempty"`
	Device Device        `json:"device"`
	Role   DeviceRoleDef `json:"role"`
}

func (accessorGroup *AccessorGroup) GetDeviceRoles() ([]DeviceRole, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceRole")
	if err != nil {
		return []DeviceRole{}, err
	}

	deviceroles, err := exctractDeviceRoleData(rows)
	if err != nil {
		return []DeviceRole{}, err
	}
	defer rows.Close()

	return deviceroles, nil
}

func (accessorGroup *AccessorGroup) AddDeviceRole(dr DeviceRole) (DeviceRole, error) {
	response, err := accessorGroup.Database.Exec("INSERT INTO DeviceRole (deviceRoleID, deviceID, deviceRoleDefinitionID) VALUES(?,?,?)", dr.ID, dr.Device.ID, dr.Role.ID)
	if err != nil {
		return DeviceRole{}, err
	}

	id, err := response.LastInsertId()
	dr.ID = int(id)

	return dr, nil
}

func exctractDeviceRoleData(rows *sql.Rows) ([]DeviceRole, error) {
	var deviceroles []DeviceRole
	var devicerole DeviceRole
	var id *int
	var dID *int
	var rID *int

	for rows.Next() {
		err := rows.Scan(&id, &dID, &rID)
		if err != nil {
			return []DeviceRole{}, err
		}

		if id != nil {
			devicerole.ID = *id
		}
		if dID != nil {
			devicerole.Device.ID = *dID
		}
		if rID != nil {
			devicerole.Role.ID = *rID
		}
		deviceroles = append(deviceroles, devicerole)
	}

	err := rows.Err()
	if err != nil {
		return []DeviceRole{}, err
	}

	return deviceroles, nil
}
