package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/labstack/echo"
)

//GetAllBuildings gets all buildings
func (handlerGroup *HandlerGroup) GetAllBuildings(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllBuildings()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetBuildingByID gets building by ID
func (handlerGroup *HandlerGroup) GetBuildingByID(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.GetBuildingByID(id)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

//GetBuildingByShortname gets building by shortname (i.e. ITB or HBLL)
func (handlerGroup *HandlerGroup) GetBuildingByShortname(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetBuildingByShortname(context.Param("shortname"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

// AddBuilding asdf
func (handlerGroup *HandlerGroup) AddBuilding(context echo.Context) error {
	var building structs.Building
	err := context.Bind(&building)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = handlerGroup.Accessors.GetBuildingByShortname(building.Shortname)
	if err == nil {
		return context.JSON(http.StatusBadRequest, "Building already exists in database")
	}

	building, err = handlerGroup.Accessors.AddBuilding(building.Name, building.Shortname, building.Description)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, building)
}
