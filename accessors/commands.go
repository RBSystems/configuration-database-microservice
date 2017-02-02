package accessors

import (
	"database/sql"
	"log"
)

/*Command represents all the information needed to issue a particular command to a device.
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

	rows, err := accessorGroup.Database.Query("SELECT * FROM DeviceCommands")
	if err != nil {
		return []DeviceCommand{}, err
	}

	defer rows.Close()

	deviceCommands, err := ExtractDeviceCommands(rows)
	if err != nil {
		return []DeviceCommand{}, err
	}

	return deviceCommands, nil

}

//GetAllCommands simply dumps the commands table
func (accessorGroup *AccessorGroup) GetAllCommands() (commands []Command, err error) {
	log.Printf("Getting all commands...")
	rows, err := accessorGroup.Database.Query("Select * FROM Commands")
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	commands, err = ExtractCommands(rows)
	log.Printf("Done.")
	return
}

//ExtractCommand pulls a command object from a set of sql.Rows
func ExtractDeviceCommands(rows *sql.Rows) (allCommands []DeviceCommand, err error) {

	for rows.Next() {
		command := DeviceCommand{}

		err = rows.Scan(&command.Name, &command.Endpoint.Name, &command.Endpoint.Path, &command.Microservice)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		allCommands = append(allCommands, command)
	}

	return
}

//ExtractCommands pulls a Command object from a set of sql.Rows
func ExtractCommands(rows *sql.Rows) (allCommands []Command, err error) {

	for rows.Next() {
		command := Command{}

		err = rows.Scan(&command.ID, &command.Name, &command.Description, &command.Priority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		allCommands = append(allCommands, command)
	}

	return
}
