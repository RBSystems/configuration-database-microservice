package accessors

import (
	"database/sql"
)

//DeviceType corresponds to the DeviceType table in the database
type DeviceClass struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	DisplayName string `json:"display-name"`
	Description string `json:"description"`
}

//GetDeviceClasses returns a dump of the table in the database
func (accessorGroup *AccessorGroup) GetDeviceTypes() ([]DeviceClass, error) {

	var toReturn []DeviceClass
	rows, err := accessorGroup.Database.Query("Select typeName, typeDescription, typeDisplayName From Device Types")
	if err != nil {
		return toReturn, err
	}

	toReturn, err = extractDeviceClassData(rows)

	return toReturn, err
}

func (accessorGroup *AccessorGroup) SetDeviceTypeByName(name string, device Device) error {

}

func extractDeviceClassData(rows *sql.Rows) ([]DeviceClass, error) {

	toReturn := []DeviceClass{}
	var id *int
	var name *string
	var displayName *string
	var description *string

	for rows.Next() {
		curVal := DeviceClass{}

		err := rows.Scan(&id, &name, &displayName, &description)
		if err != nil {
			return toReturn, err
		}

		if id != nil {
			curVal.ID = *id
		}

		if name != nil {
			curVal.Name = *name
		}

		if displayName != nil {
			curVal.DisplayName = *displayName
		}

		if description != nil {
			curVal.Description = *description
		}

		toReturn = append(toReturn, curVal)
	}
	rows.Close()

	return toReturn, nil
}
