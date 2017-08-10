package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetDeviceRoleDefs(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceRoleDefs()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddDeviceRoleDef(context echo.Context) error {
	drdName := context.Param("deviceroledefinition")
	var drd structs.DeviceRoleDef

	err := context.Bind(&drd)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if drdName != drd.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	response, err := handlerGroup.Accessors.AddDeviceRoleDef(drd)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
