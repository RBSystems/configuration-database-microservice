package accessors

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
)

//Device represents a device object as found in the DB.
type Device struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Address     string              `json:"address"`
	Input       bool                `json:"input"`
	Output      bool                `json:"output"`
	Building    Building            `json:"building"`
	Room        Room                `json:"room"`
	Type        string              `json:"type"`
	Power       string              `json:"power"`
	Blanked     *bool               `json:"blanked,omitempty"`
	Volume      *int                `json:"volume,omitempty"`
	Muted       *bool               `json:"muted,omitempty"`
	PowerStates []string            `json:"powerstates,omitempty"`
	Responding  bool                `json:"responding"`
	Ports       []PortConfiguration `json:"ports,omitempty"`
	Commands    []DeviceCommand     `json:"commands,omitempty"`
}

//GetFullName reutrns the string of building + room + name
func (d *Device) GetFullName() string {
	return (d.Building.Shortname + "-" + d.Room.Name + "-" + d.Name)
}

//PortConfiguration represents a physical port on a device (HDMI, DP, Audo, etc.)
type PortConfiguration struct {
	Source      string `json:"source"`
	Name        string `json:"name"`
	Destination string `json:"destination"`
	Host        string `json:"host"`
}

//DeviceType represents the DeviceType table in the database
type DeviceType struct {
	DeviceTypeID int    `json:"deviceTypeID"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

//PowerState represents the PowerState table in the database
type PowerState struct {
	PowerStateID int    `json:"powerStateID"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

//GetDeviceTypes simply returns the contents of the DeviceTypes table
func (accessorGroup *AccessorGroup) GetDeviceTypes() ([]DeviceType, error) {

	query := `SELECT * FROM DeviceTypes`
	log.Printf("Querying: \"%v\"", query)
	rows, err := accessorGroup.Database.Query(query)
	if err != nil {
		return []DeviceType{}, err
	}
	defer rows.Close()

	deviceTypes, err := extractDeviceTypes(rows)
	if err != nil {
		return []DeviceType{}, err
	}

	log.Printf("Done.")
	return deviceTypes, nil
}

func extractDeviceTypes(rows *sql.Rows) ([]DeviceType, error) {

	log.Printf("Extracting data...")
	var deviceTypes []DeviceType

	for rows.Next() {

		var deviceType DeviceType
		var ID *int
		var name *string
		var description *string

		err := rows.Scan(&ID, &name, &description)
		if err != nil {
			return []DeviceType{}, err
		}

		log.Printf("Creating struct...")
		if ID != nil {
			deviceType.DeviceTypeID = *ID
		}
		if name != nil {
			deviceType.Name = *name
		}
		if description != nil {
			deviceType.Description = *description
		}

		deviceTypes = append(deviceTypes, deviceType)

	}

	return deviceTypes, nil
}

//GetPowerStates simply returns the contents of the PowerStates table
func (accessorGroup *AccessorGroup) GetPowerStates() ([]PowerState, error) {

	query := `SELECT * FROM PowerStates`
	log.Printf("Querying: \"%v\"", query)
	rows, err := accessorGroup.Database.Query(query)
	if err != nil {
		return []PowerState{}, err
	}

	defer rows.Close()

	powerStates, err := extractPowerStates(rows)
	if err != nil {
		return []PowerState{}, err
	}

	log.Printf("Done.")
	return powerStates, nil
}

func extractPowerStates(rows *sql.Rows) ([]PowerState, error) {

	log.Printf("Extracting data...")
	var powerStates []PowerState

	for rows.Next() {
		var powerState PowerState
		var tableID *int
		var tableName *string
		var tableDescription *string

		err := rows.Scan(&tableID, &tableName, &tableDescription)
		if err != nil {
			return []PowerState{}, err
		}

		log.Printf("Creating struct...")
		if tableID != nil {
			powerState.PowerStateID = *tableID
		}
		if tableName != nil {
			powerState.Name = *tableName
		}
		if tableDescription != nil {
			powerState.Description = *tableDescription
		}

		powerStates = append(powerStates, powerState)

	}

	return powerStates, nil
}

/*
GetDevicesByQuery is a function that abstracts some of the execution and extraction
of data from the database when we're looking for responses based on the COMPLETE device struct.
The function MAY have the WHERE clause passed in to limit the devices found.
The function MAY have any JOIN clauses necessary to the WEHRE Clause not included in
the base query.
JOIN statements in the base query:
JOIN Rooms on Devices.roomID = Rooms.RoomID
JOIN Buildings on Rooms.buildingID = Buildings.buildingID
JOIN DeviceTypes on Devices.typeID = DeviceTypes.deviceTypeID
If empty string is passed in no WHERE clause will be appended, and thus all devices
will be returned.

Flow	->	Find all devices based on the clause passed in
			->	For each device found find the Ports
			->	For each device found find the Commands

Examples of valid parameters.
Example 1:
`JOIN deviceRole on deviceRole.deviceID = Devices.deviceID
JOIN DeviceRoleDefinition on DeviceRole.deviceRoleDefinitionID = DeviceRoleDefinition.deviceRoleDefinitionID
WHERE DeviceRoleDefinition.name LIKE 'AudioIn'`
Example 2:
`WHERE Devices.RoomID = 1`
*/
func (accessorGroup *AccessorGroup) GetDevicesByQuery(query string, parameters ...interface{}) ([]Device, error) {
	baseQuery := `SELECT DISTINCT Devices.deviceID,
  	Devices.Name as deviceName,
  	Devices.address as deviceAddress,
  	Devices.input,
  	Devices.output,
  	Rooms.roomID,
  	Rooms.name as roomName,
  	Rooms.description as roomDescription,
  	Buildings.buildingID,
  	Buildings.name as buildingName,
  	Buildings.shortName as buildingShortname,
  	Buildings.description as buildingDescription,
  	DeviceTypes.name as deviceType
  	FROM Devices
  	JOIN Rooms on Rooms.roomID = Devices.roomID
  	JOIN Buildings on Buildings.buildingID = Devices.buildingID
  	JOIN DeviceTypes on Devices.typeID = DeviceTypes.deviceTypeID
    JOIN DeviceRole on DeviceRole.deviceID = Devices.deviceID
    JOIN DeviceRoleDefinition on DeviceRole.deviceRoleDefinitionID = DeviceRoleDefinition.deviceRoleDefinitionID`

	var allDevices []Device

	log.Printf("Querying: \"%v %v %v\"", baseQuery, query, parameters)
	rows, err := accessorGroup.Database.Query(baseQuery+" "+query, parameters...)
	if err != nil {
		return []Device{}, err
	}

	defer rows.Close()

	for rows.Next() {

		device := Device{}

		err := rows.Scan(&device.ID,
			&device.Name,
			&device.Address,
			&device.Input,
			&device.Output,
			&device.Room.ID,
			&device.Room.Name,
			&device.Room.Description,
			&device.Building.ID,
			&device.Building.Name,
			&device.Building.Shortname,
			&device.Building.Description,
			&device.Type)
		if err != nil {
			return []Device{}, err
		}

		device.Commands, err = accessorGroup.GetDeviceCommandsByBuildingAndRoomAndName(device.Building.Shortname, device.Room.Name, device.Name)
		if err != nil {
			return []Device{}, err
		}

		device.Ports, err = accessorGroup.GetDevicePortsByBuildingAndRoomAndName(device.Building.Shortname, device.Room.Name, device.Name)
		if err != nil {
			return []Device{}, err
		}

		device.PowerStates, err = accessorGroup.GetPowerStatesByDeviceID(device.ID)
		if err != nil {
			return []Device{}, err
		}

		allDevices = append(allDevices, device)
	}

	return allDevices, nil
}

//GetPowerStatesByDeviceID gets the powerstates allowed for a given devices based on the
//DevicePowerStates table in the DB.
func (accessorGroup *AccessorGroup) GetPowerStatesByDeviceID(deviceID int) ([]string, error) {

	query := `SELECT PowerStates.name FROM PowerStates
	JOIN DevicePowerStates on DevicePowerStates.powerStateID = PowerStates.powerStateID
	Where DevicePowerStates.deviceID = ?`

	log.Printf("Querying: \"%v %v\"", query, deviceID)

	toReturn := []string{}
	rows, err := accessorGroup.Database.Query(query, deviceID)
	if err != nil {
		return []string{}, err
	}

	log.Printf("Extracting data...")
	for rows.Next() {
		var tableValue *string
		var value string

		err := rows.Scan(&tableValue)
		if err != nil {
			return []string{}, err
		}

		if tableValue != nil {
			value = *tableValue
		}

		toReturn = append(toReturn, value)
	}

	log.Printf("Done.")
	return toReturn, nil
}

//GetDevicesByBuildingAndRoomAndRole gets the devices in the room specified with the given role,
//as specified in the DeviceRole table in the DB
func (accessorGroup *AccessorGroup) GetDevicesByBuildingAndRoomAndRole(buildingShortname string, roomName string, roleName string) ([]Device, error) {
	log.Printf("Getting ")
	devices, err := accessorGroup.GetDevicesByQuery(`WHERE Rooms.name LIKE ? AND Buildings.shortname LIKE ? AND DeviceRoleDefinition.name LIKE ?`,
		roomName, buildingShortname, roleName)

	if err != nil {
		log.Printf("Error: %v", err.Error())
		return []Device{}, err
	}
	switch strings.ToLower(roleName) {

	}
	return devices, nil

}

//GetDevicesByRoleAndType Gets all teh devices that have a given role and type.
func (accessorGroup *AccessorGroup) GetDevicesByRoleAndType(deviceRole string, deviceType string) ([]Device, error) {
	return accessorGroup.GetDevicesByQuery(`WHERE DeviceRoleDefinition.name LIKE ? AND DeviceTypes.name LIKE ?`, deviceRole, deviceType)
}

//GetDevicesByBuildingAndRoom get all the devices in the room specified.
func (accessorGroup *AccessorGroup) GetDevicesByBuildingAndRoom(buildingShortname string, roomName string) ([]Device, error) {
	log.Printf("Getting devices in room %s and building %s", roomName, buildingShortname)

	devices, err := accessorGroup.GetDevicesByQuery(
		`WHERE Rooms.name=? AND Buildings.shortName=?`, roomName, buildingShortname)

	if err != nil {
		return []Device{}, err
	}

	return devices, nil
}

//GetDeviceCommandsByBuildingAndRoomAndName gets all the commands for the device
//specified. Note that we assume that device names are unique within a room.
func (accessorGroup *AccessorGroup) GetDeviceCommandsByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) ([]DeviceCommand, error) {

	query := `SELECT Commands.name as commandName, Endpoints.name as endpointName, Endpoints.path as endpointPath, Microservices.address as microserviceAddress, Commands.priority as commandPriority
	FROM Devices
	JOIN DeviceCommands on Devices.deviceID = DeviceCommands.deviceID JOIN Commands on DeviceCommands.commandID = Commands.commandID JOIN Endpoints on DeviceCommands.endpointID = Endpoints.endpointID JOIN Microservices ON DeviceCommands.microserviceID = Microservices.microserviceID
	JOIN Rooms ON Rooms.roomID=Devices.roomID
	JOIN Buildings ON Rooms.buildingID=Buildings.buildingID
	WHERE Rooms.name=? AND Buildings.shortName=? AND Devices.name=?`

	log.Printf("Querying: \"%v %v %v %v\"", query, roomName, buildingShortname, deviceName)

	rows, err := accessorGroup.Database.Query(query, roomName, buildingShortname, deviceName)
	if err != nil {
		return []DeviceCommand{}, err
	}

	allCommands, err := ExtractDeviceCommands(rows)
	if err != nil {
		return allCommands, err
	}

	log.Printf("Done.")
	return allCommands, nil
}

