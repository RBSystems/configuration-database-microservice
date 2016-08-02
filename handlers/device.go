package handlers

import (
	"net/http"

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
	response, err := handlerGroup.Accessors.GetDeviceByBuildingAndRoomAndName()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
