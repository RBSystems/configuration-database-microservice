package structs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

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

func UnmarshalFromFile(filepath string, toFill interface{}) error {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf(color.HiRedString("Could not read %v: %v", filepath, err.Error()))
		return err
	}

	err = json.Unmarshal(b, toFill)
	if err != nil {
		log.Printf(color.HiRedString("Could not unmarshal %s: %v", b, err.Error()))
	}
	return err
}
