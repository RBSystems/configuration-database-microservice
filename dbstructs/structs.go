package dbstructs

//Command represents the command table in the database. Do we need this guy?
type Command struct {
	CommandID   int    `json:"commandID,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Priority    int    `json:"priority,omitempty"`
}

//Endpoint represents the Endpoint table in the database
type Endpoint struct {
	EndpointID  int    `json:"endpointID,omitempty"`
	Name        string `json:"name,omitempty"`
	Path        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
