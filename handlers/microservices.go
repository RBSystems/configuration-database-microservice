package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetMicroservices(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetMicroservices()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddMicroservice(context echo.Context) error {
	msName := context.Param("microservice")
	var ms accessors.Microservice

	err := context.Bind(&ms)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if msName != ms.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	response, err := handlerGroup.Accessors.AddMicroservice(ms)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
