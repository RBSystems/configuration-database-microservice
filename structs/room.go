package structs

type Room struct {
	ID            string   `json:"__id"`
	Name          string   `json:"name"`
	Shortname     string   `json:"shortname"`
	Description   string   `json:"description"`
	Tags          []string `json:"tags"`
	Configuration string   `json:"configuration"`
	Designation   string   `json:"designation"`
	Devices       []Device `json:"devices"`
}

type RoomConfiguration struct {
	ID          string      `json:"_id"`
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
