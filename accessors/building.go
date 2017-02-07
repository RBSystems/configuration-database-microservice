package accessors

import (
	"database/sql"
	"log"
)

//Building represents a building
type Building struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Shortname   string `json:"shortname,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetAllBuildings returns a list of buildings from the database
func (accessorGroup *AccessorGroup) GetAllBuildings() ([]Building, error) {

	log.Printf("Querying database...")
	rows, err := accessorGroup.Database.Query("SELECT * FROM Buildings")
	if err != nil {
		return []Building{}, err
	}

	defer rows.Close()

	allBuildings, err := extractBuildingData(rows)
	if err != nil {
		return []Building{}, err
	}

	err = rows.Err()
	if err != nil {
		return []Building{}, err
	}

	return allBuildings, nil
}

func extractBuildingData(rows *sql.Rows) ([]Building, error) {

	log.Printf("Extracting data...")

	var allBuildings []Building

	for rows.Next() {

		var tableID *int
		var tableName *string
		var tableShortName *string
		var tableDescription *string

		var structID int
		var structName string
		var structShortName string
		var structDescription string

		err := rows.Scan(&tableID, &tableName, &tableShortName, &tableDescription)
		if err != nil {
			return []Building{}, err
		}

		if tableID != nil {
			structID = *tableID
		}
		if tableName != nil {
			structName = *tableName
		}
		if tableShortName != nil {
			structShortName = *tableShortName
		}
		if tableDescription != nil {
			structDescription = *tableDescription
		}

		log.Printf("Creating struct...")
		building := Building{
			structID,
			structName,
			structShortName,
			structDescription,
		}

		allBuildings = append(allBuildings, building)
	}

	return allBuildings, nil
}

// GetBuildingByID returns a building from the database by ID
func (accessorGroup *AccessorGroup) GetBuildingByID(id int) (Building, error) {

	row, err := accessorGroup.Database.Query("SELECT * FROM Buildings WHERE buildingID=?", id)
	if err != nil {
		return Building{}, err
	}

	defer row.Close()

	building, err := extractBuildingData(row)
	if err != nil {
		return Building{}, err
	}

	log.Printf("Done.")
	return building[0], nil
}

// GetBuildingByShortname returns a building from the database by shortname
func (accessorGroup *AccessorGroup) GetBuildingByShortname(shortname string) (Building, error) {
	row, err := accessorGroup.Database.Query("SELECT * FROM Buildings WHERE shortname=?", shortname)
	if err != nil {
		return Building{}, err
	}

	building, err := extractBuildingData(row)
	if err != nil {
		return Building{}, err
	}

	log.Printf("Done.")
	return building[0], nil
}
