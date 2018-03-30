package couch

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
)

func GetDeviceByID(ID string) (structs.Device, error) {

	toReturn := structs.Device{}
	err := MakeRequest("GET", fmt.Sprintf("devices/%v", ID), "", nil, &toReturn)
	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get Device %v. %v", ID, err.Error())
		log.L.Warn(msg)
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
		msg := fmt.Sprintf("There was a problem marshalling the query: %v", err.Error())
		log.L.Warn(msg)
		return []structs.Device{}, errors.New(msg)
	}

	toReturn := structs.DeviceQueryResponse{}
	err = MakeRequest("POST", fmt.Sprintf("devices/_find"), "application/json", b, &toReturn)

	if err != nil {
		msg := fmt.Sprintf("[couch] Could not get room %v. %v", roomID, err.Error())
		log.L.Warn(msg)
	}

	//we need to go through the devices and get their type information. Hopefully caching them so we're not making a thousand requests for duplicate types.

	return toReturn.Docs, err
}

/*
Create Device. As amazing as it may seem, this fuction creates a device in the databse.

For a device to be created, it must contain the following attributes:

	1. A valid ID
		a. The room portion corresponds to an existing room
	2. A valid name
	3. A valid type
		a. Either the ID corresponds to an existing Type, or all elements are available to create a new type. Note that if the type ID matches, but the current type doesn't match the existing ID, the current type with that ID in the Database will NOT be overwritten.
	4. A valid Class
	5. One or more roles:
		a. A role must have a valid ID and Name

Ports must pass validation - criteria are covered in the CreateDeviceType function.

If a device is passed into the fuction with a valid 'rev' field, the current device with that ID will be overwritten.
`rev` must be omitted to create a new device.
*/
func CreateDevice(toAdd structs.Device) (structs.Device, error) {

}
