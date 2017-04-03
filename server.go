package main

import (
	"net/http"
	"os"

	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/byuoitav/configuration-database-microservice/handlers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
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

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	secure.GET("/buildings", handlerGroup.GetAllBuildings)
	secure.GET("/buildings/id/:id", handlerGroup.GetBuildingByID)
	secure.GET("/buildings/:shortname", handlerGroup.GetBuildingByShortname)
	secure.GET("/buildings/shortname/:shortname", handlerGroup.GetBuildingByShortname)
	secure.GET("/buildings/:building/rooms/:room", handlerGroup.GetRoomByBuildingAndName)
	secure.GET("/buildings/:building/rooms", handlerGroup.GetRoomsByBuilding)
	secure.GET("/buildings/:building/rooms/:room/devices", handlerGroup.GetDevicesByBuildingAndRoom)
	secure.GET("/buildings/:building/rooms/:room/devices/roles/:role", handlerGroup.GetDevicesByBuildingAndRoomAndRole)
	secure.GET("/buildings/:building/rooms/:room/devices/:device", handlerGroup.GetDeviceByBuildingAndRoomAndName)

	secure.PUT("/buildings/:building/rooms/:room/devices/:device/attributes/:attribute/:value", handlerGroup.PutDeviceAttributeByDeviceAndRoomAndBuilding)

	secure.GET("/rooms", handlerGroup.GetAllRooms)
	secure.GET("/rooms/id/:id", handlerGroup.GetRoomByID)
	secure.GET("/rooms/buildings/:building", handlerGroup.GetRoomsByBuilding)

	secure.GET("/devices/roles/:role/types/:type", handlerGroup.GetDevicesByRoleAndType)
	secure.GET("/development/devices/roles/:role/types/:type", handlerGroup.GetDevDevicesByRoleAndType)
	secure.GET("/stage/devices/roles/:role/types/:type", handlerGroup.GetStageDevicesByRoleAndType)

	secure.GET("/buildings/:building/rooms/:room/configuration", handlerGroup.GetConfigurationByRoomAndBuilding)
	secure.GET("/configurations/:configuration", handlerGroup.GetConfigurationByName)
	secure.GET("/devices/ports", handlerGroup.GetPorts)
	secure.GET("/devices/types", handlerGroup.GetDeviceTypes)
	secure.GET("/devices/endpoints", handlerGroup.GetEndpoints)
	secure.GET("/devices/powerstates", handlerGroup.GetPowerStates)
	secure.GET("/devices/microservices", handlerGroup.GetMicroservices)
	secure.GET("/devices/roledefinitions", handlerGroup.GetDeviceRoleDefs)

	secure.POST("/buildings/:building", handlerGroup.AddBuilding)
	secure.POST("/buildings/:building/rooms/:room", handlerGroup.AddRoom)
	secure.POST("/ports/:port", handlerGroup.AddPort)
	secure.POST("/devicetypes/:devicetype", handlerGroup.AddDeviceType)
	secure.POST("/endpoints/:endpoint", handlerGroup.AddEndpoint)
	secure.POST("/powerstates/:powerstate", handlerGroup.AddPowerState)
	secure.POST("/microservices/:microservice", handlerGroup.AddMicroservice)
	secure.POST("/deviceroledefinitions/:deviceroledefinition", handlerGroup.AddDeviceRoleDef)

	secure.POST("/buildings/:building/rooms/:room/devices/:device/commands/:id", handlerGroup.AddDeviceCommand)
	secure.POST("/buildings/:building/rooms/:room/devices/:device/powerstates/:id", handlerGroup.AddDevicePowerState)
	secure.POST("/buildings/:building/rooms/:room/devices/:device/portconfiguration/:id", handlerGroup.AddPortConfiguration)
	secure.POST("/buildings/:building/rooms/:room/devices/:device/roles/:id", handlerGroup.AddDeviceRole)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
