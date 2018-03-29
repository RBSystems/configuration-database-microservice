package handlers

import (
	"fmt"
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

	err := context.Bind(&toAdd)
	if err != nil {
		msg := fmt.Sprintf("Invalid building, check the structure. %v", err.Error())
		log.L.Warn(msg)
		return context.JSON(http.StatusBadRequest, msg)
	}

	toAdd, err = couch.CreateBuilding(toAdd)
	if err != nil {
		log.L.Warn(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, toAdd)
}
