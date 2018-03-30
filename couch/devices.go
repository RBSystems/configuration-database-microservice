package couch

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
)

var DeviceValidationRegex *regexp.Regexp

func init() {

	DeviceValidationRegex = regexp.MustCompile(`([A-z,0-9]{2,}-[A-z,0-9]+)-[A-z]+[0-9]+`)
}

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
However in addition, if the port includes devices those devices must be valid devices

If a device is passed into the fuction with a valid 'rev' field, the current device with that ID will be overwritten.
`rev` must be omitted to create a new device.
*/
func CreateDevice(toAdd structs.Device) (structs.Device, error) {
	log.L.Infof("Starting add of Device: %v", toAdd.ID)

	log.L.Debug("Starting checks. Checking name and class.")
	if len(toAdd.Name) < 3 || len(toAdd.Class) < 3 {
		return lde(fmt.Sprintf("Couldn't create device - invalid name or Class"))
	}

	log.L.Debug("Name and class are good. Checking Roles")
	if len(toAdd.Roles) < 1 {
		return lde(fmt.Spritnf("Must include at least one role"))
	}

	for i := range toAdd.Roles {
		if err := checkRole(toAdd.Roles[i]); err != nil {
			return lde(fmt.Sprintf("Couldn't create device: %v", err.Error()))
		}
	}
	log.L.Debug("Roles are all valid. Checking ID")

	vals := deviceValidationRegex.FindAllStringSubmatch(room.ID, 1)
	if len(vals) == 0 {
		return lde(fmt.Sprintf("Couldn't create Device. Invalid deviceID format %v. Must match `[A-z,0-9]{2,}-[A-z,0-9]+-[A-z]+[0-9]+`", room.ID))
	}

	log.L.Debug("Device ID is well formed, checking for valid room.")

	_, err := GetRoomByID(vals[0][1])

	if err != nil {
		if nf, ok := err.(NotFound); ok {
			return lde(fmt.Sprintf("Trying to create a device in a non-existant Room: %v. Create the room before adding the device. message: ", vals[0][1], nf.Error()))
		}

		return lde(fmt.Sprintf("unknown problem creating the device: %v", err.Error()))
	}
	log.L.Debug("Device has a valid roomID. Checking for a valid type.")

	if len(toAdd.Type.ID) < 1 {
		return lde("Couldn't create a device, a type ID must be included")
	}

	deviceType, err := GetDeviceTypeByID(toAdd.Type.ID)
	if err != nil {
		if nf, ok := err.(NotFound); ok {
			log.L.Debug("Device Type not found, attempting to create. Message: %v", nf.Error())

			deviceType, err := CreateDeviceType(toAdd.Type)
			if err != nil {
				return lde("Trying to create a device with a non-existant device type and not enough information to create the type. Check the included type ID")
			}
			log.L.Debug("Type created")
		} else {
			ldt(fmt.Sprintf("Unkown issue creating the device: %v", err.Error()))
		}
	}
	log.L.Debug("Type is good. Checking ports.")

	for i := range toAdd.Ports {
		if err := checkPort(toAdd.Ports[i]); err != nil {
			return lde(fmt.Sprintf("Couldn't create device: %v", err.Error()))
		}
	}

	log.L.Debug("Ports are good. Checks passed. Creating device.")
}

//log device error
//alias to help cut down on cruft
func lde(msg string) (dev structs.Device, err error) {
	log.L.Warn(msg)
	err = errors.New(msg)
	return
}

func checkRole(r structs.Role) error {
	if len(r.ID) < 3 || len(r.Name) < 3 {
		return errors.New("Invalid role, check name and ID.")
	}
}

func checkPort(p structs.Port) error {
	if !validatePort(p) {
		return errors.New("Invalid port, check Name, ID, and Port Type")
	}

	//now we need to check the source and destination device
	if len(p.SourceDevice) > 0 {
		if _, err := GetDeviceByID(p.SourceDevice); err != nil {
			return errors.New(fmt.Spritnf("Invalid port %v, source device %v doesn't exist. Create it before adding it to a port", p.ID, p.SourceDevice))
		}
	}
	if len(p.DestiationDevice) > 0 {
		if _, err := GetDeviceByID(p.DestinationDevice); err != nil {
			return errors.New(fmt.Spritnf("Invalid port %v, destination device %v doesn't exist. Create it before adding it to a port", p.ID, p.DestinationDevice))
		}
	}

	//we're all good
	return nil
}
