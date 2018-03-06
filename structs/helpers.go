package structs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

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
