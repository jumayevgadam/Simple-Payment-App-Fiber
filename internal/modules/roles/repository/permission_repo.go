package repository

import (
	"context"

	permissionModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// AddPayment repo method insert a new permission in db.
func (r *RoleRepository) AddPermission(ctx context.Context, res permissionModel.PermissionRes) (int, error) {
	var permissionID int

	err := r.psqlDB.QueryRow(
		ctx,
		addPermissionQuery,
		res.PermissionType,
	).Scan(&permissionID)
	if err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return permissionID, nil
}

// GetPermissionRepo method retrieve a permission from db using identified id.
func (r *RoleRepository) GetPermission(ctx context.Context, permissionID int) (*permissionModel.PermissionData, error) {
	var permissionData permissionModel.PermissionData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&permissionData,
		getPermissionQuery,
		permissionID,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &permissionData, nil
}

// ListPermissions repo method retrieves all permissions from DB.
func (r *RoleRepository) ListPermissions(ctx context.Context, paginationData abstract.PaginationData) (
	[]*permissionModel.PermissionData, error,
) {
	var permissionDatas []*permissionModel.PermissionData
	offset := (paginationData.Page - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&permissionDatas,
		listPermissionsQuery,
		offset,
		paginationData.Limit,
	)
	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return permissionDatas, nil
}

// DeletePermission repo method removes permission from db which identified by id.
func (r *RoleRepository) DeletePermission(ctx context.Context, permissionID int) error {
	_, err := r.psqlDB.Exec(
		ctx,
		deletePermissionQuery,
		permissionID,
	)
	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

// UpdatePermission repo method edits permission_type in DB.
func (r *RoleRepository) UpdatePermission(ctx context.Context, permissionID int, updateData permissionModel.PermissionRes) (
	string, error,
) {
	var response string

	err := r.psqlDB.QueryRow(
		ctx,
		updatePermissionQuery,
		updateData.PermissionType,
		permissionID,
	).Scan(&response)
	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return response, nil
}
