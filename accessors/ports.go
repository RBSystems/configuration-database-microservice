package accessors

import (
	"database/sql"
	"log"
)

//Port represents the port table in the database
type Port struct {
	PortID      int    `json:"portID,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

//GetAllPorts returns an array of all the port objects in the database
func (accessorGroup *AccessorGroup) GetAllPorts() ([]Port, error) {

	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("SELECT * FROM Ports")
	if err != nil {
		return []Port{}, err
	}

	defer rows.Close()

	log.Printf("Extracting data...")
	ports, err := extractPortData(rows)
	if err != nil {
		return []Port{}, err
	}

	return ports, nil
}

func extractPortData(rows *sql.Rows) ([]Port, error) {
	ports := []Port{}

	for rows.Next() {
		// var tableID sql.NullInt64
		// var tableName sql.NullString
		// var tableDescription sql.NullString
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
		// if tableID.Valid {
		// 	portID, _ = tableID.Value().(int)
		// }
		// if tableName.Valid {
		// 	portName, _ = tableName.Value().(string)
		// }
		// if tableDescription.Valid {
		// 	tableDescription, _ = tableDescription.Value().(string)
		// }

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
