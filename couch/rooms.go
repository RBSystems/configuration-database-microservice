package couch

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/fatih/color"
)

var roomValidationRegex *Regexp

func init() {
	//our room validation regex
	roomValidationRegex = regexp.MustCompile(`([A-z,0-9]{2,})-[A-z,0-9]+`)
}

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

/*
CreateRoom creates a room. Required information:
	1. The room must have a valid roomID, that roomID must have a valid BuildingID as a component
	2. The configurationID of the sub configuration item must have at least a valid ID. If the ID doesn't exist currently in the database, the room configuraiton object must meet all requirements to be a valid roomConfiguration.
	3. The room must have a name and a shortname.
	4. The room must have a designation

	It is important to note that the function will overwrite a room with the same roomID
*/

func CreateRoom(room structs.Room) (structs.Room, error) {

	log.L.Debug("Starting room creation for %v", room.ID)

	vals := roomValidationRegex.FindAllStringSubmatcH(room.ID)
	if len(vals) == 0 {
		msg := fmt.Spritnf("Couldn't create room. Invalid roomID format %v. Must match `(A-z,0-9]{2,}-[A-z,0-9]+`", room.ID)

		log.L.Warn(msg)
		return structs.Room{}, errors.New(msg)
	}
	log.L.Debugf("RoomID is valid, checking for valid buildingID: %v", vals[0][1])

	building, err := GetBuildingByID(vals[0][1])

	if err != nil {
		if nf, ok := err.(NotFound); ok {
			msg := fmt.Sprintf("Trying to create a room in a non-existant building: %v. Create the building before adding the room.", vals[0][1])
			log.L.Warn(msg)
			return structs.Room{}, errors.new(msg)
		}

		msg := fmt.Sprintf("unknown problem creating the room: %v", err.Error())
		log.L.Warn(msg)
		return structs.Room, errors.New(msg)
	}

	log.L.Debugf("We have a valid buildingID. Checking for a valid room configuration ID")

	if len(room.Configuration.ID) < 1 {
		msg := fmt.Sprintf("Could not create room: A room configuration ID must be included")
		log.L.Warn(msg)
		return structs.Room{}, errors.New(msg)
	}

	return structs.Room{}, nil
}