//GetDevicePortsByBuildingAndRoomAndName gets the ports for the device
//specified. Note that we assume that device names are unique within a room.
func (accessorGroup *AccessorGroup) GetDevicePortsByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) ([]PortConfiguration, error) {

	query := `SELECT srcDevice.Name as sourceName, Ports.name as portName, destDevice.Name as DestinationDevice, hostDevice.name as HostDevice FROM Ports
	JOIN PortConfiguration ON Ports.PortID = PortConfiguration.PortID
	JOIN Devices as srcDevice on srcDevice.DeviceID = PortConfiguration.sourceDeviceID
	JOIN Devices as destDevice on destDevice.DeviceID = PortConfiguration.destinationDeviceID
	JOIN Devices as hostDevice on hostDevice.DeviceID = PortConfiguration.hostDeviceID
	JOIN Rooms ON Rooms.roomID=destDevice.roomID
	JOIN Buildings ON Rooms.buildingID=Buildings.buildingID
	WHERE Rooms.name=? AND Buildings.shortName=? AND hostDevice.name=?`
	log.Printf("Querying: \"%v %v %v %v\"", query, roomName, buildingShortname, deviceName)
	rows, err := accessorGroup.Database.Query(query, roomName, buildingShortname, deviceName)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return []PortConfiguration{}, err
	}

	defer rows.Close()

	portConfigurations, err := extractPortConfigurations(rows)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return []PortConfiguration{}, err
	}

	log.Printf("Done.")
	return portConfigurations, nil
}

