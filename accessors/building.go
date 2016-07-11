package accessors

type Building struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
}

// GetBuildingByEmail returns a building from the database by email
func (accessorGroup *AccessorGroup) GetBuildingByEmail(email string) (Building, error) {
	building := &Building{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM buildings WHERE email=?", email).Scan(&building)

	if err != nil {
		return Building{}, err
	}

	return *building, nil
}

// GetBuildingByID returns a building from the database by buildingID
func (accessorGroup *AccessorGroup) GetBuildingByID(email string) (Building, error) {
	building := &Building{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM buildings WHERE id=?", email).Scan(&building)

	if err != nil {
		return Building{}, err
	}

	return *building, nil
}

// GetBuildingID returns a building from the database by buildingID
func (accessorGroup *AccessorGroup) GetBuildingID(email string) (int, error) {
	building, err := accessorGroup.GetBuildingByEmail(email)
	if err != nil {
		return 0, err
	}

	return building.ID, nil
}

// MakeBuilding adds a building to the database
func (accessorGroup *AccessorGroup) MakeBuilding(email string) (Building, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO buildings (email) VALUES (?)", email)
	if err != nil {
		return Building{}, err
	}

	building, err := accessorGroup.GetBuildingByEmail(email)
	if err != nil {
		return Building{}, err
	}

	return building, nil
}
