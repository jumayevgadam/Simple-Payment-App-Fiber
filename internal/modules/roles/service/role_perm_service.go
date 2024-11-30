package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	rolePermModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// AddRolePermission creates a new role-permission for app.
func (s *RoleService) AddRolePermission(ctx context.Context, request rolePermModel.RolePermissionReq) (string, error) {
	response, err := s.repo.RolesRepo().AddRolePermission(ctx, request.ToStorage())
	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return response, nil
}

// GetRolePermissionByRole retrieve all permissions of identified role.
func (s *RoleService) GetPermissionsByRole(ctx context.Context, roleID int) ([]rolePermModel.RolePermissionReq, error) {
	var rolePermDTOs []rolePermModel.RolePermissionReq
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err := db.RolesRepo().GetRole(ctx, roleID)
		if err != nil {
			return errlst.ErrNoSuchRole
		}

		rolePermDAOs, err := db.RolesRepo().GetPermissionsByRole(ctx, roleID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		for _, daoRes := range rolePermDAOs {
			rolePermDTOs = append(rolePermDTOs, daoRes.ToServer())
		}

		return nil
	})
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return rolePermDTOs, nil
}

// GetRolesByPermission service retrieve all roles that can do that permission.
func (s *RoleService) GetRolesByPermission(ctx context.Context, permissionID int) ([]rolePermModel.RolePermissionReq, error) {
	var rolePermDTOs []rolePermModel.RolePermissionReq
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err := db.RolesRepo().GetPermission(ctx, permissionID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		rolePermDAOs, err := db.RolesRepo().GetRolesByPermissionID(ctx, permissionID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		for _, daoRes := range rolePermDAOs {
			rolePermDTOs = append(rolePermDTOs, daoRes.ToServer())
		}

		return nil
	})
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return rolePermDTOs, nil
}

// DeleteRolePermission service removes and send to handler response.
func (s *RoleService) DeleteRolePermission(ctx context.Context, roleID, permissionID int) error {
	err := s.repo.RolesRepo().DeleteRolePermission(ctx, roleID, permissionID)
	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}
