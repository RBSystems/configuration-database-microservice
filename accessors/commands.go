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

/*Command represents all the information needed to issue a particular command to a device.
Name: Command name
Description: command description
Priority: The relative priority of the command relative to other commands. Commands
					with a higher (closer to 1) priority will be issued to the devices first.
*/
type Command struct {
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

//CommandSorterByPriority sorts commands by priority and implements sort.Interface
type CommandSorterByPriority struct {
	Commands []Command
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

	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceCommands")
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
func (accessorGroup *AccessorGroup) GetAllCommands() (commands []Command, err error) {
	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("Select * FROM Commands")
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	commands, err = extractCommands(rows)
	log.Printf("Done.")
	return
}

//ExtractDeviceCommands pulls a command object from a set of sql.Rows
func ExtractDeviceCommands(rows *sql.Rows) (allCommands []DeviceCommand, err error) {

	log.Printf("Extracting data...")
	for rows.Next() {

		var tableName *string
		var tableEndpointName *string
		var tableEndpointPath *string
		var tableMicroservice *string
		var tablePriority *int

		var structName string
		var structEndpointName string
		var structEndpointPath string
		var structMicroservice string
		var structPriority int

		err = rows.Scan(&tableName, &tableEndpointName, &tableEndpointPath, &tableMicroservice, &tablePriority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		if tableName != nil {
			structName = *tableName
		}
		if tableEndpointName != nil {
			structEndpointName = *tableEndpointName
		}
		if tableEndpointPath != nil {
			structEndpointPath = *tableEndpointPath
		}
		if tableMicroservice != nil {
			structMicroservice = *tableMicroservice
		}
		if tablePriority != nil {
			structPriority = *tablePriority
		}

		log.Printf("Creating struct...")
		structEndpoint := Endpoint{
			structEndpointName,
			structEndpointPath,
		}
		command := DeviceCommand{
			structName,
			structEndpoint,
			structMicroservice,
			structPriority,
		}

		allCommands = append(allCommands, command)
	}

	return
}

//ExtractCommands pulls a Command object from a set of sql.Rows
func extractCommands(rows *sql.Rows) ([]Command, error) {

	log.Printf("Extracting data...")

	var commands []Command

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
			return []Command{}, nil
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
		command := Command{
			structID,
			structName,
			structDescription,
			structPriority,
		}

		commands = append(commands, command)
	}

	return commands, nil
}
