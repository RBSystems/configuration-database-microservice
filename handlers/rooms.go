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
	context.Bind(&toAdd)

	//now we call in
	err, room := couch.CreateRoom(toAdd)
	if err != nil {
		msg := fmt.Sprintf("Couldn't create room: %v", err.Error())
		log.L.Warn(color.HiRedString(msg))
		return context.JSON(http.StatusInternalServerError, msg)
	}

	return context.JSON(http.StatusOK, room)
}

func CreateRoomConfiguration(context echo.Context) error {

	return nil
}
