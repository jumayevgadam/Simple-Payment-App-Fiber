package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// AddRole service processes the request to create a new role and saves it to the database.
func (s *RoleService) AddRole(ctx context.Context, roleDTO roleModel.DTO) (int, error) {
	ctx, span := otel.Tracer("[AddRole][Service]").Start(ctx, "[AddRole]")
	defer span.End()

	roleID, err := s.repo.RolesRepo().AddRole(ctx, roleDTO.ToStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return roleID, nil
}

// GetRole Service retrieves a role by its id from the database and returns it.
func (s *RoleService) GetRole(ctx context.Context, roleID int) (roleModel.DTO, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetRole]")
	defer span.End()

	var roleDTO roleModel.DTO
	roleDAO, err := s.repo.RolesRepo().GetRole(ctx, roleID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return roleModel.DTO{}, errlst.ParseErrors(err)
	}
	roleDTO = roleDAO.ToServer()

	span.SetStatus(codes.Ok, "Successfully got role")
	return roleDTO, nil
}

// GetRoles service fetches a list of all roles from database.
func (s *RoleService) GetRoles(ctx context.Context) ([]roleModel.DTO, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetRoles]")
	defer span.End()

	var roleDTOs []roleModel.DTO
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		roleDAOs, err := db.RolesRepo().GetRoles(ctx)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		for _, role := range roleDAOs {
			roleDTOs = append(roleDTOs, role.ToServer())
		}
		return nil
	})
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.ParseErrors(err)
	}

	return roleDTOs, nil
}

// DeleteRole service processes deleting role action using its id and also deletes from database.
func (s *RoleService) DeleteRole(ctx context.Context, roleID int) error {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[DeleteRole]")
	defer span.End()

	// delete from database
	if err := s.repo.RolesRepo().DeleteRole(ctx, roleID); err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	span.SetStatus(codes.Ok, "Successfully deleted role and cleared cache")
	return nil
}

// UpdateRole service processes the new role data requests and saves it to the database.
func (s *RoleService) UpdateRole(ctx context.Context, roleDTO roleModel.DTO) (string, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[UpdateRole]")
	defer span.End()
	var (
		resFromDB string
		err       error
	)

	if err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		resFromDB, err = db.RolesRepo().UpdateRole(ctx, roleDTO.ToStorage())
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		return nil
	}); err != nil {
		tracing.ErrorTracer(span, err)
		return "", errlst.ParseErrors(err)
	}

	return resFromDB, nil
}
