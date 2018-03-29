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
)

func MakeRequest(method, endpoint, contentType string, body []byte, toFill interface{}) error {

	url := fmt.Sprintf("%v/%v", os.Getenv("COUCH_ADDRESS"), endpoint)

	log.L.Debugf("Making request to %v", url)

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
		log.Warn(msg)
		return errors.New(msg)
	}

	//otherwise we unmarshal
	err = json.Unmarshal(b, toFill)
	if err != nil {
		log.Infof("Couldn't umarshal response into the provided struct: %v", err.Error())

		//check to see if it was a known error from couch
		ce := CouchError{}
		err = json.Unmarshal(b, &ce)
		if err != nil {
			msg := fmt.Sprintf("Unknown response from couch: %s", b)
			log.L.Warn(msg)
			return errors.New(msg)
		}
		//it was an error, we can check on error types
		return checkCouchErrors(ce)
	}

	return nil
}

func checkCouchErrors(ce CouchError) error {
	switch ce.Error {
	case "not_found":
		return NotFound{fmt.Sprintf("The ID requested was unknown. Message: %v.", ce.Reason)}
	default:
		msg := fmt.Sprintf("Unknown error type: %v. Message: %v", ce.Error, ce.Reason)
		logger.L.Warn(msg)
		return errors.New(msg)
	}
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

type CouchError struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

type NotFound struct {
	msg string
}

func (n *NotFound) Error() string {
	return n.msg
}
