package health

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/byuoitav/configuration-database-microservice/handlers"
	"github.com/byuoitav/event-router-microservice/eventinfrastructure"
	"github.com/byuoitav/event-router-microservice/healthinfrastructure"
	"github.com/labstack/echo"
	"github.com/xuther/go-message-router/common"
	"github.com/xuther/go-message-router/publisher"
)

const version string = "0.9"
const port string = "7004"

var Publisher publisher.Publisher
var accessorGroup *accessors.AccessorGroup
var handlerGroup *handlers.HandlerGroup

func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, healthinfrastructure.BuildVersion(version))
}

func GetHealth() map[string]string {

	log.Printf("[HealthCheck] Checking microservice health: ")

	healthReport := make(map[string]string)

	healthReport["Web Server Status"] = "ok"
	healthReport["Initialized"] = "ok"

	vals, err := accessorGroup.GetAllBuildings()
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

func Status(context echo.Context) error {
	report := GetHealth()

	return context.JSON(http.StatusOK, report)
}

func startPublisher() {

	database := os.Getenv("CONFIGURATION_DATABASE_USERNAME") + ":" + os.Getenv("CONFIGURATION_DATABASE_PASSWORD") + "@tcp(" + os.Getenv("CONFIGURATION_DATABASE_HOST") + ":" + os.Getenv("CONFIGURATION_DATABASE_PORT") + ")/" + os.Getenv("CONFIGURATION_DATABASE_NAME")
	// Constructs a new accessor group and connects it to the database
	accessorGroup = new(accessors.AccessorGroup)
	accessorGroup.Open(database)

	// Constructs a new controller group and gives it the accessor group
	handlerGroup = new(handlers.HandlerGroup)
	handlerGroup.Accessors = accessorGroup

	Publisher, err := publisher.NewPublisher(port, 1000, 10)
	if err != nil {
		errstr := fmt.Sprintf("Could not start publisher. Error: %v\n", err.Error())
		log.Fatalf(errstr)
	}

	go func() {
		Publisher.Listen()
		if err != nil {
			errstr := fmt.Sprintf("Could not start publisher listening. Error: %v\n", err.Error())
			log.Fatalf(errstr)
		} else {
			log.Printf("Publisher started on port ", port)
			//this is where we would tell the router to listen to us

		}
	}()
}

func Publish(event eventinfrastructure.Event) {
	toSend, err := json.Marshal(&event)
	if err != nil {
		log.Printf("ERROR sending event: %v", err.Error())
	}

	header := [24]byte{}
	copy(header[:], eventinfrastructure.Health)

	log.Printf("Publishing event: %s", toSend)
	Publisher.Write(common.Message{MessageHeader: header, MessageBody: toSend})
}

func StartupCheckAndReport() {
	startPublisher()
	healthinfrastructure.SendSuccessfulStartup(GetHealth, "AV_API", Publish)
}
