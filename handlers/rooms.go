package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/couch"
	"github.com/labstack/echo"
)

func GetRoomByBuildingAndName(context echo.Context) error {

	id, err := GetCompanyIDFromJWT(context)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	room, err := couch.GetRoomByID(id, context.Param("building")+"-"+context.Param("room"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, room)
}

func GetRoomsByBuilding(context echo.Context) error {

	id, err := GetCompanyIDFromJWT(context)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	room, err := couch.GetRoomsByBuilding(id, context.Param("buildingid"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, room)
	return nil
}
