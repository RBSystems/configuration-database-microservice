package accessors

import "errors"

type room struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	VLAN     int      `json:"vlan"`
	Building Building `json:"building"`
}

type roomRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	VLAN     int    `json:"vlan"`
	Building string `json:"building,omitempty"`
}

// GetAllRooms returns a list of rooms from the database
func (accessorGroup *AccessorGroup) GetAllRooms() ([]Room, error) {
	allBuildings := []Building{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM buildings")
	if err != nil {
		return []Room{}, err
	}

	for rows.Next() {
		building := Building{}

		err := rows.Scan(&building.ID, &building.Name, &building.Shortname)
		if err != nil {
			return []Room{}, err
		}

		allBuildings = append(allBuildings, building)
	}

	allRooms := []Room{}

	rows, err = accessorGroup.Database.Query("SELECT * FROM rooms")
	if err != nil {
		return []Room{}, err
	}

	defer rows.Close()

	for rows.Next() {
		room := Room{}

		err := rows.Scan(&room.ID, &room.Name, &room.Building.ID, &room.VLAN)
		if err != nil {
			return []Room{}, err
		}

		for i := 0; i < len(allBuildings); i++ {
			if allBuildings[i].ID == room.Building.ID {
				room.Building = allBuildings[i]
				break
			}
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

	err := accessorGroup.Database.QueryRow("SELECT * FROM rooms WHERE id=?", id).Scan(&room.ID, &room.Name, &room.Building.ID, &room.VLAN)
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

		err := rows.Scan(&room.ID, &room.Name, &room.Building.ID, &room.VLAN)
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
	room.Building = building

	err = accessorGroup.Database.QueryRow("SELECT * FROM rooms WHERE building=? AND name=?", building.ID, name).Scan(&room.ID, &room.Name, &room.Building.ID, &room.VLAN)
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

	_, err = accessorGroup.Database.Exec("INSERT INTO rooms (name, building, vlan) VALUES (?, ?, ?)", name, building.ID, vlan)
	if err != nil {
		return Room{}, err
	}

	room, err := accessorGroup.GetRoomByBuildingAndName(building.Shortname, name)
	if err != nil {
		return Room{}, err
	}

	return room, nil
}
