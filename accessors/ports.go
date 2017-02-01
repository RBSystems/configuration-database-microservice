package accessors

import "github.com/byuoitav/configuration-database-microservice/dbstructs"

//GetAllPorts returns an array of all the port objects in the database
func (accessorGroup *AccessorGroup) GetAllPorts() ([]dbstructs.Port, error) {
	ports := []dbstructs.Port{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Ports")
	if err != nil {
		return ports, err
	}

	//why do rows need to be closed?
	defer rows.Close()

	for rows.Next() {
		port := dbstructs.Port{}

		err = rows.Scan(&port.PortID, &port.Name, &port.Description)
		if err != nil {
			return []dbstructs.Port{}, err
		}

		ports = append(ports, port)
	}

	err = rows.Err()
	if err != nil {
		return []dbstructs.Port{}, err
	}

	return ports, nil
}
