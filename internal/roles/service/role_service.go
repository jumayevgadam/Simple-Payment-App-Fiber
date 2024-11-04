package service

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/database"
	roleModel "github.com/jumayevgadaym/tsu-toleg/internal/models/role"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"go.opentelemetry.io/otel"
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
