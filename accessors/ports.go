package accessors

//PortType corresponds to the Ports table in the Database and really should be called Port
//TODO:Change struct name to "Port"
type PortType struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (accessorGroup *AccessorGroup) GetAllPorts() ([]PortType, error) {

	return []PortType{}, nil
}

//AddPort adds an entry to the Ports table in the database
func (accessorGroup *AccessorGroup) AddPort(portToAdd PortType) (PortType, error) {

	return PortType{}, nil
}

func extractPortData() ([]PortType, error) {
	return []PortType{}, nil
}
