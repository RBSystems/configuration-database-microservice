package accessors

import "github.com/byuoitav/configuration-database-microservice/dbstructs"

//GetAllMicroservices returns an array of Microservice structs
func (accessorGroup *AccessorGroup) GetAllMicroservices() ([]dbstructs.Microservice, error) {
	microservices := []dbstructs.Microservice{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Microservices")
	if err != nil {
		return microservices, err
	}

	defer rows.Close()

	for rows.Next() {
		microservice := dbstructs.Microservice{}

		err = rows.Scan(&dbstructs.Microservice.MicroserviceID,
			&dbstructs.Microservice.Name,
			&dbstructs.Microservice.Address,
			&dbstructs.Microservice.Description)
		if err != nil {
			return []dbstructs.Microservice{}, err
		}
	}

	return microservices, nil
}
