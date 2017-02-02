package accessors

import "database/sql"

//Endpoint represents a path on a microservice.
type Endpoint struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

//GetEndpoints queries the endpoints table for the name and path columns
func (accessorGroup *AccessorGroup) GetEndpoints() ([]Endpoint, error) {

	rows, err := accessorGroup.Database.Query("SELECT name, path FROM Endpoints")
	if err != nil {
		return []Endpoint{}, err
	}

	endpoints, err := extractEndpoints(rows)
	if err != nil {
		return []Endpoint{}, err
	}

	return endpoints, nil
}

func extractEndpoints(rows *sql.Rows) ([]Endpoint, error) {

	endpoints := []Endpoint{}

	for rows.Next() {

		endpoint := Endpoint{}
		err := rows.Scan(&endpoint.Name, &endpoint.Path)
		if err != nil {
			return []Endpoint{}, err
		}
		endpoints = append(endpoints, endpoint)

	}

	return endpoints, nil
}
