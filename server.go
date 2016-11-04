package main

import (
	"log"
	"os"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/byuoitav/configuration-database-microservice/handlers"
	"github.com/byuoitav/wso2jwt"
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
	router.Use(middleware.CORS())

	router.Get("/health", health.Check)

	router.Get("/buildings", handlerGroup.GetAllBuildings)
	router.Get("/buildings/id/:id", handlerGroup.GetBuildingByID)
	router.Get("/buildings/:shortname", handlerGroup.GetBuildingByShortname)
	router.Get("/buildings/shortname/:shortname", handlerGroup.GetBuildingByShortname)
	router.Get("/buildings/:building/rooms/:room", handlerGroup.GetRoomByBuildingAndName)
	router.Get("/buildings/:building/rooms", handlerGroup.GetRoomsByBuilding)
	router.Get("/buildings/:building/rooms/:room/devices", handlerGroup.GetDevicesByBuildingAndRoom)
	router.Get("/buildings/:building/rooms/:room/devices/roles/:role", handlerGroup.GetDevicesByBuildingAndRoomAndRole)
	router.Get("/buildings/:building/rooms/:room/devices/:device", handlerGroup.GetDeviceByBuildingAndRoomAndName)

	router.Put("/buildings/:building/rooms/:room/devices/:device/attributes/:attribute/:value", handlerGroup.PutDeviceAttributeByDeviceAndRoomAndBuilding, wso2jwt.ValidateJWT())

	router.Get("/rooms", handlerGroup.GetAllRooms)
	router.Get("/rooms/id/:id", handlerGroup.GetRoomByID)
	router.Get("/rooms/buildings/:building", handlerGroup.GetRoomsByBuilding)

	router.Get("/devices/roles/:role/types/:type", handlerGroup.GetDevicesByRoleAndType)

	router.Get("/buildings/:building/rooms/:room/configuration", handlerGroup.GetConfigurationByRoomAndBuilding)
	router.Get("/configurations/:configuration", handlerGroup.GetConfigurationByName)

	log.Println("The Configuration Database microservice is listening on " + port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
