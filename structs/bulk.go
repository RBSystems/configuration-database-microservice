package structs

type BulkResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
	Rows      []struct {
		ID    string `json:"id"`
		Key   string `json:"key"`
		Value struct {
			Rev string `json:"rev"`
		} `json:"value"`
		Doc interface{} `json:"doc"`
	} `json:"rows"`
}

type BulkBuildingResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
	Rows      []struct {
		ID    string `json:"id"`
		Key   string `json:"key"`
		Value struct {
			Rev string `json:"rev"`
		} `json:"value"`
		Doc Building `json:"doc"`
	} `json:"rows"`
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

type BulkDeviceResponse struct {
	TotalRows int `json:"total_rows"`
	Offset    int `json:"offset"`
	Rows      []struct {
		ID    string `json:"id"`
		Key   string `json:"key"`
		Value struct {
			Rev string `json:"rev"`
		} `json:"value"`
		Doc Device `json:"doc"`
	} `json:"rows"`
}
