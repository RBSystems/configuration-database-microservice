package handlers

import (
	"log"
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) SetDeviceAttribute(context echo.Context) error {
	var info accessors.DeviceAttributeInfo

	err := context.Bind(&info)
	if err != nil {
		return err
	}

	device, err := handlerGroup.Accessors.SetDeviceAttribute(info)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, device)
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
	response, err := handlerGroup.Accessors.GetDevicesByRoleAndType(context.Param("role"), context.Param("type"), "production")
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetStageDevicesByRoleAndType(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetDevicesByRoleAndType(context.Param("role"), context.Param("type"), "stage")
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) GetBranchDevicesByRoleAndType(context echo.Context) error {
	branch := context.Param("branch")
	log.Printf("Getting %v devices by role %s and type %s", branch, context.Param("role"), context.Param("type"))
	response, err := handlerGroup.Accessors.GetDevicesByRoleAndType(context.Param("role"), context.Param("type"), branch)
	if err != nil {
		log.Printf("[error] %s", err.Error())
		return context.String(http.StatusBadRequest, err.Error())
	}

	log.Printf("Response: %s", response)
	return context.JSON(http.StatusOK, response)

}

func (handlerGroup *HandlerGroup) AddDevice(context echo.Context) error {
	buildingSN := context.Param("building")
	roomN := context.Param("room")
	dN := context.Param("device")
	var d accessors.Device

	err := context.Bind(&d)

	if dN != d.Name {
		return context.JSON(http.StatusBadRequest, "Parameter and device name must match!")
	}

	building, err := handlerGroup.Accessors.GetBuildingByShortname(buildingSN)
	if err != nil {
		return err
	}
	d.Building = building

	room, err := handlerGroup.Accessors.GetRoomByBuildingAndName(buildingSN, roomN)
	if err != nil {
		return err
	}
	d.Room = room

	response, err := handlerGroup.Accessors.AddDevice(d)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
