package accessors

import (
	"database/sql"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

//GetDeviceClasses returns a dump of the table in the database
func (accessorGroup *AccessorGroup) GetDeviceClasses() ([]structs.DeviceType, error) {

	var DeviceClasses []structs.DeviceType

	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceClasses")
	if err != nil {
		return []structs.DeviceType{}, err
	}

	DeviceClasses, err = extractDeviceTypeData(rows)
	if err != nil {
		return []structs.DeviceType{}, err
	}
	defer rows.Close()

	return DeviceClasses, nil
}

func (accessorGroup *AccessorGroup) AddDeviceType(deviceType structs.DeviceType) (structs.DeviceType, error) {
	result, err := accessorGroup.Database.Exec("Insert into DeviceClasses (deviceClassID, name, description) VALUES(?,?,?)", deviceType.ID, deviceType.Name, deviceType.Description)
	if err != nil {
		return structs.DeviceType{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.DeviceType{}, err
	}

	deviceType.ID = int(id)
	return deviceType, nil
}

func (accessorGroup *AccessorGroup) GetDeviceTypeByID(id int) (structs.DeviceType, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM DeviceClasses WHERE deviceClassID = ?", id)

	dt, err := extractDeviceType(row)
	if err != nil {
		return structs.DeviceType{}, err
	}

	return dt, nil
}

func (accessorGroup *AccessorGroup) GetDeviceTypeByName(name string) (structs.DeviceType, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM DeviceClasses WHERE name = ?", name)

	dt, err := extractDeviceType(row)
	if err != nil {
		return structs.DeviceType{}, err
	}

	return dt, nil
}

func extractDeviceTypeData(rows *sql.Rows) ([]structs.DeviceType, error) {

	var deviceTypes []structs.DeviceType
	var deviceType structs.DeviceType
	var id *int
	var name *string
	var description *string

	for rows.Next() {

		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return []structs.DeviceType{}, err
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

func extractDeviceType(row *sql.Row) (structs.DeviceType, error) {
	var dt structs.DeviceType
	var id *int
	var name *string
	var description *string

	err := row.Scan(&id, &name, &description)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return structs.DeviceType{}, err
	}
	if id != nil {
		dt.ID = *id
	}
	if name != nil {
		dt.Name = *name
	}
	if description != nil {
		dt.Description = *name
	}

	return dt, nil
}
