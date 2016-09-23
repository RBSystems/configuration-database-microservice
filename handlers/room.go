package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetAllRooms(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllRooms()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

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

func (handlerGroup *HandlerGroup) GetRoomsByBuilding(context echo.Context) error {
	room, err := strconv.Atoi(context.Param("room"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetRoomsByBuilding(room)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetRoomByBuildingAndName(context echo.Context) error {
	log.Printf("Getting room by building and name...")
	response, err := handlerGroup.Accessors.GetRoomByBuildingAndName(context.Param("building"), context.Param("room"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeRoom(context echo.Context) error {
	room := accessors.RoomRequest{}
	err := context.Bind(&room)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeRoom(room.Name, room.Building, room.VLAN)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
