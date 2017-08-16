package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetPowerStates(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetPowerStates()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddPowerState(context echo.Context) error {
	psName := context.Param("powerstate")
	var ps structs.PowerState

	err := context.Bind(&ps)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if psName != ps.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	response, err := handlerGroup.Accessors.AddPowerState(ps)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
