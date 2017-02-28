package accessors

import (
	"database/sql"
	"log"
)

/*DeviceCommand represents all the information needed to issue a particular command to a device.

Name: Command name
Endpoint: the endpoint within the microservice
Microservice: the location of the microservice to call to communicate with the device.
Priority: The relative priority of the command relative to other commands. Commands
					with a higher (closer to 1) priority will be issued to the devices first.
*/
type DeviceCommand struct {
	Name         string   `json:"name"`
	Endpoint     Endpoint `json:"endpoint"`
	Microservice string   `json:"microservice"`
	Priority     int      `json:"priority"`
}

/* RawCommand represents all the information needed to issue a particular command to a device.
Name: Command name
Description: command description
Priority: The relative priority of the command relative to other commands. Commands
					with a higher (closer to 1) priority will be issued to the devices first.
*/
type RawCommand struct {
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

//CommandSorterByPriority sorts commands by priority and implements sort.Interface
type CommandSorterByPriority struct {
	Commands []RawCommand
}

//Len is part of sort.Interface
func (c *CommandSorterByPriority) Len() int {
	return len(c.Commands)
}

//Swap is part of sort.Interface
func (c *CommandSorterByPriority) Swap(i, j int) {
	c.Commands[i], c.Commands[j] = c.Commands[j], c.Commands[i]
}

//Less is part of sort.Interface
func (c *CommandSorterByPriority) Less(i, j int) bool {
	return c.Commands[i].Priority < c.Commands[j].Priority
}

//GetAllDeviceCommands simply dumps the DeviceCommands
func (accessorGroup *AccessorGroup) GetAllDeviceCommands() ([]DeviceCommand, error) {

	query := `SELECT * FROM DeviceCommands`
	log.Printf("Querying: \"%v\"", query)
	rows, err := accessorGroup.Database.Query(query)
	if err != nil {
		return []DeviceCommand{}, err
	}

	defer rows.Close()

	deviceCommands, err := ExtractDeviceCommands(rows)
	if err != nil {
		return []DeviceCommand{}, err
	}

	log.Printf("Done.")
	return deviceCommands, nil

}

//GetAllCommands simply dumps the commands table
func (accessorGroup *AccessorGroup) GetAllCommands() (commands []RawCommand, err error) {

	query := `Select * FROM Commands`
	log.Printf("Querying: \"%v\"", query)
	rows, err := accessorGroup.Database.Query(query)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	commands, err = extractCommands(rows)
	log.Printf("Done.")
	return
}

//ExtractDeviceCommands pulls a command object from a set of sql.Rows
func ExtractDeviceCommands(rows *sql.Rows) ([]DeviceCommand, error) {

	log.Printf("Extracting data...")
	var deviceCommands []DeviceCommand

	for rows.Next() {

		var deviceCommand DeviceCommand
		var tableName *string
		var tableEndpointName *string
		var tableEndpointPath *string
		var tableMicroservice *string
		var tablePriority *int

		err := rows.Scan(&tableName, &tableEndpointName, &tableEndpointPath, &tableMicroservice, &tablePriority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return []DeviceCommand{}, err
		}

		log.Printf("Creating struct...")
		if tableName != nil {
			deviceCommand.Name = *tableName
		}
		if tableEndpointName != nil {
			deviceCommand.Endpoint.Name = *tableEndpointName
		}
		if tableEndpointPath != nil {
			deviceCommand.Endpoint.Path = *tableEndpointPath
		}
		if tableMicroservice != nil {
			deviceCommand.Microservice = *tableMicroservice
		}
		if tablePriority != nil {
			deviceCommand.Priority = *tablePriority
		}

		deviceCommands = append(deviceCommands, deviceCommand)
	}

	return deviceCommands, nil
}

//ExtractCommands pulls a Command object from a set of sql.Rows
func extractCommands(rows *sql.Rows) ([]RawCommand, error) {

	log.Printf("Extracting data...")

	var commands []RawCommand

	for rows.Next() {

		var tableID *int
		var tableName *string
		var tableDescription *string
		var tablePriority *int

		var structID int
		var structName string
		var structDescription string
		var structPriority int

		err := rows.Scan(&tableID, &tableName, &tableDescription, &tablePriority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return []RawCommand{}, nil
		}

		if tableID != nil {
			structID = *tableID
		}
		if tableName != nil {
			structName = *tableName
		}
		if tableDescription != nil {
			structDescription = *tableDescription
		}
		if tablePriority != nil {
			structPriority = *tablePriority
		}

		log.Printf("Creating struct...")
		command := RawCommand{
			structID,
			structName,
			structDescription,
			structPriority,
		}

		commands = append(commands, command)
	}

	return commands, nil
}
