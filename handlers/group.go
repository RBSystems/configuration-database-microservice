package handlers

import "github.com/byuoitav/configuration-database-microservice/accessors"

// HandlerGroup holds all config information for the handlers
type HandlerGroup struct {
	Accessors *accessors.AccessorGroup
}
