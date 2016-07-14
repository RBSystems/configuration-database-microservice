package handlers

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

func (handlerGroup *HandlerGroup) GetAllBuildings(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllBuildings()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

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

func (handlerGroup *HandlerGroup) GetBuildingByShortname(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetBuildingByShortname(context.Param("shortname"))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) MakeBuilding(context echo.Context) error {
	building := accessors.Building{}
	err := context.Bind(&building)
	if err != nil {
		return err
	}

	response, err := handlerGroup.Accessors.MakeBuilding(building.Name, building.Shortname)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}
