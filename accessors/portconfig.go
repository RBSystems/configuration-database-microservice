package accessors

import (
	"database/sql"
	"strconv"
)

type PortConfiguration struct {
	ID                int    `json:"id,omitempty"`
	DestinationDevice Device `json:"destination-device"`
	Port              Port   `json:"port"`
	SourceDevice      Device `json:"source-device"`
}

func (accessorGroup *AccessorGroup) GetPortConfigurations() ([]PortConfiguration, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM PortConfiguration")
	if err != nil {
		return []PortConfiguration{}, err
	}

	portconfigurations, err := exctractPortConfigurationData(rows)
	if err != nil {
		return []PortConfiguration{}, err
	}
	defer rows.Close()

	return portconfigurations, nil
}

func (accessorGroup *AccessorGroup) AddPortConfiguration(pc PortConfiguration) (PortConfiguration, error) {
	response, err := accessorGroup.Database.Exec("INSERT INTO PortConfiguration (portConfigurationID, destinationDeviceID, portID, sourceDeviceID) VALUES(?,?,?,?)", pc.ID, pc.DestinationDevice.ID, pc.ID, pc.SourceDevice.ID) // the second pc.ID should be changed to pc.Port.ID
	if err != nil {
		return PortConfiguration{}, err
	}

	id, err := response.LastInsertId()
	pc.ID = int(id)

	return pc, nil
}

func exctractPortConfigurationData(rows *sql.Rows) ([]PortConfiguration, error) {
	var portconfigurations []PortConfiguration
	var portconfiguration PortConfiguration
	var id *int
	var ddID *int
	var pID *int
	var sdID *int

	for rows.Next() {
		err := rows.Scan(&id, &ddID, &pID, &sdID)
		if err != nil {
			return []PortConfiguration{}, err
		}

		if id != nil {
			portconfiguration.ID = *id
		}
		if ddID != nil {
			portconfiguration.DestinationDevice.ID = *ddID
		}
		if pID != nil {
			portconfiguration.Port.Name = strconv.Itoa(*pID) // Port.Name should be changed to Port.ID, and strconv.Itoa() removed
		}
		if sdID != nil {
			portconfiguration.SourceDevice.ID = *sdID
		}

		portconfigurations = append(portconfigurations, portconfiguration)
	}

	err := rows.Err()
	if err != nil {
		return []PortConfiguration{}, err
	}

	return portconfigurations, nil
}
