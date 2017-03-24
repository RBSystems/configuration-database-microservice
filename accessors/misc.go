package accessors

type DeviceType struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//GetDeviceTypes returns a dump of the table in the database
func (accessorGroup *AccessorGroup) GetDeviceTypes() ([]DeviceType, error) {

	var DeviceTypes []DeviceType

	rows, err := accesssorGroup.Database.Query("SELECT * FROM DeviceTypes")
	if err != nil {
		return []DeviceType{}, err
	}

	DeviceTypes, err = extractDeviceTypeData(rows)
	if err != nil {
		return []DeviceType{}, err
	}

	return DeviceTypes, nil
}

func extractDeviceTypeData([]DeviceType, error) {

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
			deviceType.id = *id
		}
		if name != nil {
			deviceType.name = *name
		}
		if description != nil {
			deviceType.description = *description
		}

		deviceTypes = append(deviceType)
	}

	return deviceTypes, nil
}
