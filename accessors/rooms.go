package accessors

import (
	"database/sql"
	"errors"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

// GetAllRooms returns a list of rooms from the database
func (accessorGroup *AccessorGroup) GetAllRooms() ([]structs.Room, error) {
	allBuildings := []structs.Building{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Buildings")
	if err != nil {
		return []structs.Room{}, err
	}

	for rows.Next() {
		building := structs.Building{}

		err = rows.Scan(&building.ID, &building.Name, &building.Shortname, &building.Description)
		if err != nil {
			return []structs.Room{}, err
		}

		allBuildings = append(allBuildings, building)
	}

	allRooms := []structs.Room{}

	//	rows, err = accessorGroup.Database.Query("SELECT * FROM Rooms WHERE roomDesignation = 'production'")
	rows, err = accessorGroup.Database.Query("SELECT * FROM Rooms ")
	if err != nil {
		return []structs.Room{}, err
	}

	defer rows.Close()

	allRooms, err = accessorGroup.ExtractRoomData(rows)
	if err != nil {
		return []structs.Room{}, err
	}

	return allRooms, nil
}

func (accessorGroup *AccessorGroup) GetAllRoomDesignations() ([]string, error) {
	toReturn := []string{}

	rows, err := accessorGroup.Database.Query("SELECT DISTINCT roomDesignation FROM Rooms")
	if err != nil {
		return toReturn, err
	}

	for rows.Next() {
		var curStr string
		err = rows.Scan(&curStr)
		toReturn = append(toReturn, curStr)
	}

	return toReturn, nil
}

// GetRoomByID returns a room from the database by ID
func (accessorGroup *AccessorGroup) GetRoomByID(id int) (structs.Room, error) {
	room := &structs.Room{}

	err := accessorGroup.Database.QueryRow("SELECT * FROM Rooms WHERE roomID=?", id).Scan(&room.ID, &room.Name, &room.Building.ID, &room.Description, &room.ConfigurationID, &room.RoomDesignation)
	if err != nil {
		return structs.Room{}, err
	}

	return *room, nil
}

//ExtractRoomData pulls data from a sql query
func (accessorGroup *AccessorGroup) ExtractRoomData(rows *sql.Rows) (rooms []structs.Room, err error) {

	for rows.Next() {
		room := structs.Room{}

		err = rows.Scan(
			&room.ID,
			&room.Name,
			&room.Building.ID,
			&room.Description,
			&room.ConfigurationID,
			&room.RoomDesignation,
		)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return
		}
		rooms = append(rooms, room)
	}
	return
}

// GetRoomsByBuilding returns a room from the database by building
func (accessorGroup *AccessorGroup) GetRoomsByBuilding(building string) ([]structs.Room, error) {

	//rows, err := accessorGroup.Database.Query(`SELECT Rooms.roomID,
	//Rooms.name, Rooms.buildingID, Rooms.description, Rooms.configurationID, Rooms.roomDesignation FROM Rooms
	//JOIN Buildings ON Rooms.buildingID = Buildings.buildingID WHERE Buildings.shortName=? AND Rooms.roomDesignation = 'production'`, building)
	rows, err := accessorGroup.Database.Query(`SELECT Rooms.roomID,
	Rooms.name, Rooms.buildingID, Rooms.description, Rooms.configurationID, Rooms.roomDesignation FROM Rooms
	JOIN Buildings ON Rooms.buildingID = Buildings.buildingID WHERE Buildings.shortName=?`, building)
	if err != nil {
		return []structs.Room{}, err
	}

	defer rows.Close()

	allRooms, err := accessorGroup.ExtractRoomData(rows)
	if err != nil {
		return []structs.Room{}, err
	}
	return allRooms, nil
}

// Getstructs.RoomByBuildingAndName returns a room from the database by building shortname and room name
func (accessorGroup *AccessorGroup) GetRoomByBuildingAndName(buildingShortname string, name string) (structs.Room, error) {
	log.Printf("Getting building info for %s - %s...", buildingShortname, name)
	building, err := accessorGroup.GetBuildingByShortname(buildingShortname)
	//
	log.Printf("TEST: building.ID = %v", building.ID)
	//
	if err != nil {
		return structs.Room{}, err
	}

	room := structs.Room{}
	log.Printf("Getting room info for %s-%s...", buildingShortname, name)
	row, err := accessorGroup.Database.Query("SELECT * FROM Rooms WHERE buildingID=? AND name=?", building.ID, name)
	if err != nil {
		return structs.Room{}, err
	}
	defer row.Close()

	rooms, err := accessorGroup.ExtractRoomData(row)
	if err != nil {
		return structs.Room{}, err
	}

	if len(rooms) < 1 {
		return structs.Room{}, errors.New("No rooms found with that name.")
	}

	room = rooms[0]
	room.Building = building

	log.Printf("Getting device info for %s-%s...", buildingShortname, name)
	room.Devices, err = accessorGroup.GetDevicesByBuildingAndRoom(buildingShortname, name)
	if err != nil {
		return room, err
	}

	log.Printf("Getting configuration information for %s-%s, room key %v...", buildingShortname, name, room.ConfigurationID)
	room.Configuration, err = accessorGroup.GetConfigurationByConfigurationID(room.ConfigurationID)
	if err != nil {
		return room, err
	}

	log.Printf("Done.")
	return room, nil
}

func (accessorGroup *AccessorGroup) AddRoom(buildingShortName string, roomToAdd structs.Room) (structs.Room, error) {
	log.Printf("Adding room %v to building %v...", roomToAdd.Name, buildingShortName)

	building, err := accessorGroup.GetBuildingByShortname(buildingShortName)
	if err != nil {
		return structs.Room{}, err
	}

	result, err := accessorGroup.Database.Exec("INSERT into Rooms (name, buildingID, description, configurationID, roomDesignation) VALUES (?,?,?,?,?)",
		roomToAdd.Name, building.ID, roomToAdd.Description, roomToAdd.ConfigurationID, roomToAdd.RoomDesignation)
	if err != nil {
		return structs.Room{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.Room{}, err
	}

	roomToAdd.ID = int(id) // cast id into an int
	roomToAdd.Building = building

	return roomToAdd, nil
}
