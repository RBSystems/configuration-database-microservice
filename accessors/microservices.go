package accessors

import (
	"database/sql"
	"log"
)

type Microservice struct {
	MicroserviceID int    `json:"microserviceID,omitempty"`
	Name           string `json:"name,omitempty"`
	Address        string `json:"address,omitempty"`
	Description    string `json:"description,omitempty"`
}

//GetAllMicroservices returns an array of Microservice structs
func (accessorGroup *AccessorGroup) GetAllMicroservices() ([]Microservice, error) {

	query := `SELECT * FROM Microservices`
	log.Printf("Querying: \"%v\"", query)
	rows, err := accessorGroup.Database.Query(query)
	if err != nil {
		return []Microservice{}, err
	}

	microservices, err := extractMicroserviceData(rows)
	if err != nil {
		return []Microservice{}, err
	}

	defer rows.Close()

	log.Printf("Done.")
	return microservices, nil
}

//ExtractMicroserviceData scans the sql columns
func extractMicroserviceData(rows *sql.Rows) ([]Microservice, error) {

	log.Printf("Extracting data...")
	var microservices []Microservice

	for rows.Next() {
		var microservice Microservice
		var ID *int
		var name *string
		var address *string
		var description *string

		err := rows.Scan(
			&ID,
			&name,
			&address,
			&description)
		if err != nil {
			return []Microservice{}, err
		}

		log.Printf("Creating struct...")
		if ID != nil {
			microservice.MicroserviceID = *ID
		}
		if name != nil {
			microservice.Name = *name
		}
		if address != nil {
			microservice.Address = *address
		}
		if description != nil {
			microservice.Description = *description
		}

		microservices = append(microservices, microservice)
	}

	return microservices, nil
}
