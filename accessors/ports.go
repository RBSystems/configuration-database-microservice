package accessors

import (
	"database/sql"
	"log"
)

//PortType corresponds to the Ports table in the Database and really should be called Port
//TODO:Change struct name to "Port"
type PortType struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeviceTypePort struct {
	DeviceTypePortID        int      `json:"id"`
	DeviceTypeID            int      `json:"type-id"`
	DeviceTypeName          string   `json:"type-name"`
	Port                    PortType `json:"port-info"`
	Description             string   `json:"type-port-description"`
	FriendlyName            string   `json:"friendlyName"`
	SourceDesitnationMirror bool     `json:"mirror-source-dest"`
}

func (accessorGroup *AccessorGroup) GetAllPorts() ([]PortType, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Ports")
	if err != nil {
		return []PortType{}, err
	}

	allPorts, err := extractPortData(rows)
	if err != nil {
		return []PortType{}, err
	}
	defer rows.Close()

	return allPorts, nil
}

//AddPort adds an entry to the Ports table in the database
func (accessorGroup *AccessorGroup) AddPort(portToAdd PortType) (PortType, error) {

	result, err := accessorGroup.Database.Exec("INSERT into Ports (portID, name, description) VALUES(?,?,?)", portToAdd.ID, portToAdd.Name, portToAdd.Description)
	if err != nil {
		return PortType{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return PortType{}, err
	}

	portToAdd.ID = int(id)
	return portToAdd, nil
}

func (accessorGroup *AccessorGroup) GetPortTypeByName(name string) (PortType, error) {
	row := accessorGroup.Database.QueryRow("SELECT * FROM Ports  WHERE name = ? ", name)

	p, err := extractPortType(row)
	if err != nil {
		return PortType{}, err
	}

	return p, nil
}

func (accessorGroup *AccessorGroup) GetPortsByDeviceTypeName(typeName string) ([]DeviceTypePort, error) {
	log.Printf("Getting ports for class %v", typeName)

	query :=
		`
	SELECT dtp.deviceTypePortID, dtp.portID, dtp.description, dtp.friendlyName, dtp.sourceDestinationMirror, p.name, dt.typeName, dt.deviceTypeID, p.description
	FROM DeviceTypePorts dtp
	JOIN Ports p on dtp.portID = p.portiD
	JOIN DeviceTypes dt on dt.deviceTypeID = dtp.deviceTypeID
	WHERE dt.typeName = ?
	`

	rows, err := accessorGroup.Database.Query(query, typeName)
	if err != nil {
		log.Printf("error: %v", err.Error())
		return []DeviceTypePort{}, err
	}

	log.Printf("Query executed successfully")

	val, err := extractDeviceTypePortData(rows)
	log.Printf("Found %v ports", len(val))
	return val, err
}

func extractDeviceTypePortData(rows *sql.Rows) ([]DeviceTypePort, error) {

	toReturn := []DeviceTypePort{}
	var portID *int
	var deviceTypeID *int
	var deviceTypePortID *int
	var sourceDestMirror *bool

	var deviceTypeName *string
	var deviceTypePortDescription *string
	var deviceTypePortName *string

	var portName *string
	var portDesc *string

	for rows.Next() {
		curValue := DeviceTypePort{}
		curPort := PortType{}
		err := rows.Scan(
			&deviceTypePortID,
			&portID,
			&deviceTypePortDescription,
			&deviceTypePortName,
			&sourceDestMirror,
			&portName,
			&deviceTypeName,
			&deviceTypeID,
			&portDesc)

		if err != nil {
			return []DeviceTypePort{}, err
		}

		curValue.Port = curPort

		if deviceTypePortID != nil {
			curValue.DeviceTypePortID = *deviceTypePortID
		}
		if portID != nil {
			curValue.Port.ID = *portID
		}
		if deviceTypePortDescription != nil {
			curValue.Description = *deviceTypePortDescription
		}
		if deviceTypePortName != nil {
			curValue.FriendlyName = *deviceTypePortName
		}
		if sourceDestMirror != nil {
			curValue.SourceDesitnationMirror = *sourceDestMirror
		}
		if portName != nil {
			curValue.Port.Name = *portName
		}
		if deviceTypeName != nil {
			curValue.DeviceTypeName = *deviceTypeName
		}
		if deviceTypeID != nil {
			curValue.DeviceTypeID = *deviceTypeID
		}
		if portDesc != nil {
			curValue.Port.Description = *portDesc
		}

		toReturn = append(toReturn, curValue)
	}

	return toReturn, nil
}

func extractPortData(rows *sql.Rows) ([]PortType, error) {

	var allPorts []PortType
	var port PortType
	var id *int
	var name *string
	var description *string

	for rows.Next() {
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return []PortType{}, err
		}

		if id != nil {
			port.ID = *id
		}
		if name != nil {
			port.Name = *name
		}
		if description != nil {
			port.Description = *description
		}

		allPorts = append(allPorts, port)
	}

	err := rows.Err()
	if err != nil {
		return []PortType{}, err
	}

	return allPorts, nil
}

func extractPortType(row *sql.Row) (PortType, error) {
	var p PortType
	var id *int
	var name *string
	var description *string

	err := row.Scan(&id, &name, &description)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return PortType{}, err
	}
	if id != nil {
		p.ID = *id
	}
	if name != nil {
		p.Name = *name
	}
	if description != nil {
		p.Description = *description
	}

	return p, nil
}
