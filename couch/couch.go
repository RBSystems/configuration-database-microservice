package couch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/byuoitav/configuration-database-microservice/log"
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

	if method != "GET" && method != "DELETE" {
		req.Header.Add("content-type", contentType)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {

		log.L.Info("Got a non-200 response from: %v. Code: %v", endpoint, resp.StatusCode)

		ce := CouchError{}
		err = json.Unmarshal(b, &ce)
		if err != nil {
			msg := fmt.Sprintf("Received a non-200 response from %v. Body: %s", url, b)
			log.L.Warn(msg)
			return errors.New(msg)
		}
		return checkCouchErrors(ce)
	}

	if toFill == nil {
		return nil
	}

	//otherwise we unmarshal
	err = json.Unmarshal(b, toFill)
	if err != nil {
		log.L.Infof("Couldn't umarshal response into the provided struct: %v", err.Error())

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
	log.L.Debugf("Checking for couch error type: %v", ce.Error)
	switch strings.ToLower(ce.Error) {
	case "not_found":
		log.L.Debug("Error type found: Not Found.")
		return &NotFound{fmt.Sprintf("The ID requested was unknown. Message: %v.", ce.Reason)}
	case "conflict":
		log.L.Debug("Error type found: Conflict.")
		return &Confict{fmt.Sprintf("There was a conflict updating/creating the document: %v", ce.Reason)}
	case "bad_request":
		log.L.Debug("Error type found: Bad Request.")
		return &BadRequest{fmt.Sprintf("The request was bad: %v", ce.Reason)}
	default:
		msg := fmt.Sprintf("Unknown error type: %v. Message: %v", ce.Error, ce.Reason)
		log.L.Warn(msg)
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

type CouchUpsertResponse struct {
	OK  bool   `json:"ok"`
	ID  string `json:"id"`
	Rev string `json:"rev"`
}

type CouchError struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

type NotFound struct {
	msg string
}

func (n NotFound) Error() string {
	return n.msg
}

type Confict struct {
	msg string
}

func (c Confict) Error() string {
	return c.msg
}

type BadRequest struct {
	msg string
}

func (br BadRequest) Error() string {
	return br.msg
}
