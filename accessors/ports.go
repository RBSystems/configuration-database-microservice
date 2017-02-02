package accessors

import "database/sql"

//Port represents the port table in the database
type Port struct {
	PortID      int    `json:"portID,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

//GetAllPorts returns an array of all the port objects in the database
func (accessorGroup *AccessorGroup) GetAllPorts() ([]Port, error) {

	rows, err := accessorGroup.Database.Query("SELECT * FROM Ports")
	if err != nil {
		return []Port{}, err
	}

	//why do rows need to be closed?
	defer rows.Close()

	ports, err := accessorGroup.ExtractPortData(rows)
	if err != nil {
		return []Port{}, err
	}

	return ports, nil
}

//ExtractPortData performs the scan of the rows in an SQL table
func (accessorGroup *AccessorGroup) ExtractPortData(rows *sql.Rows) ([]Port, error) {
	ports := []Port{}

	for rows.Next() {
		port := Port{}

		err := rows.Scan(&port.PortID, &port.Name, &port.Description)
		if err != nil {
			return []Port{}, err
		}

		ports = append(ports, port)
	}

	err := rows.Err()
	if err != nil {
		return []Port{}, err
	}

	return ports, nil
}
