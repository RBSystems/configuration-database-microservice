package accessors

import "database/sql"

func (accessorGroup *AccessorGroup) GetAllEndpoints() ([]Endpoint, error) {

	rows, err := accessorGroup.Database.Query("SELECT * FROM Endpoints")
	if err != nil {
		return []Endpoint{}, err
	}
	defer rows.Close()

	endpoints, err := exctractEndpointData(rows)
	if err != nil {
		return []Endpoint{}, err
	}

	return endpoints, nil
}

func (accessorGroup *AccessorGroup) AddEndpoint(toAdd Endpoint) (Endpoint, error) {

	response, err := accessorGroup.Database.Exec("INSERT INTO Endpoints (name, path, description) VALUES(?,?,?)", toAdd.Name, toAdd.Path, toAdd.Description)
	if err != nil {
		return Endpoint{}, err
	}

	id, err := response.LastInsertId()
	toAdd.ID = id

	return toAdd, nil
}

func (accessorGroup *AccessorGroup) RemoveEndpointByName(name string) error {

	_, err := accessorGroup.Database.Exec("DELETE FROM Endpoints WHERE name=?", name)
	if err != nil {
		return err
	}

	return nil
}

func exctractEndpointData(rows *sql.Rows) ([]Endpoint, error) {

	var endpoints []Endpoint
	var endpoint Endpoint
	var id *int64
	var name *string
	var path *string
	var description *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &path, &description)
		if err != nil {
			return []Endpoint{}, err
		}

		if id != nil {
			endpoint.ID = *id
		}
		if name != nil {
			endpoint.Name = *name
		}
		if path != nil {
			endpoint.Path = *path
		}
		if description != nil {
			endpoint.Description = *description
		}

		endpoints = append(endpoints, endpoint)

	}

	err := rows.Err()
	if err != nil {
		return []Endpoint{}, err
	}

	return endpoints, nil
}
