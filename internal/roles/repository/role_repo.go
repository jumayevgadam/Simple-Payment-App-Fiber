package repository

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	roleModel "github.com/jumayevgadaym/tsu-toleg/internal/models/role"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// RoleRepository is
type RoleRepository struct {
	psqlDB connection.DB
}

// NewRoleRepository is
func NewRoleRepository(psqlDB connection.DB) *RoleRepository {
	return &RoleRepository{psqlDB: psqlDB}
}

// AddRole method is
func (r *RoleRepository) AddRole(ctx context.Context, roleDAO *roleModel.DAO) (int, error) {
	var roleID int

	if err := r.psqlDB.QueryRow(
		ctx,
		addRoleQuery,
		roleDAO.RoleName,
	).Scan(&roleID); err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return roleID, nil
}

// GetRole repo is
func (r *RoleRepository) GetRole(ctx context.Context, roleID int) (*roleModel.DAO, error) {
	var roleDAO roleModel.DAO

	if err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&roleDAO,
		getRoleQuery,
		roleID,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &roleDAO, nil
}
