package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetDeviceClasses(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceTypes()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
