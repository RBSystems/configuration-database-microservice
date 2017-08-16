package accessors

import (
	"database/sql"

	"github.com/byuoitav/configuration-database-microservice/structs"
)

func (accessorGroup *AccessorGroup) GetPortConfiguration(building string, room string, device string) ([]structs.PortConfiguration, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM PortConfiguration")
	if err != nil {
		return []structs.PortConfiguration{}, err
	}

	portconfigurations, err := exctractPortConfigurationData(rows)
	if err != nil {
		return []structs.PortConfiguration{}, err
	}
	defer rows.Close()

	return portconfigurations, nil
}

func (accessorGroup *AccessorGroup) AddPortConfiguration(pc structs.PortConfiguration) (structs.PortConfiguration, error) {
	response, err := accessorGroup.Database.Exec("INSERT INTO PortConfiguration (portConfigurationID, destinationDeviceID, portID, sourceDeviceID, hostDeviceID) VALUES(?,?,?,?,?)", pc.ID, pc.DestinationDeviceID, pc.PortID, pc.SourceDeviceID, pc.HostDeviceID)
	if err != nil {
		return structs.PortConfiguration{}, err
	}

	id, err := response.LastInsertId()
	pc.ID = int(id)

	return pc, nil
}

func exctractPortConfigurationData(rows *sql.Rows) ([]structs.PortConfiguration, error) {
	var portconfigurations []structs.PortConfiguration
	var portconfiguration structs.PortConfiguration
	var id *int
	var ddID *int
	var pID *int
	var sdID *int
	var hID *int

	for rows.Next() {
		err := rows.Scan(&id, &ddID, &pID, &sdID, &hID)
		if err != nil {
			return []structs.PortConfiguration{}, err
		}

		if id != nil {
			portconfiguration.ID = *id
		}
		if ddID != nil {
			portconfiguration.DestinationDeviceID = *ddID
		}
		if pID != nil {
			portconfiguration.PortID = *pID
		}
		if sdID != nil {
			portconfiguration.SourceDeviceID = *sdID
		}
		if hID != nil {
			portconfiguration.HostDeviceID = *hID
		}

		portconfigurations = append(portconfigurations, portconfiguration)
	}

	err := rows.Err()
	if err != nil {
		return []structs.PortConfiguration{}, err
	}

	return portconfigurations, nil
}
