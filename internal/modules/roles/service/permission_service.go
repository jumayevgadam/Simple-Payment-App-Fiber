package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	permissionModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// AddPermission service method creates a new permission and handles buisiness logic.
func (s *RoleService) AddPermission(ctx context.Context, request permissionModel.PermissionReq) (int, error) {
	resID, err := s.repo.RolesRepo().AddPermission(ctx, request.ToStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return resID, nil
}

// GetPermission service method retrieve a permission by id.
func (s *RoleService) GetPermission(ctx context.Context, permissionID int) (*permissionModel.Permission, error) {
	permissionRes, err := s.repo.RolesRepo().GetPermission(ctx, permissionID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return permissionRes.ToServer(), nil
}

// ListPermissions service method retrieve a list of permissions by using pagination.
func (s *RoleService) ListPermissions(ctx context.Context, paginationReq abstract.PaginationQuery) (
	[]*permissionModel.Permission, error,
) {
	var permissions []*permissionModel.Permission
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		permissionDAOs, err := db.RolesRepo().ListPermissions(ctx, paginationReq.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		for _, response := range permissionDAOs {
			permissions = append(permissions, response.ToServer())
		}

		return nil
	})

	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return permissions, nil
}

// DeletePermission service gets a delete response from repository and sends to handler.
func (s *RoleService) DeletePermission(ctx context.Context, permissionID int) error {
	err := s.repo.RolesRepo().DeletePermission(ctx, permissionID)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdatePermission service edits fields of permissions using identified id.
func (s *RoleService) UpdatePermission(ctx context.Context, permissionID int, updateReq permissionModel.PermissionReq) (
	string, error,
) {
	var (
		res string
		err error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// Check that permission, that exist or not in db.
		_, err := db.RolesRepo().GetPermission(ctx, permissionID)
		if err != nil {
			return errlst.ErrNotFound
		}

		res, err = db.RolesRepo().UpdatePermission(ctx, permissionID, updateReq.ToStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return res, nil
}
