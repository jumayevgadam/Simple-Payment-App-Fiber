package service

import (
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
)

// Ensure RoleService implements the roles.Service interface.
var (
	_ roles.Service = (*RoleService)(nil)
)

// RoleService performs buisiness logic in role management.
type RoleService struct {
	repo database.DataStore
}

// NewRoleService creates and returns a new instance of RoleService.
func NewRoleService(repo database.DataStore) *RoleService {
	return &RoleService{repo: repo}
}
