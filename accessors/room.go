package accessors

import "errors"

type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Building string `json:"building"`
	VLAN     int    `json:"vlan"`
}

// GetAllRooms returns a list of rooms from the database
func (accessorGroup *AccessorGroup) GetAllRooms() ([]Room, error) {
	allRooms := []Room{}

	rows, err := accessorGroup.Database.Query("SELECT rooms.id, buildings.shortname, rooms.name, rooms.vlan FROM rooms JOIN buildings ON rooms.building=buildings.ID")
	if err != nil {
		return []Room{}, err
	}

	defer rows.Close()

	for rows.Next() {
		room := Room{}

		err := rows.Scan(&room.ID, &room.Building, &room.Name, &room.VLAN)
		if err != nil {
			return []Room{}, err
		}

		allRooms = append(allRooms, room)
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
	err := accessorGroup.Database.QueryRow("SELECT * FROM rooms WHERE id=?", id).Scan(&room.ID, &room.Name, &room.Building, &room.VLAN)

	if err != nil {
		return Room{}, err
	}

	return *room, nil
}

// GetRoomsByBuilding returns a room from the database by building
func (accessorGroup *AccessorGroup) GetRoomsByBuilding(building int) ([]Room, error) {
	allRooms := []Room{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM rooms WHERE building=?", building)
	if err != nil {
		return []Room{}, err
	}

	defer rows.Close()

	for rows.Next() {
		room := Room{}

		err := rows.Scan(&room.ID, &room.Name, &room.Building, &room.VLAN)
		if err != nil {
			return []Room{}, err
		}

		allRooms = append(allRooms, room)
	}

	err = rows.Err()
	if err != nil {
		return []Room{}, err
	}

	return allRooms, nil
}

// GetRoomByBuildingAndName returns a room from the database by building shortname and room name
func (accessorGroup *AccessorGroup) GetRoomByBuildingAndName(buildingShortname string, name string) (Room, error) {
	building, err := accessorGroup.GetBuildingByShortname(buildingShortname)
	if err != nil {
		return Room{}, err
	}

	room := &Room{}
	err = accessorGroup.Database.QueryRow("SELECT * FROM rooms WHERE building=? AND name=?", building.ID, name).Scan(&room.ID, &room.Name, &room.Building, &room.VLAN)
	if err != nil {
		return Room{}, err
	}

	return *room, nil
}

// MakeRoom adds a room to the database
func (accessorGroup *AccessorGroup) MakeRoom(name string, buildingShortname string, vlan int) (Room, error) {
	building, err := accessorGroup.GetBuildingByShortname(buildingShortname)
	if err != nil {
		return Room{}, errors.New("Could not find a building with the \"" + buildingShortname + "\" shortname")
	}

	_, err = accessorGroup.Database.Query("INSERT INTO rooms (name, building, vlan) VALUES (?, ?, ?)", name, building.ID, vlan)
	if err != nil {
		return Room{}, err
	}

	room, err := accessorGroup.GetRoomByBuildingAndName(building.Shortname, name)
	if err != nil {
		return Room{}, err
	}

	return room, nil
}
