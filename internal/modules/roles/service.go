package roles

import (
	"context"

	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
)

// Service interface for performing role crud operations in this layer.
type Service interface {
	AddRole(ctx context.Context, roleDTO roleModel.DTO) (int, error)
	GetRole(ctx context.Context, roleID int) (roleModel.DTO, error)
	GetRoles(ctx context.Context) ([]roleModel.DTO, error)
	DeleteRole(ctx context.Context, roleID int) error
	UpdateRole(ctx context.Context, roleDTO roleModel.DTO) (string, error)
}
