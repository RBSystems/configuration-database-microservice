package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetAllDevices(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllDevices()
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

func (hanlderGroup *HandlerGroup) PutDeviceAttributeByDeviceAndRoomAndBuilding(context echo.Context) error {
	values := make(map[string]string)
	context.Bind(&values)
	respose, err := handlerGroup.Accessors.PutDeviceAttributeByDeviceAndRoomAndBuilding(
		context.Param("building"),
		context.Param("room"),
		context.Param("device"),
		context.Param("attribute"),
		values["value"])
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

func (handlerGroup *HandlerGroup) MakeDevice(context echo.Context) error {
	room := accessors.RoomRequest{}
	err := context.Bind(&room)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeRoom(room.Name, room.Building, room.VLAN)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
