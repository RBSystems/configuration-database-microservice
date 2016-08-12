package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetAllDevices(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllDevices()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetDeviceByBuildingAndRoomAndName(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceByBuildingAndRoomAndName(context.Param("building"), context.Param("room"), context.Param("device"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeDevice(context echo.Context) error {
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
