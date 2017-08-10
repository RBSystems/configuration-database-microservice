package accessors

import (
	"database/sql"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

//GetAllCommands simply dumps the commands table
func (accessorGroup *AccessorGroup) GetAllCommands() (commands []structs.RawCommand, err error) {
	log.Printf("Getting all commands...")
	rows, err := accessorGroup.Database.Query("Select * FROM Commands")
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}
	defer rows.Close()

	commands, err = ExtractRawCommands(rows)
	log.Printf("Done.")
	return
}

func (accessorGroup *AccessorGroup) GetRawCommandByName(name string) (structs.RawCommand, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM Commands WHERE name = ? ", name)

	rc, err := extractRawCommand(row)
	if err != nil {
		return structs.RawCommand{}, err
	}

	return rc, nil
}

//ExtractCommand pulls a command object from a set of sql.Rows
func ExtractCommand(rows *sql.Rows) (allCommands []structs.Command, err error) {

	for rows.Next() {
		command := structs.Command{}

		err = rows.Scan(&command.Name, &command.Endpoint.Name, &command.Endpoint.Path, &command.Microservice)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		allCommands = append(allCommands, command)
	}

	return
}

//ExtractRawCommands pulls a RawCommand object from a set of sql.Rows
func ExtractRawCommands(rows *sql.Rows) (allCommands []structs.RawCommand, err error) {

	for rows.Next() {
		command := structs.RawCommand{}

		err = rows.Scan(&command.ID, &command.Name, &command.Description, &command.Priority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		allCommands = append(allCommands, command)
	}

	return
}

func (accessorGroup *AccessorGroup) AddRawCommand(rc structs.RawCommand) (structs.RawCommand, error) {
	result, err := accessorGroup.Database.Exec("Insert into Commands (commandID, name, description, priority) VALUES(?,?,?,?)", rc.ID, rc.Name, rc.Description, rc.Priority)
	if err != nil {
		return structs.RawCommand{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.RawCommand{}, err
	}

	rc.ID = int(id)
	return rc, nil
}

func extractRawCommand(row *sql.Row) (structs.RawCommand, error) {
	var rc structs.RawCommand
	var id *int
	var name *string
	var description *string
	var priority *int

	err := row.Scan(&id, &name, &description, &priority)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return structs.RawCommand{}, err
	}
	if id != nil {
		rc.ID = *id
	}
	if name != nil {
		rc.Name = *name
	}
	if description != nil {
		rc.Description = *description
	}
	if priority != nil {
		rc.Priority = *priority
	}

	return rc, nil
}
