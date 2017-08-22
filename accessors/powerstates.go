package accessors

import (
	"database/sql"
	"log"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

func (accessorGroup *AccessorGroup) GetPowerStates() ([]structs.PowerState, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM PowerStates")
	if err != nil {
		return []structs.PowerState{}, err
	}

	powerstates, err := extractPowerStates(rows)
	if err != nil {
		return []structs.PowerState{}, err
	}
	defer rows.Close()

	return powerstates, nil
}

func (accessorGroup *AccessorGroup) AddPowerState(powerstate structs.PowerState) (structs.PowerState, error) {
	result, err := accessorGroup.Database.Exec("Insert into PowerStates (powerStateID, name, description) VALUES(?,?,?)", powerstate.ID, powerstate.Name, powerstate.Description)
	if err != nil {
		return structs.PowerState{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return structs.PowerState{}, err
	}

	powerstate.ID = int(id)
	return powerstate, nil
}

func (accessorGroup *AccessorGroup) GetPowerStateByID(id int) (structs.PowerState, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM PowerStates WHERE powerStateID = ?", id)

	ps, err := extractPowerState(row)
	if err != nil {
		return structs.PowerState{}, err
	}

	return ps, nil
}

func (accessorGroup *AccessorGroup) GetPowerStateByName(name string) (structs.PowerState, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM PowerStates WHERE name = ?", name)

	ps, err := extractPowerState(row)
	if err != nil {
		return structs.PowerState{}, err
	}

	return ps, nil
}

func extractPowerStates(rows *sql.Rows) ([]structs.PowerState, error) {
	var powerstates []structs.PowerState
	var ps structs.PowerState
	var id *int
	var name *string
	var description *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return []structs.PowerState{}, err
		}
		if id != nil {
			ps.ID = *id
		}
		if name != nil {
			ps.Name = *name
		}
		if description != nil {
			ps.Description = *description
		}

		powerstates = append(powerstates, ps)
	}
	return powerstates, nil
}

func extractPowerState(row *sql.Row) (structs.PowerState, error) {
	var ps structs.PowerState
	var id *int
	var name *string
	var description *string

	err := row.Scan(&id, &name, &description)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return structs.PowerState{}, err
	}
	if id != nil {
		ps.ID = *id
	}
	if name != nil {
		ps.Name = *name
	}
	if description != nil {
		ps.Description = *description
	}

	return ps, nil
}
