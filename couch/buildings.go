package couch

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
)

//GetBuildingByID gets the company's building with the corresponding ID from the couch database
func GetBuildingByID(id string) (structs.Building, error) {

	toReturn := structs.Building{}
	err := MakeRequest("GET", fmt.Sprintf("buildings/%v", id), "", nil, &toReturn)
	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get building %v. %v", id, err.Error())
		log.L.Warn(msg)
	}

	return toReturn, err
}

//GetAllBuildings returns all buildings for the company specified
func GetAllBuildings() ([]structs.Building, error) {

	toFill := structs.BuildingQueryResponse{}

	err := MakeRequest("GET", fmt.Sprintf("buildings/_all_docs?limit=1000&include_docs=true"), "", nil, &toFill)
	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get buildings. %v", err.Error())
		log.L.Warn(msg)
	}

	toReturn := []structs.Building{}
	for _, row := range toFill.Rows {
		toReturn = append(toReturn, row.Doc)
	}

	return toReturn, err
}

/*
AddBuilding adds a building. The building must have at least:
1) ID
2) Name

The function will also overwrite the existing building providing the _rev field is set properly
*/
func CreateBuilding(toAdd structs.Building) (structs.Building, error) {
	log.L.Debugf("Starting adding a building: %v", toAdd.Name)

	if len(toAdd.ID) < 2 || len(toAdd.Name) < 2 {
		msg := "Cannot create building, must have at least a name and an ID"
		log.L.Warn(msg)
	}

	//there's not a lot to buildings, so we can just add

	log.L.Debug("Checks passed, creating building.")

	b, err := json.Marshal(toAdd)
	if err != nil {

		msg := fmt.Sprintf("Couldn't marshal building into JSON. Error: %v", err.Error())
		log.L.Error(msg) // this is a slightly bigger deal
		return toAdd, errors.New(msg)
	}

	resp := CouchUpsertResponse{}

	err = MakeRequest("PUT", fmt.Sprintf("buildings/%v", toAdd.ID), "", b, &resp)
	if err != nil {
		log.L.Debugf("%v", err)
		if conflict, ok := err.(*Confict); ok {
			msg := fmt.Sprintf("Error: %v Make changes on an updated version of the configuration.", conflict.Error())
			log.L.Warn(msg)
			return toAdd, errors.New(msg)
		}
		//ther was some other problem
		msg := fmt.Sprintf("unknown problem creating the Building: %v", err.Error())
		log.L.Warn(msg)
		return toAdd, errors.New(msg)
	}

	log.L.Debug("Building created, retriving new configuration from database.")

	//return the created config
	toAdd, err = GetBuildingByID(toAdd.ID)
	if err != nil {
		msg := fmt.Sprintf("There was a problem getting the newly created building: %v", err.Error())
		log.L.Warn(msg)
		return toAdd, errors.New(msg)
	}

	log.L.Debug("Done.")
	return toAdd, nil
}

func DeleteBuilding(id string) error {
	building, err := GetBuildingByID(id)
	if err != nil {
		msg := fmt.Sprintf("There was a problem deleting the building: %v", err.Error())
		log.L.Warn(msg)
		return errors.New(msg)
	}

	err = MakeRequest("DELETE", fmt.Sprintf("buildings/%s?rev=%v", id, building.Rev), "", nil, nil)
	if err != nil {
		msg := fmt.Sprintf("There was a problem deleting the building: %v", err.Error())
		log.L.Warn(msg)
		return errors.New(msg)
	}

	return nil
}
