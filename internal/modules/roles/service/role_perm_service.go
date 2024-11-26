package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	rolePermModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// AddRolePermission creates a new role-permission for app.
func (s *RoleService) AddRolePermission(ctx context.Context, request rolePermModel.RolePermissionReq) (string, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[AddRolePermission]")
	defer span.End()

	response, err := s.repo.RolesRepo().AddRolePermission(ctx, request.ToStorage())
	if err != nil {
		tracing.ErrorTracer(span, err)
		return "", errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully saved to role-permission to db.")
	return response, nil
}

// GetRolePermissionByRole retrieve all permissions of identified role.
func (s *RoleService) GetPermissionsByRole(ctx context.Context, roleID int) ([]rolePermModel.RolePermissionReq, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetRolePermissionByRole]")
	defer span.End()

	var rolePermDTOs []rolePermModel.RolePermissionReq
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err := db.RolesRepo().GetRole(ctx, roleID)
		if err != nil {
			tracing.ErrorTracer(span, errlst.ErrNotFound)
			return errlst.ParseErrors(err)
		}

		rolePermDAOs, err := db.RolesRepo().GetPermissionsByRole(ctx, roleID)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		for _, daoRes := range rolePermDAOs {
			rolePermDTOs = append(rolePermDTOs, daoRes.ToServer())
		}

		return nil
	})
	if err != nil {
		tracing.ErrorTracer(span, errlst.ErrTransactionFailed)
		return nil, errlst.ParseErrors(err)
	}

	return rolePermDTOs, nil
}

// GetRolesByPermission service retrieve all roles that can do that permission.
func (s *RoleService) GetRolesByPermission(ctx context.Context, permissionID int) ([]rolePermModel.RolePermissionReq, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetRolePermissionByPermission]")
	defer span.End()

	var rolePermDTOs []rolePermModel.RolePermissionReq
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err := db.RolesRepo().GetPermission(ctx, permissionID)
		if err != nil {
			tracing.ErrorTracer(span, errlst.ErrNotFound)
			return errlst.ParseErrors(err)
		}

		rolePermDAOs, err := db.RolesRepo().GetRolesByPermission(ctx, permissionID)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		for _, daoRes := range rolePermDAOs {
			rolePermDTOs = append(rolePermDTOs, daoRes.ToServer())
		}

		return nil
	})
	if err != nil {
		tracing.ErrorTracer(span, errlst.ErrTransactionFailed)
		return nil, errlst.ParseErrors(err)
	}

	return rolePermDTOs, nil
}

// DeleteRolePermission service removes and send to handler response.
func (s *RoleService) DeleteRolePermission(ctx context.Context, roleID, permissionID int) error {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[DeleteRolePermission]")
	defer span.End()

	err := s.repo.RolesRepo().DeleteRolePermission(ctx, roleID, permissionID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully removed role_permission from db.")
	return nil
}
