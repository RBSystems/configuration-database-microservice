package handlers

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/couch"
	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/fatih/color"
	"github.com/labstack/echo"
)

func GetRoomByBuildingAndName(context echo.Context) error {

	room, err := couch.GetRoomByID(context.Param("building") + "-" + context.Param("room"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, room)
}

func GetRoomsByBuilding(context echo.Context) error {

	room, err := couch.GetRoomsByBuilding(context.Param("buildingid"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, room)
	return nil
}

func CreateRoom(context echo.Context) error {
	toAdd := structs.Room{}
	err := context.Bind(&toAdd)
	if err != nil {
		msg := fmt.Sprintf("Invalid room format. %v", err.Error())
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}

	//now we call in
	room, err := couch.CreateRoom(toAdd)
	if err != nil {
		msg := fmt.Sprintf("Couldn't create room: %v", err.Error())
		log.L.Warn(color.HiRedString(msg))
		return context.JSON(http.StatusInternalServerError, msg)
	}

	return context.JSON(http.StatusOK, room)
}

func CreateRoomConfiguration(context echo.Context) error {
	toAdd := structs.RoomConfiguration{}
	err := context.Bind(&toAdd)
	if err != nil {
		msg := fmt.Sprintf("Invalid room configuration format. %v", err.Error())
		log.L.Warn(msg)
	}

	config, err := couch.CreateRoomConfiguration(toAdd)
	if err != nil {
		msg := fmt.Sprintf("Couldn't create room configuration: %v", err.Error())
		log.L.Warn(color.HiRedString(msg))
		return context.JSON(http.StatusInternalServerError, msg)
	}

	return context.JSON(http.StatusOK, config)
}
