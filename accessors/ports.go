package accessors

import "database/sql"

//PortType corresponds to the Ports table in the Database and really should be called Port
//TODO:Change struct name to "Port"
type PortType struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (accessorGroup *AccessorGroup) GetAllPorts() ([]PortType, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Ports")
	if err != nil {
		return []PortType{}, err
	}

	defer rows.Close()
	allPorts, err := extractPortData(rows)

	return allPorts, nil
}

//AddPort adds an entry to the Ports table in the database
func (accessorGroup *AccessorGroup) AddPort(portToAdd PortType) (PortType, error) {

	result, err := accessorGroup.Database.Exec("INSERT into Ports (portID, name, description) VALUES(?,?,?)", portToAdd.ID, portToAdd.Name, portToAdd.Description)
	if err != nil {
		return PortType{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return PortType{}, err
	}

	portToAdd.ID = int(id)
	return portToAdd, nil
}

func extractPortData(rows *sql.Rows) ([]PortType, error) {

	var allPorts []PortType
	var port PortType
	var id *int
	var name *string
	var description *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return []PortType{}, err
		}

		if id != nil {
			port.ID = *id
		}
		if name != nil {
			port.Name = *name
		}
		if description != nil {
			port.Description = *description
		}

		allPorts = append(allPorts, port)
	}

	err := rows.Err()
	if err != nil {
		return []PortType{}, err
	}

	return allPorts, nil
}
