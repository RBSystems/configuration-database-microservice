package accessors

import (
	"database/sql"
	"log"
)

//RoomConfiguration reflects a defined room configuration with the commands and
//command keys incldued therein.
type RoomConfiguration struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	RoomKey     string                 `json:"roomKey"`
	Description string                 `json:"description"`
	Commands    []ConfigurationCommand `json:"commands"`
}

//ConfigurationCommand commands is the command information correlated with the
//specifics for the configuration (key and priority)
type ConfigurationCommand struct {
	CommandID   int    `json:"commandID"`
	CommandName string `json:"commandName"`
	Priority    int    `json:"priority"`
	CommandKey  string `json:"commandKey"`
}

//GetConfigurationByRoomAndBuilding will get the configuration information tied to a given room.
func (accessorGroup *AccessorGroup) GetConfigurationByRoomAndBuilding(building string, room string) (toReturn RoomConfiguration, err error) {

	rm, err := accessorGroup.GetRoomByBuildingAndName(building, room)
	if err != nil {
		return
	}

	toReturn, err = accessorGroup.GetConfigurationByConfigurationID(rm.ConfigurationID)
	return
}

//GetConfigurationByConfigurationName gets a configuraiton by name.
func (accessorGroup *AccessorGroup) GetConfigurationByConfigurationName(name string) (config RoomConfiguration, err error) {
	config, err = accessorGroup.GetConfigurationByQuery(`WHERE name = ?`, name)
	return
}

//GetConfigurationByConfigurationID gets a room configuration by it's ID, and fills the commands
//struct with the relevant ConfigurationCommands
func (accessorGroup *AccessorGroup) GetConfigurationByConfigurationID(configurationID int) (config RoomConfiguration, err error) {
	config, err = accessorGroup.GetConfigurationByQuery(`WHERE roomConfigurationID = ?`, configurationID)
	return

}

//GetConfigurationByQuery performs similarly to accessorGroup.GetDevicesByQuery
//You provide a WHERE statement to append to the base query, essentially allowing you to
//get any subset of information without duplicaiton of the necessary actions to extact and
//fill the data. Note that this is meant to only access the TOP 1 of any objects returned.
func (accessorGroup *AccessorGroup) GetConfigurationByQuery(queryAddition string, params ...interface{}) (config RoomConfiguration, err error) {
	baseQuery := `
	Select roomConfigurationID, name, description, roomConfigurationKey
	FROM RoomConfiguration
	`
	limit := `
	LIMIT 1
	`

	rows, err := accessorGroup.Database.Query(baseQuery+" "+queryAddition+" "+limit, params...)
	if err != nil {
		return
	}

	config, err = accessorGroup.ExtractRoomConfiguration(rows)
	if err != nil {
		return
	}

	config.Commands, err = accessorGroup.GetCommandsForConfigurationByID(config.ID)
	return
}

//GetCommandsForConfigurationByID gets the elements form the vConfiguraitonMapping table for a given configurationID
func (accessorGroup *AccessorGroup) GetCommandsForConfigurationByID(configurationID int) (allCommands []ConfigurationCommand, err error) {
	//Get configuration commands
	query := `
	Select CommandID, CommandName, CodeKey, Priority
	FROM vConfigurationMapping
	WHERE ConfigurationID = ?`

	rows, err := accessorGroup.Database.Query(query, configurationID)
	if err != nil {
		return
	}

	allCommands, err = accessorGroup.ExtractConfigurationCommand(rows)

	return
}

//ExtractRoomConfiguration pulls the items from the row to fill the config item.
func (accessorGroup *AccessorGroup) ExtractRoomConfiguration(rows *sql.Rows) (config RoomConfiguration, err error) {
	rows.Next()
	err = rows.Scan(&config.ID, &config.Name, &config.Description, &config.RoomKey)
	return
}

//ExtractConfigurationCommand pulls a set ConfigurationCommand of objects from a set of sql.Rows
func (accessorGroup *AccessorGroup) ExtractConfigurationCommand(rows *sql.Rows) (allCommands []ConfigurationCommand, err error) {

	for rows.Next() {
		command := ConfigurationCommand{}

		err = rows.Scan(&command.CommandID, &command.CommandName, &command.CommandKey, &command.Priority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		allCommands = append(allCommands, command)
	}

	return
}
