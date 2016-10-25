package handlers

import (
	"net/http"

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
