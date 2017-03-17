package accessors

import (
	"database/sql"
	"log"
)

/* Command represents all the information needed to issue a particular command to a device.

Name: Command name
Endpoint: the endpoint within the microservice
Microservice: the location of the microservice to call to communicate with the device.
Priority: The relative priority of the command relative to other commands. Commands
					with a higher (closer to 1) priority will be issued to the devices first.
*/
type Command struct {
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
func (accessorGroup *AccessorGroup) GetAllDeviceCommands() ([]Command, error) {

	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceCommands")
	if err != nil {
		return []Command{}, err
	}

	defer rows.Close()

	deviceCommands, err := ExtractCommands(rows)
	if err != nil {
		return []Command{}, err
	}

	log.Printf("Done.")
	return deviceCommands, nil

}

func (accessorGroup *AccessorGroup) GetAllCommands() (commands []RawCommand, err error) {
	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("Select * FROM Commands")
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
	defer rows.Close()

	commands, err = extractRawCommands(rows)
	log.Printf("Done.")
	return
}

//ExtractCommands pulls a command object from a set of sql.Rows
func ExtractCommands(rows *sql.Rows) (allCommands []Command, err error) {

	log.Printf("Extracting data...")
	for rows.Next() {

		var command Command
		var endpoint Endpoint

		var tableName *string
		var tableEndpointName *string
		var tableEndpointPath *string
		var tableMicroservice *string
		var tablePriority *int

		err = rows.Scan(&tableName, &tableEndpointName, &tableEndpointPath, &tableMicroservice, &tablePriority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		if tableName != nil {
			command.Name = *tableName
		}
		if tableEndpointName != nil {
			endpoint.Name = *tableEndpointName
		}
		if tableEndpointPath != nil {
			endpoint.Path = *tableEndpointPath
		}
		if tableMicroservice != nil {
			command.Microservice = *tableMicroservice
		}
		if tablePriority != nil {
			command.Priority = *tablePriority
		}

		allCommands = append(allCommands, command)
	}

	return
}

//ExtractCommands pulls a Command object from a set of sql.Rows
func extractRawCommands(rows *sql.Rows) ([]RawCommand, error) {

	log.Printf("Extracting data...")

	var commands []RawCommand

	for rows.Next() {

		var command RawCommand

		var tableID *int
		var tableName *string
		var tableDescription *string
		var tablePriority *int

		err := rows.Scan(&tableID, &tableName, &tableDescription, &tablePriority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return []RawCommand{}, nil
		}

		if tableID != nil {
			command.ID = *tableID
		}
		if tableName != nil {
			command.Name = *tableName
		}
		if tableDescription != nil {
			command.Description = *tableDescription
		}
		if tablePriority != nil {
			command.Priority = *tablePriority
		}

		commands = append(commands, command)
	}

	return commands, nil
}
