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

	router.Get("/buildings", handlerGroup.GetAllBuildings, wso2jwt.ValidateJWT())
	router.Get("/buildings/id/:id", handlerGroup.GetBuildingByID, wso2jwt.ValidateJWT())
	router.Get("/buildings/shortname/:shortname", handlerGroup.GetBuildingByShortname, wso2jwt.ValidateJWT())
	router.Get("/buildings/:building/rooms/:room", handlerGroup.GetRoomByBuildingAndName, wso2jwt.ValidateJWT())
	router.Get("/buildings/:building/rooms/:room/devices", handlerGroup.GetDevicesByBuildingAndRoom, wso2jwt.ValidateJWT())
	// router.Get("/buildings/:building/rooms/:room/devices/roles", handlerGroup.GetDevicesByBuildingAndRoomAndRole, wso2jwt.ValidateJWT())
	router.Get("/buildings/:building/rooms/:room/devices/roles/:role", handlerGroup.GetDevicesByBuildingAndRoomAndRole, wso2jwt.ValidateJWT())
	router.Get("/buildings/:building/rooms/:room/devices/:device", handlerGroup.GetDeviceByBuildingAndRoomAndName, wso2jwt.ValidateJWT())

	router.Put("/buildings/:building/rooms/:room/devices/:devices/attributes/:attribute", handlerGroup.PutDeviceAttributeByDeviceAndRoomAndBuilding, wso2jwt.ValidateJWT())

	router.Post("/buildings", handlerGroup.MakeBuilding, wso2jwt.ValidateJWT())

	router.Get("/rooms", handlerGroup.GetAllRooms, wso2jwt.ValidateJWT())
	router.Get("/rooms/id/:id", handlerGroup.GetRoomByID, wso2jwt.ValidateJWT())
	router.Get("/rooms/buildings/:building", handlerGroup.GetRoomsByBuilding, wso2jwt.ValidateJWT())

	router.Post("/rooms", handlerGroup.MakeRoom, wso2jwt.ValidateJWT())

	router.Get("/devices", handlerGroup.GetAllDevices, wso2jwt.ValidateJWT())
	router.Get("/devices/roles/:role/types/:type", handlerGroup.GetDevicesByRoleAndType, wso2jwt.ValidateJWT())
	router.Post("/devices", handlerGroup.MakeDevice, wso2jwt.ValidateJWT())

	log.Println("The Configuration Database microservice is listening on " + port)
	server := fasthttp.New(port)
	server.ReadBufferSize = 1024 * 10 // Needed to interface properly with WSO2
	router.Run(server)
}
