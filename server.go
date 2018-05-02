package main

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/handlers"
	"github.com/byuoitav/device-monitoring-microservice/statusinfrastructure"
	"github.com/labstack/echo"
)

func main() {
	port := ":8886"
	router := echo.New()

	router.GET("/mstatus", GetStatus)

	// get core items
	router.GET("/buildings", handlers.GetAllBuildings)
	router.GET("/buildings/:buildingid", handlers.GetBuildingByID)
	router.GET("/buildings/:buildingid/rooms", handlers.GetRoomsByBuilding)
	router.GET("/buildings/:buildingid/rooms/:roomid", handlers.GetRoomByBuildingAndName)
	router.GET("/buildings/:buildingid/rooms/:roomid/devices", handlers.GetDevicesByRoom)
	router.GET("/buildings/:buildingid/rooms/:roomid/devices/:deviceid", handlers.GetDeviceByID)

	// create core items
	router.POST("/buildings/:buildingid", handlers.CreateBuilding)
	router.POST("/buildings/:buildingid/rooms/:roomid", handlers.CreateRoom)
	router.POST("/buildings/:buildingid/rooms/:roomid/devices/:deviceid", handlers.CreateDevice)

	// create ancilliary items
	router.POST("/rooms/configurations/:configurationid", handlers.CreateRoomConfiguration)
	router.POST("/devices/types/:typeid", handlers.CreateDeviceType)

	// helper endpoints
	// router.GET("/buildings/:building/rooms/:room/devices/roles/:role", handlers.GetDevicesByBuildingAndRoomAndRole)
	// router.GET("/devices/roles/:role/types/:type", handlers.GetDevicesByRoleAndType)

	/*
		router.GET("/rooms/designations", handlers.GetAllRoomDesignations)


		router.GET("/buildings/:building/rooms/:room/configuration", handlers.GetConfigurationByRoomAndBuilding)
		router.GET("/configurations/:configuration", handlers.GetConfigurationByName)
		router.GET("/configurations", handlers.GetConfigurations)

		router.GET("/devices/ports", handlers.GetPorts)
		router.GET("/devices/types", handlers.GetDeviceTypes)
		router.GET("/devices/classes", handlers.GetDeviceClasses)
		router.GET("/devices/endpoints", handlers.GetEndpoints)
		router.GET("/devices/commands", handlers.GetAllCommands)
		router.GET("/devices/powerstates", handlers.GetPowerStates)
		router.GET("/devices/microservices", handlers.GetMicroservices)
		router.GET("/devices/roledefinitions", handlers.GetDeviceRoleDefs)

		router.GET("/classes/:class/ports", handlers.GetPortsByDeviceType)
	*/

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}

func GetStatus(context echo.Context) error {
	var s statusinfrastructure.Status
	var err error

	s.Version, err = statusinfrastructure.GetVersion("version.txt")
	if err != nil {
		return context.JSON(http.StatusOK, "Failed to open version.txt")
	}
	return context.JSON(http.StatusOK, s)
}
