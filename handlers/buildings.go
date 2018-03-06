package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/couch"
	"github.com/labstack/echo"
)

func GetAllBuildings(context echo.Context) error {

	id, err := GetCompanyIDFromJWT(context)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	buildings, err := couch.GetAllBuildings(id)
	if err != nil {
		//there's an error
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, buildings)
}

func GetBuildingByID(context echo.Context) error {

	id, err := GetCompanyIDFromJWT(context)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	buildingID := context.Param("buildingid")
	if len(buildingID) <= 0 {
		return context.JSON(http.StatusBadRequest, "No building ID")
	}

	building, err := couch.GetBuildingByID(id, buildingID)
	if err != nil {
		//there's an error
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, building)
}
