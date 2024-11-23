package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
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

// AddRole method inserts a new role into the database.
func (r *RoleRepository) AddRole(ctx context.Context, roleDAO roleModel.DAO) (int, error) {
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

// GetRole repo retrieves a role from the database by its id.
func (r *RoleRepository) GetRole(ctx context.Context, roleID int) (roleModel.DAO, error) {
	var roleDAO roleModel.DAO

	if err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&roleDAO,
		getRoleQuery,
		roleID,
	); err != nil {
		return roleModel.DAO{}, errlst.ParseSQLErrors(err)
	}

	return roleDAO, nil
}

// GetRoles repo fetches a list of all roles from database.
func (r *RoleRepository) GetRoles(ctx context.Context) ([]roleModel.DAO, error) {
	var roleDAOs []roleModel.DAO

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

// DeleteRole repo deletes role by identified id from database.
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

// FetchCurrentRoleName repo fetches role name by identified id from database.
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

// UpdateRole repo updates details of role by using id.
func (r *RoleRepository) UpdateRole(ctx context.Context, roleDAO roleModel.DAO) (string, error) {
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
