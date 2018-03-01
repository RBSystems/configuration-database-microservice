package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetDeviceClasses(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceTypes()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

type SetTypeIDStruct struct {
	TypeID int `json:"type-id"`
}

func (handlerGroup *HandlerGroup) SetDeviceTypeByID(context echo.Context) error {
	var vals SetTypeIDStruct
	err := context.Bind(&vals)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	deviceIDString := context.Param("deviceID")

	deviceID, err := strconv.Atoi(deviceIDString)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	err = handlerGroup.Accessors.SetDeviceTypeByID(vals.TypeID, deviceID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	device, err := handlerGroup.Accessors.GetDeviceById(deviceID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, device)
}
