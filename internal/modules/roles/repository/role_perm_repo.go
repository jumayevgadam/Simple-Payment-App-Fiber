package repository

import (
	"context"

	rolePermModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// AddRolePermission repo inserts a new role-permission to db.
func (r *RoleRepository) AddRolePermission(ctx context.Context, data rolePermModel.RolePermissionRes) (string, error) {
	var res string

	err := r.psqlDB.QueryRow(
		ctx,
		addRolePermissionQuery,
		data.RoleID,
		data.PermissionID,
	).Scan(&res)
	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return res, nil
}

// GetRolePermissionByRole repo retrieve all permissions of identified role.
func (r *RoleRepository) GetPermissionsByRole(ctx context.Context, roleID int) ([]rolePermModel.RolePermissionRes, error) {
	var res []rolePermModel.RolePermissionRes

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&res,
		getPermissionsByRoleQuery,
		roleID,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return res, nil
}

// GetRolesByPermission repo retrieve all roles of identifided permission.
func (r *RoleRepository) GetRolesByPermissionID(ctx context.Context, permissionID int) ([]rolePermModel.RolePermissionRes, error) {
	var res []rolePermModel.RolePermissionRes

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&res,
		getRolesByPermissionIDQuery,
		permissionID,
	)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return res, nil
}

// GetRolesByPermission repo takes roles from identified permission,we will use this in role based middleware.
func (r *RoleRepository) GetRolesByPermission(ctx context.Context, permissionType string) ([]rolePermModel.DAO, error) {
	var roles []rolePermModel.DAO

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&roles,
		getRolesByPermissionQuery,
		permissionType,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return roles, nil
}

// DeleteRolePermission repo removes role permission by identified role_id and permission_id.
func (r *RoleRepository) DeleteRolePermission(ctx context.Context, roleID, permissionID int) error {
	_, err := r.psqlDB.Exec(
		ctx,
		deleteRolePermissionQuery,
		roleID,
		permissionID,
	)
	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

func (r *RoleRepository) GetPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error) {
	var permissions []string

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&permissions,
		getPermissionsByRoleIDQuery,
		roleID,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return permissions, nil
}
