package accessors

import "database/sql"

type Microservice struct {
	MicroserviceID int    `json:"microserviceID,omitempty"`
	Name           string `json:"name,omitempty"`
	Address        string `json:"address,omitempty"`
	Description    string `json:"description,omitempty"`
}

//GetAllMicroservices returns an array of Microservice structs
func (accessorGroup *AccessorGroup) GetAllMicroservices() ([]Microservice, error) {

	rows, err := accessorGroup.Database.Query("SELECT * FROM Microservices")
	if err != nil {
		return []Microservice{}, err
	}

	microservices, err := accessorGroup.ExtractMicroserviceData(rows)
	if err != nil {
		return []Microservice{}, err
	}

	defer rows.Close()

	return microservices, nil
}

//ExtractMicroserviceData scans the sql columns
func (accessorGroup *AccessorGroup) ExtractMicroserviceData(rows *sql.Rows) ([]Microservice, error) {
	microservices := []Microservice{}

	for rows.Next() {
		microservice := Microservice{}

		err := rows.Scan(
			&microservice.MicroserviceID,
			&microservice.Name,
			&microservice.Address,
			&microservice.Description)
		if err != nil {
			return []Microservice{}, err
		}
	}

	return microservices, nil
}
