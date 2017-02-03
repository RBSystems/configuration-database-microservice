package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//GetAllMicroservices returns a dump of all the microservices in the database
func (handlerGroup *HandlerGroup) GetAllMicroservices(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllMicroservices()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
