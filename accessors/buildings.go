package accessors

import "github.com/byuoitav/configuration-database-microservice/structs"

// GetAllBuildings returns a list of buildings from the database
func (accessorGroup *AccessorGroup) GetAllBuildings() ([]structs.Building, error) {
	allBuildings := []structs.Building{}

	rows, err := accessorGroup.Database.Query("SELECT * FROM Buildings")
	if err != nil {
		return []structs.Building{}, err
	}

	defer rows.Close()

	for rows.Next() {
		building := structs.Building{}

		err = rows.Scan(&building.ID, &building.Name, &building.Shortname, &building.Description)
		if err != nil {
			return []structs.Building{}, err
		}

		allBuildings = append(allBuildings, building)
	}

	err = rows.Err()
	if err != nil {
		return []structs.Building{}, err
	}

	return allBuildings, nil
}

// GetBuildingByID returns a building from the database by ID
func (accessorGroup *AccessorGroup) GetBuildingByID(id int) (structs.Building, error) {
	building := &structs.Building{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Buildings WHERE buildingID=?", id).Scan(&building.ID, &building.Name, &building.Shortname, &building.Description)
	if err != nil {
		return structs.Building{}, err
	}

	return *building, nil
}

// GetBuildingByShortname returns a building from the database by shortname
func (accessorGroup *AccessorGroup) GetBuildingByShortname(shortname string) (structs.Building, error) {
	building := &structs.Building{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Buildings WHERE shortname=?", shortname).Scan(&building.ID, &building.Name, &building.Shortname, &building.Description)
	if err != nil {
		return structs.Building{}, err
	}

	return *building, nil
}

//AddBuilding adds a building

func (accessorGroup *AccessorGroup) AddBuilding(name string, shortname string, description string) (structs.Building, error) {

	result, err := accessorGroup.Database.Exec(`INSERT into Buildings (name, shortname, description) VALUES (?,?,?)`, name, shortname, description)
	if err != nil {
		return structs.Building{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.Building{}, err
	}

	building := structs.Building{
		Name:        name,
		Shortname:   shortname,
		Description: description,
	}
	building.ID = int(id) // cast id into an int

	return building, err
}
