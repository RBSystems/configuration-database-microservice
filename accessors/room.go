package accessors

import (
	"database/sql"
	"errors"
	"log"
)

//Room represents a room object as represented in the DB.
type Room struct {
	ID              int               `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	Building        Building          `json:"building,omitempty"`
	Devices         []Device          `json:"devices,omitempty"`
	ConfigurationID int               `json:"configurationID,omitempty"`
	Configuration   RoomConfiguration `json:"configuration"`
}

// GetAllRooms returns a list of rooms from the database
func (accessorGroup *AccessorGroup) GetAllRooms() ([]Room, error) {
	allBuildings := []Building{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Buildings")
	if err != nil {
		return []Room{}, err
	}

	for rows.Next() {
		building := Building{}

		err = rows.Scan(&building.ID, &building.Name, &building.Shortname)
		if err != nil {
			return []Room{}, err
		}

		allBuildings = append(allBuildings, building)
	}

	rows, err = accessorGroup.Database.Query("SELECT * FROM Rooms")
	if err != nil {
		return []Room{}, err
	}

	defer rows.Close()

	allRooms, err := accessorGroup.ExtractRoomData(rows)
	if err != nil {
		return []Room{}, err
	}

	err = rows.Err()
	if err != nil {
		return []Room{}, err
	}

	return allRooms, nil
}

// GetRoomByID returns a room from the database by ID
func (accessorGroup *AccessorGroup) GetRoomByID(id int) (Room, error) {
	room := &Room{}

	err := accessorGroup.Database.QueryRow("SELECT * FROM rooms WHERE id=?", id).Scan(&room.ID, &room.Name, &room.Building.ID, &room.Description)
	if err != nil {
		return Room{}, err
	}

	return *room, nil
}

//ExtractRoomData pulls data from a sql query
func (accessorGroup *AccessorGroup) ExtractRoomData(rows *sql.Rows) ([]Room, error) {

	log.Printf("Extracting data...")

	var rooms []Room

	for rows.Next() {
		var room Room
		var tableID *int
		var tableName *string
		var tableBuildingID *int
		var tableDescription *string
		var tableConfigurationID *int

		var err = rows.Scan(
			&tableID,
			&tableName,
			&tableBuildingID,
			&tableDescription,
			&tableConfigurationID,
		)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return []Room{}, err
		}

		log.Printf("Creating struct...")
		if tableID != nil {
			room.ID = *tableID
		}
		if tableName != nil {
			room.Name = *tableName
		}
		if tableBuildingID != nil {
			room.Building.ID = *tableBuildingID
		}
		if tableDescription != nil {
			room.Description = *tableDescription
		}
		if tableConfigurationID != nil {
			room.ConfigurationID = *tableConfigurationID
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

// GetRoomsByBuilding returns a room from the database by building
func (accessorGroup *AccessorGroup) GetRoomsByBuilding(building string) ([]Room, error) {

	rows, err := accessorGroup.Database.Query(`SELECT * FROM Rooms
		JOIN Buildings ON Rooms.buildingID = Buildings.buildingID WHERE buildingShortname=?`, building)
	if err != nil {
		return []Room{}, err
	}
	allRooms, err := accessorGroup.ExtractRoomData(rows)
	if err != nil {
		return []Room{}, err
	}
	return allRooms, nil
}

// GetRoomByBuildingAndName returns a room from the database by building shortname and room name
func (accessorGroup *AccessorGroup) GetRoomByBuildingAndName(buildingShortname string, name string) (Room, error) {
	log.Printf("Getting building info for %s - %s...", buildingShortname, name)
	building, err := accessorGroup.GetBuildingByShortname(buildingShortname)
	if err != nil {
		return Room{}, err
	}

	room := Room{}
	log.Printf("Getting room info for %s-%s...", buildingShortname, name)
	row, err := accessorGroup.Database.Query("SELECT * FROM Rooms WHERE buildingID=? AND name=?", building.ID, name)
	if err != nil {
		return Room{}, err
	}
	rooms, err := accessorGroup.ExtractRoomData(row)
	if err != nil {
		return Room{}, err
	}

	if len(rooms) < 1 {
		return Room{}, errors.New("No rooms found with that name.")
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
