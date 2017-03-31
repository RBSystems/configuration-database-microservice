package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetDevicePowerStates(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDevicePowerStates()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddDevicePowerState(context echo.Context) error {
	dpsID := context.Param("id")
	var dps accessors.DevicePowerState

	err := context.Bind(&dps)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	// have to convert dc.ID to a string to compare it to a string (dcID)
	if dpsID != strconv.Itoa(dps.ID) {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json id must match!")
	}

	response, err := handlerGroup.Accessors.AddDevicePowerState(dps)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
