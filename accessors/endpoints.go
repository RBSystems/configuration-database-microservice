package accessors

import (
	"database/sql"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

func (accessorGroup *AccessorGroup) GetAllEndpoints() ([]structs.Endpoint, error) {

	rows, err := accessorGroup.Database.Query("SELECT * FROM Endpoints")
	if err != nil {
		return []structs.Endpoint{}, err
	}

	endpoints, err := exctractEndpointData(rows)
	if err != nil {
		return []structs.Endpoint{}, err
	}
	defer rows.Close()

	return endpoints, nil
}

func (accessorGroup *AccessorGroup) AddEndpoint(toAdd structs.Endpoint) (structs.Endpoint, error) {

	response, err := accessorGroup.Database.Exec("INSERT INTO Endpoints (name, path, description) VALUES(?,?,?)", toAdd.Name, toAdd.Path, toAdd.Description)
	if err != nil {
		return structs.Endpoint{}, err
	}

	id, err := response.LastInsertId()
	toAdd.ID = int(id)

	return toAdd, nil
}

func (accessorGroup *AccessorGroup) RemoveEndpointByName(name string) error {

	_, err := accessorGroup.Database.Exec("DELETE FROM Endpoints WHERE name=?", name)
	if err != nil {
		return err
	}

	return nil
}

func (accessorGroup *AccessorGroup) GetEndpointByName(name string) (structs.Endpoint, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM Endpoints WHERE name = ? ", name)

	e, err := extractEndpoint(row)
	if err != nil {
		return structs.Endpoint{}, err
	}

	return e, nil
}

func exctractEndpointData(rows *sql.Rows) ([]structs.Endpoint, error) {

	var endpoints []structs.Endpoint
	var endpoint structs.Endpoint
	var id *int
	var name *string
	var path *string
	var description *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &path, &description)
		if err != nil {
			return []structs.Endpoint{}, err
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
		return []structs.Endpoint{}, err
	}

	return endpoints, nil
}

func extractEndpoint(row *sql.Row) (structs.Endpoint, error) {
	var e structs.Endpoint
	var id *int
	var name *string
	var path *string
	var description *string

	err := row.Scan(&id, &name, &path, &description)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return structs.Endpoint{}, err
	}
	if id != nil {
		e.ID = *id
	}
	if name != nil {
		e.Name = *name
	}
	if path != nil {
		e.Path = *path
	}
	if description != nil {
		e.Description = *description
	}

	return e, nil
}
