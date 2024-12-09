package handler

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	roleOps "github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
)

// Ensuring RoleHandler implements methods of roleOps.Handlers.
var (
	_ roleOps.Handlers = (*RoleHandler)(nil)
)

// RoleHandler is for calling methods from service.
type RoleHandler struct {
	service services.DataService
}

// NewRoleHandler creates and returns a new instance of RoleHandler.
func NewRoleHandler(service services.DataService) *RoleHandler {
	return &RoleHandler{service: service}
}
