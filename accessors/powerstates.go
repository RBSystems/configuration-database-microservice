package accessors

import (
	"database/sql"
	"log"
)

type PowerState struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (accessorGroup *AccessorGroup) GetPowerStates() ([]PowerState, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM PowerStates")
	if err != nil {
		return []PowerState{}, err
	}

	powerstates, err := extractPowerStates(rows)
	if err != nil {
		return []PowerState{}, err
	}

	return powerstates, nil
}

func (accessorGroup *AccessorGroup) AddPowerState(powerstate PowerState) (PowerState, error) {
	result, err := accessorGroup.Database.Exec("Insert into PowerStates (powerStateID, name, description) VALUES(?,?,?)", powerstate.ID, powerstate.Name, powerstate.Description)
	if err != nil {
		return PowerState{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return PowerState{}, err
	}

	powerstate.ID = int(id)
	return powerstate, nil
}

func extractPowerStates(rows *sql.Rows) ([]PowerState, error) {
	var powerstates []PowerState

	for rows.Next() {
		ps := PowerState{}

		err := rows.Scan(
			&ps.ID,
			&ps.Name,
			&ps.Description,
		)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return []PowerState{}, err
		}

		powerstates = append(powerstates, ps)
	}
	return powerstates, nil
}
