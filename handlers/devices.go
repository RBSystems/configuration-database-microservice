package handlers

import (
	"fmt"
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/couch"
	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func GetDevicesByRoom(context echo.Context) error {
	roomID := context.Param("roomid")
	buildingID := context.Param("buildingid")
	if len(roomID) < 1 || len(buildingID) < 1 {
		return context.JSON(http.StatusBadRequest, "Need a roomid and buildingid")
	}

	devs, err := couch.GetDevicesByRoom(buildingID + "-" + roomID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, devs)
}

func GetDeviceByID(context echo.Context) error {
	room := context.Param("roomid")
	building := context.Param("buildingid")
	dev := context.Param("deviceid")
	if len(room) < 1 || len(building) < 1 || len(dev) < 1 {
		return context.JSON(http.StatusBadRequest, "Need a roomid and buildingid")
	}

	device, err := couch.GetDeviceByID(fmt.Sprintf("%v-%v-%v", building, room, dev))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, device)

}

func CreateDevice(context echo.Context) error {

	dev := structs.Device{}
	err := context.Bind(&dev)
	if err != nil {
		msg := fmt.Sprintf("Invalid device. Validae that payload reflects a valid device")
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}

	device, err := couch.CreateDevice(dev)
	if err != nil {
		msg := fmt.Sprintf("Couldn't create device. Error: %v")
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}
	return context.JSON(http.StatusOK, device)
}

func CreateDeviceType(context echo.Context) error {

	ty := structs.DeviceType{}
	err := context.Bind(&ty)
	if err != nil {
		msg := fmt.Sprintf("Invalid device type. Validae that payload reflects a valid device")
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}

	devtype, err := couch.CreateDeviceType(ty)
	if err != nil {
		msg := fmt.Sprintf("Couldn't create device type. Error: %v")
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}
	return context.JSON(http.StatusOK, devtype)
}

func GetDevicesByBuildingAndRoomAndRole(context echo.Context) error {
	room := context.Param("room")
	building := context.Param("building")
	role := context.Param("role")

	if len(room) < 2 || len(building) < 2 || len(role) < 2 {
		msg := fmt.Sprintf("Invalid parameters. Must include a valid buliding, room and role")
		log.L.Warn(msg)
		return context.JSON(http.StatusInternalServerError, msg)
	}

	devs, err := couch.GetDevicesByRoomAndRole(fmt.Sprintf("%v-%v", building, room), role)
	if err != nil {
		msg := fmt.Sprintf("error: %v", err.Error())
		log.L.Warn(msg)
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, devs)
}

func GetDevicesByRoleAndType(context echo.Context) error {
	role := context.Param("role")
	deviceType := context.Param("type")

	devs, err := couch.GetDevicesByRoleAndType(role, deviceType)
	if err != nil {
		msg := fmt.Sprintf("error: %v", err.Error())
		log.L.Warn(msg)
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, devs)
}
