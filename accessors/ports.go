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

	query := `SELECT * FROM Ports`
	log.Printf("Querying: \"%v\"", query)
	rows, err := accessorGroup.Database.Query(query)
	if err != nil {
		return []Port{}, err
	}

	defer rows.Close()

	ports, err := extractPortData(rows)
	if err != nil {
		return []Port{}, err
	}

	log.Printf("Done.")
	return ports, nil
}

func extractPortData(rows *sql.Rows) ([]Port, error) {

	log.Printf("Extracting data...")

	var ports []Port

	for rows.Next() {

		var port Port
		var tableID *int
		var tableName *string
		var tableDescription *string

		err := rows.Scan(&tableID, &tableName, &tableDescription)
		if err != nil {
			return []Port{}, err
		}

		log.Printf("Creating struct...")
		if tableID != nil {
			port.PortID = *tableID
		}
		if tableName != nil {
			port.Name = *tableName
		}
		if tableDescription != nil {
			port.Description = *tableDescription
		}

		ports = append(ports, port)
	}

	err := rows.Err()
	if err != nil {
		return []Port{}, err
	}

	return ports, nil
}
