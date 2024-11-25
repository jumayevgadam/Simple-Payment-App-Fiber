package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	permissionModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// AddPermission service method creates a new permission and handles buisiness logic.
func (s *RoleService) AddPermission(ctx context.Context, request permissionModel.PermissionReq) (int, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[AddPermission]")
	defer span.End()

	resID, err := s.repo.RolesRepo().AddPermission(ctx, request.ToStorage())
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully added a new permission, handled in service well")
	return resID, nil
}

// GetPermission service method retrieve a permission by id.
func (s *RoleService) GetPermission(ctx context.Context, permissionID int) (*permissionModel.Permission, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetPermission]")
	defer span.End()

	permissionRes, err := s.repo.RolesRepo().GetPermission(ctx, permissionID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully retrieved permission in service from repo")
	return permissionRes.ToServer(), nil
}

// ListPermissions service method retrieve a list of permissions by using pagination.
func (s *RoleService) ListPermissions(ctx context.Context, paginationReq abstract.PaginationQuery) ([]*permissionModel.Permission, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[ListPermissions]")
	defer span.End()

	var permissions []*permissionModel.Permission
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		permissionDAOs, err := db.RolesRepo().ListPermissions(ctx, paginationReq.ToStorage())
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		for _, response := range permissionDAOs {
			permissions = append(permissions, response.ToServer())
		}

		return nil
	})
	if err != nil {
		tracing.ErrorTracer(span, errlst.ErrTransactionFailed)
		return nil, errlst.ParseErrors(err)
	}

	return permissions, nil
}

// DeletePermission service gets a delete response from repository and sends to handler.
func (s *RoleService) DeletePermission(ctx context.Context, permissionID int) error {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[DeletePermission]")
	defer span.End()

	err := s.repo.RolesRepo().DeletePermission(ctx, permissionID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "successfully got delete permission response from repo")
	return nil
}

// UpdatePermission service edits fields of permissions using identified id.
func (s *RoleService) UpdatePermission(ctx context.Context, permissionID int, updateReq permissionModel.PermissionReq) (string, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[UpdatePermission]")
	defer span.End()

	var (
		res string
		err error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// Check that permission, that exist or not in db.
		_, err := db.RolesRepo().GetPermission(ctx, permissionID)
		if err != nil {
			tracing.ErrorTracer(span, errlst.ErrNotFound)
			return errlst.ParseErrors(err)
		}

		res, err = db.RolesRepo().UpdatePermission(ctx, permissionID, updateReq.ToStorage())
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		tracing.ErrorTracer(span, errlst.ErrTransactionFailed)
		return "", errlst.ParseErrors(err)
	}

	return res, nil
}
