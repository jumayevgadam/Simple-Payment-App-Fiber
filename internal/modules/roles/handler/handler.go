package handler

import roleOps "github.com/jumayevgadam/tsu-toleg/internal/modules/roles"

// Ensuring RoleHandler implements methods of roleOps.Handlers.
var (
	_ roleOps.Handlers = (*RoleHandler)(nil)
)

// RoleHandler is for calling methods from service.
type RoleHandler struct {
	service roleOps.Service
}

// NewRoleHandler creates and returns a new instance of RoleHandler.
func NewRoleHandler(service roleOps.Service) *RoleHandler {
	return &RoleHandler{service: service}
}
