package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) AddDeviceRole(context echo.Context) error {
	drID := context.Param("id")
	var dr structs.DeviceRole

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
