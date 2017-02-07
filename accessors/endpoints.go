package accessors

import (
	"database/sql"
	"log"
)

//Endpoint represents a path on a microservice.
type Endpoint struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

//GetEndpoints queries the endpoints table for the name and path columns
func (accessorGroup *AccessorGroup) GetEndpoints() ([]Endpoint, error) {

	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("SELECT name, path FROM Endpoints")
	if err != nil {
		return []Endpoint{}, err
	}

	endpoints, err := extractEndpoints(rows)
	if err != nil {
		return []Endpoint{}, err
	}

	log.Printf("Done.")
	return endpoints, nil
}

func extractEndpoints(rows *sql.Rows) ([]Endpoint, error) {

	log.Printf("Extracting data...")
	endpoints := []Endpoint{}

	for rows.Next() {

		var endpoint Endpoint
		var name *string
		var path *string

		err := rows.Scan(&name, &path)
		if err != nil {
			return []Endpoint{}, err
		}

		log.Printf("Creating struct...")
		if name != nil {
			endpoint.Name = *name
		}
		if path != nil {
			endpoint.Path = *path
		}

		endpoints = append(endpoints, endpoint)

	}

	return endpoints, nil
}