func extractPortConfigurations(rows *sql.Rows) ([]PortConfiguration, error) {

	log.Printf("Extracting data...")

	var portConfigurations []PortConfiguration

	for rows.Next() {

		var portConfiguration PortConfiguration
		var tableSource *string
		var tableName *string
		var tableDestinaion *string
		var tableHost *string

		err := rows.Scan(&tableSource, &tableName, &tableDestinaion, &tableHost)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return []PortConfiguration{}, err
		}

		log.Printf("Creating struct...")
		if tableSource != nil {
			portConfiguration.Source = *tableSource
		}
		if tableName != nil {
			portConfiguration.Name = *tableName
		}
		if tableDestinaion != nil {
			portConfiguration.Destination = *tableDestinaion
		}
		if tableHost != nil {
			portConfiguration.Host = *tableHost
		}

		portConfigurations = append(portConfigurations, portConfiguration)
	}

	return portConfigurations, nil

}

//GetDeviceByBuildingAndRoomAndName gets the device
//specified. Note that we assume that device names are unique within a room.
func (accessorGroup *AccessorGroup) GetDeviceByBuildingAndRoomAndName(buildingShortname string, roomName string, deviceName string) (Device, error) {
	dev, err := accessorGroup.GetDevicesByQuery("WHERE Buildings.shortName = ? AND Rooms.name = ? AND Devices.name = ?", buildingShortname, roomName, deviceName)
	if err != nil {
		return Device{}, err
	}

	return dev[0], nil
}

