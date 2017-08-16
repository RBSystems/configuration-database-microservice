package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetEndpoints(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllEndpoints()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddEndpoint(context echo.Context) error {
	endpointName := context.Param("endpoint")
	var endpoint structs.Endpoint

	err := context.Bind(&endpoint)
	if endpointName != endpointName {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := handlerGroup.Accessors.AddEndpoint(endpoint)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
