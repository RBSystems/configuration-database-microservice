package couch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func MakeRequest(method, endpoint, contentType string, body []byte, toFill interface{}) error {

	url := fmt.Sprintf("%v/%v", os.Getenv("COUCH_ADDRESS"), endpoint)

	log.Printf(color.HiBlueString("Making request to %v", url))

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	//add auth
	req.SetBasicAuth(os.Getenv("COUCH_USERNAME"), os.Getenv("COUCH_PASSWORD"))

	if method != "GET" || method != "DELETE" {
		req.Header.Add("content-type", contentType)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		msg := fmt.Sprintf("Received a non-200 response from %v. Body: %s", url, b)
		log.Printf(color.HiRedString(msg))
		return errors.New(msg)
	}

	//otherwise we unmarshal
	err = json.Unmarshal(b, toFill)
	log.Printf("%s", b)
	if err != nil {
		return err
	}

	return nil
}

type IDPrefixQuery struct {
	Selector struct {
		ID struct {
			GT string `json:"$gt"`
			LT string `json:"$lt"`
		} `json:"_id"`
	} `json:"selector"`
	Limit int `json:"limit"`
}