//PutDeviceAttributeByDeviceAndRoomAndBuilding allows you to change attribute values for devices
//Currently sets volume and muted.
func (accessorGroup *AccessorGroup) PutDeviceAttributeByDeviceAndRoomAndBuilding(building string, room string, device string, attribute string, attributeValue string) (Device, error) {
	switch strings.ToLower(attribute) {
	case "volume":
		statement := `update AudioDevices SET volume = ? WHERE deviceID =
			(Select deviceID from Devices
				JOIN Rooms on Rooms.roomID = Devices.roomID
				JOIN Buildings on Buildings.buildingID = Rooms.buildingID
				WHERE Devices.name LIKE ? AND Rooms.name LIKE ? AND Buildings.shortName LIKE ?)`
		val, err := strconv.Atoi(attributeValue)
		if err != nil {
			return Device{}, err
		}

		_, err = accessorGroup.Database.Exec(statement, val, device, room, building)
		if err != nil {
			return Device{}, err
		}
		break

	case "muted":
		var valToSet bool
		switch attributeValue {
		case "true":
			valToSet = true
			break
		case "false":
			valToSet = false
			break
		default:
			return Device{}, errors.New("Invalid attribute value, must be a boolean.")
		}
		statement := `update AudioDevices SET muted = ? WHERE deviceID =
			(Select deviceID from Devices
				JOIN Rooms on Rooms.roomID = Devices.roomID
				JOIN Buildings on Buildings.buildingID = Rooms.buildingID
				WHERE Devices.name LIKE ? AND Rooms.name LIKE ? AND Buildings.shortName LIKE ?)`
		_, err := accessorGroup.Database.Exec(statement, valToSet, device, room, building)
		if err != nil {
			return Device{}, err
		}
		break
	}

	dev, err := accessorGroup.GetDeviceByBuildingAndRoomAndName(building, room, device)
	return dev, err
}
