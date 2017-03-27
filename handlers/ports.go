package handlers

import "net/http"

func (handlerGroup *HandlerGroup) AddPort(context echo.context) error {
	var portToAdd PortType
	portName = context.Param("port")
	err := context.Bind(portToAdd)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}
	if portName != portToAdd.Name {
		return context.JSON(http.StatusBadRequest, "Endpoint parameter and json name must match!")
	}

	portToAdd, err = handlerGroup.accessorGroup.AddPort(portToAdd)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, nil)
}

func (handlerGroup *HandlerGroup) GetAllPorts(context echo.context) error {

	allPorts, err := handlerGroup.Accessors.GetAllPorts()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, allPorts)
}
