package couch

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/fatih/color"
)

func GetRoomByID(id string) (structs.Room, error) {

	toReturn := structs.Room{}
	err := MakeRequest("GET", fmt.Sprintf("rooms/%v", id), "", nil, &toReturn)

	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get room %v. %v", id, err.Error())
		log.Printf(color.HiRedString(msg))
	}

	//we need to get the room configuration information
	//we need to get devices

	return toReturn, err
}

func GetRoomsByBuilding(buildingID string) ([]structs.Room, error) {
	//we query from the - to . (the character after - to get all the elements in the room
	query := IDPrefixQuery{}
	query.Selector.ID.GT = fmt.Sprintf("%v-", buildingID)
	query.Selector.ID.LT = fmt.Sprintf("%v.", buildingID)
	query.Limit = 1000 //some arbitrarily large number for now.

	b, err := json.Marshal(query)
	if err != nil {
		log.Printf(color.HiRedString("There was a problem marshalling the query: %v", err.Error()))
		return []structs.Room{}, err
	}

	log.Printf("%s", b)

	toReturn := structs.RoomQueryResponse{}
	err = MakeRequest("POST", fmt.Sprintf("rooms/_find"), "application/json", b, &toReturn)

	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get room %v. %v", buildingID, err.Error())
		log.Printf(color.HiRedString(msg))
	}

	return toReturn.Docs, err
}
