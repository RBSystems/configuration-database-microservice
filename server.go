package main

import (
	"log"
	"os"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/byuoitav/configuration-database-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	database := os.Getenv("CONFIGURATION_DATABASE_USERNAME") + ":" + os.Getenv("CONFIGURATION_DATABASE_PASSWORD") + "@tcp(" + os.Getenv("CONFIGURATION_DATABASE_HOST") + ":" + os.Getenv("CONFIGURATION_DATABASE_PORT") + ")/" + os.Getenv("CONFIGURATION_DATABASE_NAME")

	// Constructs a new accessor group and connects it to the database
	accessorGroup := new(accessors.AccessorGroup)
	accessorGroup.Open(database)

	// Constructs a new controller group and gives it the accessor group
	handlerGroup := new(handlers.HandlerGroup)
	handlerGroup.Accessors = accessorGroup

	port := ":8006"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())

	router.Get("/health", health.Check)

	router.Get("/buildings", handlerGroup.GetAllBuildings)
	router.Get("/buildings/id/:id", handlerGroup.GetBuildingByID)
	router.Get("/buildings/shortname/:shortname", handlerGroup.GetBuildingByShortname)
	router.Get("/building/:building/room/:room", handlerGroup.GetRoomByBuildingAndName)
	router.Get("/building/:building/room/:room/device/:device", handlerGroup.GetDeviceByBuildingAndRoomAndName)

	router.Post("/buildings", handlerGroup.MakeBuilding)

	router.Get("/rooms", handlerGroup.GetAllRooms)
	router.Get("/rooms/id/:id", handlerGroup.GetRoomByID)
	router.Get("/rooms/building/:building", handlerGroup.GetRoomsByBuilding)

	router.Post("/rooms", handlerGroup.MakeRoom)

	router.Get("/devices", handlerGroup.GetAllDevices)

	log.Println("The Configuration Database microservice is listening on " + port)
	router.Run(fasthttp.New(port))
}
