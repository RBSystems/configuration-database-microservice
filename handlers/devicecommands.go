package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) AddDeviceCommand(context echo.Context) error {
	id := context.Param("id")

	var dc accessors.DeviceCommand
	err := context.Bind(&dc)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	// have to convert dc.ID to a string to compare it to a string (dcID)
	if id != strconv.Itoa(dc.ID) {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json id must match!")
	}

	response, err := handlerGroup.Accessors.AddDeviceCommand(dc)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
