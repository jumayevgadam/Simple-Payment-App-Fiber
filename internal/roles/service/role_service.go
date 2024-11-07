package service

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	roleModel "github.com/jumayevgadaym/tsu-toleg/internal/models/role"
	"github.com/jumayevgadaym/tsu-toleg/internal/roles"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
)

var (
	_ roles.Service = (*RoleService)(nil)
)

// RoleService is
type RoleService struct {
	repo database.DataStore
}

// NewRoleService is
func NewRoleService(repo database.DataStore) *RoleService {
	return &RoleService{repo: repo}
}

// AddRole service is
func (s *RoleService) AddRole(ctx context.Context, roleDTO *roleModel.DTO) (int, error) {
	ctx, span := otel.Tracer("[AddRole][Service]").Start(ctx, "[AddRole]")
	defer span.End()

	roleID, err := s.repo.RolesRepo().AddRole(ctx, roleDTO.ToStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return roleID, nil
}

// GetRole Service is
func (s *RoleService) GetRole(ctx context.Context, roleID int) (*roleModel.DTO, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetRole]")
	defer span.End()

	roleDAO, err := s.repo.RolesRepo().GetRole(ctx, roleID)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.ParseErrors(err)
	}

	return roleDAO.ToServer(), nil
}

// GetRoles service is
func (s *RoleService) GetRoles(ctx context.Context) ([]*roleModel.DTO, error) {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[GetRoles]")
	defer span.End()

	rolesDAO, err := s.repo.RolesRepo().GetRoles(ctx)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return nil, errlst.ParseErrors(err)
	}

	var roleDTO []*roleModel.DTO
	for _, roleDAO := range rolesDAO {
		roleDTO = append(roleDTO, roleDAO.ToServer())
	}

	return roleDTO, nil
}

// DeleteRole service is
func (s *RoleService) DeleteRole(ctx context.Context, roleID int) error {
	ctx, span := otel.Tracer("[RoleService]").Start(ctx, "[DeleteRole]")
	defer span.End()

	if err := s.repo.RolesRepo().DeleteRole(ctx, roleID); err != nil {
		tracing.ErrorTracer(span, err)
		return errlst.ParseErrors(err)
	}

	return nil
}

// UpdateRole service is
func (s *RoleService) UpdateRole(ctx context.Context, roleDTO *roleModel.DTO) (string, error) {
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
