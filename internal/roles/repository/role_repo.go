package repository

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	roleModel "github.com/jumayevgadaym/tsu-toleg/internal/models/role"
	"github.com/jumayevgadaym/tsu-toleg/internal/roles"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

var (
	_ roles.Repository = (*RoleRepository)(nil)
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

// GetRoles repo is
func (r *RoleRepository) GetRoles(ctx context.Context) ([]*roleModel.DAO, error) {
	var roleDAOs []*roleModel.DAO

	if err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&roleDAOs,
		getRolesQuery,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return roleDAOs, nil
}

// DeleteRole repo is
func (r *RoleRepository) DeleteRole(ctx context.Context, roleID int) error {
	_, err := r.psqlDB.Exec(
		ctx,
		deleteRoleQuery,
		roleID,
	)
	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

// FetchCurrentRoleName repo is
func (r *RoleRepository) FetchCurrentRoleName(ctx context.Context, roleDAO *roleModel.DAO) (string, error) {
	var currentRoleName string

	if err := r.psqlDB.QueryRow(
		ctx,
		fetchCurrentRoleQuery,
		roleDAO.ID,
	).Scan(&currentRoleName); err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return currentRoleName, nil
}

// UpdateRole repo is
func (r *RoleRepository) UpdateRole(ctx context.Context, roleDAO *roleModel.DAO) (string, error) {
	var res string

	if err := r.psqlDB.QueryRow(
		ctx,
		updateRoleQuery,
		roleDAO.RoleName,
		roleDAO.ID,
	).Scan(&res); err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return res, nil
}
