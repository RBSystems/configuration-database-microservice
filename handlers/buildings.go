package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/couch"
	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

func GetAllBuildings(context echo.Context) error {

	buildings, err := couch.GetAllBuildings()
	if err != nil {
		//there's an error
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, buildings)
}

func GetBuildingByID(context echo.Context) error {
	buildingID := context.Param("buildingid")
	if len(buildingID) <= 0 {
		return context.JSON(http.StatusBadRequest, "No building ID")
	}

	building, err := couch.GetBuildingByID(buildingID)
	if err != nil {
		//there's an error
		return context.JSON(http.StatusInternalServerError, err)
	}

	return context.JSON(http.StatusOK, building)
}

func CreateBuilding(context echo.Context) error {
	toAdd := structs.Building{}

	err := context.Bind(toAdd)
	if err != nil {
		msg := "Invalid building, check the structure."
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}
	return nil
}
