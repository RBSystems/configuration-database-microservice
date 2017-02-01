package dbstructs

//Port represents the port table in the database
type Port struct {
	PortID      int    `json:"portID,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

//Command represents the command table in the database. Do we need this guy?
type Command struct {
	CommandID   int    `json:"commandID,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Priority    int    `json:"priority,omitempty"`
}

//Microservice represents the Microservice table in the Database
type Microservice struct {
	MicroserviceID int    `json:"microserviceID,omitempty"`
	Name           string `json:"name,omitempty"`
	Address        string `json:"address,omitempty"`
	Description    string `json:"description,omitempty"`
}

//Endpoint represents the Endpoint table in the database
type Endpoint struct {
	EndpointID  int    `json:"endpointID,omitempty"`
	Name        string `json:"name,omitempty"`
	Path        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
