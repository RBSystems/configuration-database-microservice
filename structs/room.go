package structs

type Room struct {
	ID            string            `json:"_id"`
	Rev           string            `json:"_rev,omitempty"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Tags          []string          `json:"tags"`
	Configuration RoomConfiguration `json:"configuration"`
	Designation   string            `json:"designation"`
	Devices       []Device          `json:"devices"`
}

type RoomConfiguration struct {
	ID          string      `json:"_id"`
	Rev         string      `json:"_rev,omitempty"`
	Name        string      `json:"name"`
	Evaluators  []Evaluator `json:"evaluators"`
	Description string      `json:"description"`
}

type Evaluator struct {
	ID          string `json:"_id"`
	CodeKey     string `json:"code-key"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type RoomQueryResponse struct {
	Docs     []Room `json:"docs"`
	Bookmark string `json:"bookmark"`
	Warning  string `json:"warning"`
}

type BulkRoomResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
	Rows      []struct {
		ID    string `json:"id"`
		Key   string `json:"key"`
		Value struct {
			Rev string `json:"rev"`
		} `json:"value"`
		Doc Room `json:"doc"`
	} `json:"rows"`
}
