package accessors

import (
	"database/sql"
	"log"
)

//GetAllPorts returns an array of all the port objects in the database
func (accessorGroup *AccessorGroup) GetAllPorts() ([]Port, error) {

	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("SELECT * FROM Ports")
	if err != nil {
		return []Port{}, err
	}

	defer rows.Close()

	ports, err := extractPortData(rows)
	if err != nil {
		return []Port{}, err
	}

	return ports, nil
}

func extractPortData(rows *sql.Rows) ([]Port, error) {

	log.Printf("Extracting data...")

	ports := []Port{}

	for rows.Next() {

		var tableID *int
		var tableName *string
		var tableDescription *string

		var portID int
		var portName string
		var portDescription string

		err := rows.Scan(&tableID, &tableName, &tableDescription)
		if err != nil {
			return []Port{}, err
		}

		if tableID != nil {
			portID = *tableID
		}
		if tableName != nil {
			portName = *tableName
		}
		if tableDescription != nil {
			portDescription = *tableDescription
		}

		log.Printf("Creating Port struct...")
		port := Port{
			portID,
			portName,
			portDescription,
		}

		ports = append(ports, port)
	}

	err := rows.Err()
	if err != nil {
		return []Port{}, err
	}

	return ports, nil
}
