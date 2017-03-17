package accessors

import "database/sql"

type Building struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Shortname   string `json:"shortname,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetAllBuildings returns a list of buildings from the database
func (accessorGroup *AccessorGroup) GetAllBuildings() ([]Building, error) {

	rows, err := accessorGroup.Database.Query("SELECT * FROM Buildings")
	if err != nil {
		return []Building{}, err
	}

	defer rows.Close()

	allBuildings, err := extractBuildingData(rows)
	if err != nil {
		return []Building{}, err
	}

	return allBuildings, nil
}

// GetBuildingByID returns a building from the database by ID
func (accessorGroup *AccessorGroup) GetBuildingByID(id int) (Building, error) {
	var building Building
	var ID *int
	var name *string
	var shortname *string
	var description *string

	err := accessorGroup.Database.QueryRow("SELECT * FROM Buildings WHERE buildingID=?", id).Scan(&ID, &name, &shortname, &description)
	if err != nil {
		return Building{}, err
	}

	if ID != nil {
		building.ID = *ID
	}
	if name != nil {
		building.Name = *name
	}
	if shortname != nil {
		building.Shortname = *shortname
	}
	if description != nil {
		building.Description = *description
	}

	return building, nil
}

// GetBuildingByShortname returns a building from the database by shortname
func (accessorGroup *AccessorGroup) GetBuildingByShortname(shortname string) (Building, error) {
	var building Building
	var ID *int
	var name *string
	var description *string

	err := accessorGroup.Database.QueryRow("SELECT * FROM Buildings WHERE shortname=?", shortname).Scan(&ID, &name, &shortname, &description)
	if err != nil {
		return Building{}, err
	}

	building.Shortname = shortname
	if ID != nil {
		building.ID = *ID
	}
	if name != nil {
		building.Name = *name
	}
	if description != nil {
		building.Description = *description
	}

	return building, nil
}

func extractBuildingData(rows *sql.Rows) ([]Building, error) {

	var buildings []Building

	var id *int
	var name *string
	var shortname *string
	var description *string

	for rows.Next() {

		var building Building

		err := rows.Scan(&id, &name, &shortname, &description)
		if err != nil {
			return buildings, err
		}

		if id != nil {
			building.ID = *id
		}
		if name != nil {
			building.Name = *name
		}
		if shortname != nil {
			building.Shortname = *shortname
		}
		if description != nil {
			building.Description = *description
		}
	}

	err := rows.Err()
	if err != nil {
		return []Building{}, err
	}

	return buildings, nil
}
