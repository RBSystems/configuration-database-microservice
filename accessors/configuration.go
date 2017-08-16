package accessors

import (
	"database/sql"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

//GetConfigurationByRoomAndBuilding will get the configuration information tied to a given room.
func (accessorGroup *AccessorGroup) GetConfigurationByRoomAndBuilding(building string, room string) (toReturn structs.RoomConfiguration, err error) {

	rm, err := accessorGroup.GetRoomByBuildingAndName(building, room)
	if err != nil {
		return
	}

	toReturn, err = accessorGroup.GetConfigurationByConfigurationID(rm.ConfigurationID)
	return
}

//GetConfigurationByConfigurationName gets a configuraiton by name.
func (accessorGroup *AccessorGroup) GetConfigurationByConfigurationName(name string) (config structs.RoomConfiguration, err error) {
	config, err = accessorGroup.GetConfigurationByQuery(`WHERE name = ?`, name)
	return
}

//GetConfigurationByConfigurationID gets a room configuration by it's ID, and fills the commands
//struct with the relevant ConfigurationEvaluators
func (accessorGroup *AccessorGroup) GetConfigurationByConfigurationID(configurationID int) (config structs.RoomConfiguration, err error) {
	config, err = accessorGroup.GetConfigurationByQuery(`WHERE roomConfigurationID = ?`, configurationID)
	return

}

//GetConfigurationByQuery performs similarly to accessorGroup.GetDevicesByQuery
//You provide a WHERE statement to append to the base query, essentially allowing you to
//get any subset of information without duplicaiton of the necessary actions to extact and
//fill the data. Note that this is meant to only access the TOP 1 of any objects returned.
func (accessorGroup *AccessorGroup) GetConfigurationByQuery(queryAddition string, params ...interface{}) (config structs.RoomConfiguration, err error) {
	baseQuery := `
	Select roomConfigurationID, name, description, roomConfigurationKey, roomInitializationKey
	FROM RoomConfiguration
	`
	limit := `
	LIMIT 1
	`

	rows, err := accessorGroup.Database.Query(baseQuery+" "+queryAddition+" "+limit, params...)
	if err != nil {
		return
	}
	defer rows.Close()

	config, err = accessorGroup.ExtractRoomConfiguration(rows)
	if err != nil {
		return
	}

	config.Evaluators, err = accessorGroup.GetEvaluatorsForConfigurationByID(config.ID)

	return
}

//GetEvaluatorsForConfigurationByID gets the elements form the vConfiguraitonMapping table for a given configurationID
func (accessorGroup *AccessorGroup) GetEvaluatorsForConfigurationByID(configurationID int) (allEvaluators []structs.ConfigurationEvaluator, err error) {
	//Get configuration commands
	query := `
	Select EvaluatorKey, Priority
	FROM vConfigurationMapping
	WHERE ConfigurationID = ?`

	rows, err := accessorGroup.Database.Query(query, configurationID)
	if err != nil {
		return
	}
	defer rows.Close()

	allEvaluators, err = accessorGroup.ExtractConfigurationEvaluator(rows)

	return
}

func (accessorGroup *AccessorGroup) GetConfigurations() ([]structs.RoomConfiguration, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM RoomConfiguration")
	if err != nil {
		return []structs.RoomConfiguration{}, err
	}

	rcs, err := extractRoomConfigurations(rows)
	if err != nil {
		return []structs.RoomConfiguration{}, err
	}
	defer rows.Close()

	return rcs, nil
}

func extractRoomConfigurations(rows *sql.Rows) ([]structs.RoomConfiguration, error) {
	var rcs []structs.RoomConfiguration
	var rc structs.RoomConfiguration
	var id *int
	var name *string
	var roomkey *string
	var description *string
	var roominitkey *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &roomkey, &description, &roominitkey)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
		if id != nil {
			rc.ID = *id
		}
		if name != nil {
			rc.Name = *name
		}
		if roomkey != nil {
			rc.RoomKey = *roomkey
		}
		if description != nil {
			rc.Description = *description
		}
		if roominitkey != nil {
			rc.RoomInitKey = *roominitkey
		}
		rcs = append(rcs, rc)
	}
	return rcs, nil
}

//ExtractRoomConfiguration pulls the items from the row to fill the config item.
func (accessorGroup *AccessorGroup) ExtractRoomConfiguration(rows *sql.Rows) (config structs.RoomConfiguration, err error) {
	rows.Next()

	err = rows.Scan(&config.ID, &config.Name, &config.Description, &config.RoomKey, &config.RoomInitKey)

	return
}

//ExtractConfigurationEvaluator pulls a set ConfigurationEvaluator of objects from a set of sql.Rows
func (accessorGroup *AccessorGroup) ExtractConfigurationEvaluator(rows *sql.Rows) (allEvaluators []structs.ConfigurationEvaluator, err error) {

	for rows.Next() {
		command := structs.ConfigurationEvaluator{}

		err = rows.Scan(&command.EvaluatorKey, &command.Priority)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}

		allEvaluators = append(allEvaluators, command)
	}

	return
}
