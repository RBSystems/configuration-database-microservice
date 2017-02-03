package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//GetAllDeviceTypes returns a dump of all device types in the database
func (handlerGroup *HandlerGroup) GetAllDeviceTypes(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceTypes()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetAllPorts returns all the ports in the database
func (handlerGroup *HandlerGroup) GetAllPorts(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllPorts()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetPowerStates returns a dump of all the power StatusBadRequest
func (handlerGroup *HandlerGroup) GetPowerStates(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetPowerStates()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetDevicesByBuildingAndRoom(context echo.Context) error {

	response, err := handlerGroup.Accessors.GetDevicesByBuildingAndRoom(context.Param("building"), context.Param("room"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetDevicesByBuildingAndRoomAndRole(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDevicesByBuildingAndRoomAndRole(context.Param("building"), context.Param("room"), context.Param("role"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) PutDeviceAttributeByDeviceAndRoomAndBuilding(context echo.Context) error {
	response, err := handlerGroup.Accessors.PutDeviceAttributeByDeviceAndRoomAndBuilding(
		context.Param("building"),
		context.Param("room"),
		context.Param("device"),
		context.Param("attribute"),
		context.Param("value"))

	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetDeviceByBuildingAndRoomAndName(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDeviceByBuildingAndRoomAndName(context.Param("building"), context.Param("room"), context.Param("device"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetDevicesByRoleAndType(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDevicesByRoleAndType(context.Param("role"), context.Param("type"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
