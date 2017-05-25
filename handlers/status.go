package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var Initialized bool

const version string = "0.9"

func (hg *HandlerGroup) Version(context echo.Context) error {
	toReturn = make(map[string]string)
	toReturn["version"] = version
	return context.JSON(http.StatusOK, toReturn)
}

//Return a map of status checks
func (hg *HandlerGroup) Status(context echo.Context) error {
	report := hg.GetHealth()

	return context.JSON(http.StatusOK, report)
}

func (hg *HandlerGroup) GetHealth() map[string]string {

	log.Printf("[HealthCheck] Checking microservice health: ")

	healthReport := make(map[string]string)

	healthReport["Web Server Status"] = "ok"

	if Initialized {
		healthReport["Initialized"] = "ok"
	} else {
		healthReport["Initialized"] = "ERROR"
	}

	vals, err := hg.Accessors.GetAllBuildings()
	if len(vals) < 1 || err != nil {
		healthReport["Database Connectivity"] = "ERROR"
	} else {
		healthReport["Database Connectivity"] = "ok"
	}

	log.Printf("[HealthCheck] Done. Report:")
	for k, v := range healthReport {
		log.Printf("%v: %v", k, v)
	}
	log.Printf("[HealthCheck] End.")

	return healthReport
}

func (hg *HandlerGroup) SendSuccessfulStartup() error {

	log.Printf("[HealthCheck] Reporting microsrevice startup complete")

	log.Printf("[HealthCheck] Checking Health...")
	statusReport := hg.GetHealth()
	allSuccess := true
	for _, v := range statusReport {
		if v != "ok" {
			allSuccess = false
		}
	}

	MicroserviceName := "ConfigurationDatabaseMicroservice"
	addr := fmt.Sprintf("%v/report/%s", os.Getenv("LOCAL_HEALTH_SERVICE_ADDRESS"), MicroserviceName)
	report := make(map[string]interface{})
	report["version"] = version
	if allSuccess {
		report["success"] = "ok"
	} else {
		report["success"] = "errors"
	}
	report["report"] = statusReport

	b, err := json.Marshal(&report)
	if err != nil {
		log.Printf("[HealthCheck] Error compiling report: %v", err.Error())
		return err
	}

	log.Printf("[HealthCheck] Reporting to %v...", addr)
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("[HealthCheck] Error sending report: %v", err.Error())
		return err
	}
	//check if it's a success code
	if resp.StatusCode != 200 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[HealthCheck] Error reading response body: %v", err.Error())
			return err
		}
		log.Printf("[HealthCheck] There was an error sending the report: %s", b)
		return errors.New(string(b))
	}

	//success
	log.Printf("[HealthCheck] Report successufully sent.")
	return nil
}
