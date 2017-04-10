package handlers

import (
	"net/http"

	"github.com/byuoitav/configuration-database-microservice/accessors"
	"github.com/labstack/echo"
)

/*GetAllCommands simply returns a dump of the commands table in the DB.
 */
func (handlerGroup *HandlerGroup) GetAllCommands(context echo.Context) error {
	response, err := handlerGroup.Accessors.GetAllCommands()
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (handlerGroup *HandlerGroup) AddCommand(context echo.Context) error {
	cmdName := context.Param("command")
	var cmd accessors.RawCommand

	err := context.Bind(&cmd)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if cmdName != cmd.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	response, err := handlerGroup.Accessors.AddRawCommand(cmd)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, response)
	}

	return context.JSON(http.StatusOK, response)
}
