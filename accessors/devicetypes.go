package accessors

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

//GetDeviceClasses returns a dump of the table in the database
func (accessorGroup *AccessorGroup) GetDeviceTypes() ([]structs.DeviceClass, error) {

	var toReturn []structs.DeviceClass
	rows, err := accessorGroup.Database.Query("Select deviceTypeID, typeName, typeDescription, typeDisplayName From DeviceTypes")
	if err != nil {
		return toReturn, err
	}

	toReturn, err = extractDeviceClassData(rows)

	return toReturn, err
}

func (accessorGroup *AccessorGroup) SetDeviceTypeByID(id int, deviceID int) error {
	log.Printf("Updating type id of device %v to %v", deviceID, id)

	query := "UPDATE devices SET typeID = ? WHERE deviceID = ?"

	res, err := accessorGroup.Database.Exec(query, id, deviceID)
	if err != nil {
		return err
	}

	if num, err := res.RowsAffected(); num != 1 || err != nil {
		if err != nil {
			return err
		}

		err = errors.New(fmt.Sprintf("There was a problem updating the device type: incorrect number of rows affected: %v. ", res.RowsAffected))
		return err
	}

	log.Printf("Done.")
	return nil
}

func extractDeviceClassData(rows *sql.Rows) ([]structs.DeviceClass, error) {

	toReturn := []structs.DeviceClass{}
	var id *int
	var name *string
	var displayName *string
	var description *string

	for rows.Next() {
		curVal := structs.DeviceClass{}

		err := rows.Scan(&id, &name, &description, &displayName)
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

func (accessorGroup *AccessorGroup) GetDeviceClassByName(name string) (structs.DeviceClass, error) {
	row, err := accessorGroup.Database.Query("Select deviceTypeID, typeName, typeDescription, typeDisplayName From DeviceTypes WHERE typeName = ?", name)
	if err != nil {
		return structs.DeviceClass{}, err
	}
	defer row.Close()

	dt, err := extractDeviceClassData(row)
	if err != nil || len(dt) < 1 {
		if len(dt) < 1 {
			return structs.DeviceClass{}, errors.New("No device types found")
		}
		return structs.DeviceClass{}, err
	}

	return dt[0], nil
}
