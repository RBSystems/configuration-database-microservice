package accessors

import (
	"database/sql"
	"log"
)

type Microservice struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func (accessorGroup *AccessorGroup) GetMicroservices() ([]Microservice, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Microservices")
	if err != nil {
		return []Microservice{}, err
	}

	microservices, err := extractMicroservices(rows)
	if err != nil {
		return []Microservice{}, err
	}

	return microservices, nil
}

func (accessorGroup *AccessorGroup) AddMicroservice(microservice Microservice) (Microservice, error) {
	result, err := accessorGroup.Database.Exec("Insert into Microservices (microserviceID, name, address, description) VALUES(?,?,?,?)", microservice.ID, microservice.Name, microservice.Address, microservice.Description)
	if err != nil {
		return Microservice{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Microservice{}, err
	}

	microservice.ID = int(id)
	return microservice, nil
}

func extractMicroservices(rows *sql.Rows) ([]Microservice, error) {
	var microservices []Microservice
	var microservice Microservice
	var id *int
	var name *string
	var address *string
	var description *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &address, &description)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return []Microservice{}, err
		}
		if id != nil {
			microservice.ID = *id
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
