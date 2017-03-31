package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetDeviceRoles(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceRoles()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddDeviceRole(context echo.Context) error {
	drID := context.Param("id")
	var dr accessors.DeviceRole

	err := context.Bind(&dr)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	// have to convert dc.ID to a string to compare it to a string (dcID)
	if drID != strconv.Itoa(dr.ID) {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json id must match!")
	}

	response, err := handlerGroup.Accessors.AddDeviceRole(dr)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
