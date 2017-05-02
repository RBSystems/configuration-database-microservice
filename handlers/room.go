package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

//GetAllRooms gets all rooms
func (handlerGroup *HandlerGroup) GetAllRooms(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllRooms()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetRoomByID returns the room with a given ID.
func (handlerGroup *HandlerGroup) GetRoomByID(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetRoomByID(id)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetRoomsByBuilding returns all the rooms in a given building
func (handlerGroup *HandlerGroup) GetRoomsByBuilding(context echo.Context) error {
	building := context.Param("building")

	log.Printf("calling Accessors.GetRoomsByBuilding")
	response, err := handlerGroup.Accessors.GetRoomsByBuilding(building)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetRoomByBuildingAndName returns the room by building and name
func (handlerGroup *HandlerGroup) GetRoomByBuildingAndName(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetRoomByBuildingAndName(context.Param("building"), context.Param("room"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetConfigurationByName gets the configuration by name
func (handlerGroup *HandlerGroup) GetConfigurationByName(context echo.Context) error {
	name := context.Param("configuration")

	response, err := handlerGroup.Accessors.GetConfigurationByConfigurationName(name)

	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//
func (handlerGroup *HandlerGroup) GetConfigurations(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetConfigurations()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetConfigurationByRoomAndBuilding gets the configuration by room and building
func (handlerGroup *HandlerGroup) GetConfigurationByRoomAndBuilding(context echo.Context) error {
	building := context.Param("building")
	room := context.Param("room")

	response, err := handlerGroup.Accessors.GetConfigurationByRoomAndBuilding(building, room)

	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddRoom(context echo.Context) error {
	buildingSN := context.Param("building")
	roomN := context.Param("room")
	var roomToAdd accessors.Room
	err := context.Bind(&roomToAdd)

	if roomN != roomToAdd.Name {
		return context.JSON(http.StatusBadRequest, "Parameter and room name must match!")
	}

	roomToAdd.Name = roomN

	response, err := handlerGroup.Accessors.AddRoom(buildingSN, roomToAdd)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
