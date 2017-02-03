package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//GetAllEndpoints returns a dump of all the endpoints in the database
func (handlerGroup *HandlerGroup) GetAllEndpoints(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetEndpoints()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
