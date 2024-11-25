package repository

import (
	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
)

// Ensures RoleRepository implements the roles.Repository interface.
var (
	_ roles.Repository = (*RoleRepository)(nil)
)

// RoleRepository handles database operations related to roles.
type RoleRepository struct {
	psqlDB connection.DB
}

// NewRoleRepository creates and returns a new instance of RoleRepository.
func NewRoleRepository(psqlDB connection.DB) *RoleRepository {
	return &RoleRepository{psqlDB: psqlDB}
}
