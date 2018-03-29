package structs

type Device struct {
	ID          string     `json:"_id"`
	Address     string     `json:"address"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	DisplayName string     `json:"display-name"`
	Type        DeviceType `json:"type"`
	Class       string     `json:"class"`
	Roles       []Role     `json:"roles"`
	Ports       []Port     `json:"ports"`
}

type DeviceType struct {
	ID          string       `json:"_id"`
	Name        string       `json:"name"`
	Class       string       `json:"class"`
	Description string       `json:"description"`
	Ports       []Port       `json:"ports"`
	PowerStates []PowerState `json:"power-states"`
	Commands    []Command    `json:"commands"`
}

type PowerState struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Port struct {
	ID                string `json:"_id"`
	FriendlyName      string `json:"friendly-name"`
	Name              string `json:"name"`
	SourceDevice      string `json:"source-device"`
	DestinationDevice string `json:"destination-device"`
	Description       string `json:"description"`
	PortType          string `json:"port-type"`
}

type Role struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Command struct {
	Microservice struct {
		ID          string `json:"_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
	} `json:"microservice"`
	Endpoint struct {
		ID          string `json:"_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Path        string `json:"path"`
	} `json:"endpoint"`
	Command struct {
		ID          string `json:"_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"command"`
}

type DeviceQueryResponse struct {
	Docs     []Device `json:"docs"`
	Bookmark string   `json:"bookmark"`
	Warning  string   `json:"warning"`
}
