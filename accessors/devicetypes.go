package accessors

import "database/sql"

//DeviceType corresponds to the DeviceType table in the database
type DeviceType struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//GetDeviceTypes returns a dump of the table in the database
func (accessorGroup *AccessorGroup) GetDeviceTypes() ([]DeviceType, error) {

	var DeviceTypes []DeviceType

	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceTypes")
	if err != nil {
		return []DeviceType{}, err
	}

	DeviceTypes, err = extractDeviceTypeData(rows)
	if err != nil {
		return []DeviceType{}, err
	}

	return DeviceTypes, nil
}

func (accessorGroup *AccessorGroup) AddDeviceType(deviceType DeviceType) (DeviceType, error) {
	result, err := accessorGroup.Database.Exec("Insert into DeviceTypes (deviceTypeID, name, description) VALUES(?,?,?)", deviceType.ID, deviceType.Name, deviceType.Description)
	if err != nil {
		return DeviceType{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return DeviceType{}, err
	}

	deviceType.ID = int(id)
	return deviceType, nil
}

func extractDeviceTypeData(rows *sql.Rows) ([]DeviceType, error) {

	var deviceTypes []DeviceType
	var deviceType DeviceType
	var id *int
	var name *string
	var description *string

	for rows.Next() {

		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return []DeviceType{}, err
		}

		if id != nil {
			deviceType.ID = *id
		}
		if name != nil {
			deviceType.Name = *name
		}
		if description != nil {
			deviceType.Description = *description
		}

		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes, nil
}
