package couch

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/fatih/color"
)

func GetDeviceByID(ID string) (structs.Device, error) {

	toReturn := structs.Device{}
	err := MakeRequest("GET", fmt.Sprintf("devices/%v", ID), "", nil, &toReturn)

	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get Device %v. %v", ID, err.Error())
		log.Printf(color.HiRedString(msg))
	}

	return toReturn, err
}

func GetDevicesByRoom(roomID string) ([]structs.Device, error) {
	//we query from the - to . (the character after - to get all the elements in the room
	query := IDPrefixQuery{}
	query.Selector.ID.GT = fmt.Sprintf("%v-", roomID)
	query.Selector.ID.LT = fmt.Sprintf("%v.", roomID)
	query.Limit = 1000 //some arbitrarily large number for now.

	b, err := json.Marshal(query)
	if err != nil {
		log.Printf(color.HiRedString("There was a problem marshalling the query: %v", err.Error()))
		return []structs.Device{}, err
	}

	toReturn := structs.DeviceQueryResponse{}
	err = MakeRequest("POST", fmt.Sprintf("devices/_find"), "application/json", b, &toReturn)

	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get room %v. %v", roomID, err.Error())
		log.Printf(color.HiRedString(msg))
	}

	//we need to go through the devices and get their type information. Hopefully caching them so we're not making a thousand requests for duplicate types.

	return toReturn.Docs, err
}

func GetDeviceTypes(companyID string) ([]structs.DeviceType, error) {
	return []structs.DeviceType{}, nil
}

func GetDeviceTypesByID(companyID, deviceTypeID string) (structs.DeviceType, error) {
	return structs.DeviceType{}, nil
}
