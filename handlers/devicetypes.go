package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetDeviceTypes(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceClasses()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddDeviceType(context echo.Context) error {
	deviceTypeName := context.Param("devicetype")
	var deviceType structs.DeviceType

	err := context.Bind(&deviceType)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if deviceTypeName != deviceType.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	response, err := handlerGroup.Accessors.AddDeviceType(deviceType)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
