package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) AddPort(context echo.Context) error {
	portName := context.Param("port")
	var portToAdd structs.PortType

	err := context.Bind(&portToAdd)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if portName != portToAdd.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	response, err := handlerGroup.Accessors.AddPort(portToAdd)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetPorts(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllPorts()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetPortsByDeviceType(context echo.Context) error {

	response, err := handlerGroup.Accessors.GetPortsByDeviceTypeName(context.Param("class"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, response)
}
