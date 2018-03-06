package couch

import (
	"fmt"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/fatih/color"
)

//GetBuildingByID gets the company's building with the corresponding ID from the couch database
func GetBuildingByID(id string) (structs.Building, error) {

	toReturn := structs.Building{}
	err := MakeRequest("GET", fmt.Sprintf("buildings/%v", id), "", nil, &toReturn)

	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get building %v. %v", id, err.Error())
		log.Printf(color.HiRedString(msg))
	}

	return toReturn, err
}

//GetAllBuildings returns all buildings for the company specified
func GetAllBuildings() ([]structs.Building, error) {

	toFill := structs.BuildingQueryResponse{}

	err := MakeRequest("GET", fmt.Sprintf("buildings/_all_docs?limit=1000&include_docs=true"), "", nil, &toFill)
	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get buildings. %v", err.Error())
		log.Printf(color.HiRedString(msg))
	}

	toReturn := []structs.Building{}
	for _, row := range toFill.Rows {
		toReturn = append(toReturn, row.Doc)
	}

	return toReturn, err
}
