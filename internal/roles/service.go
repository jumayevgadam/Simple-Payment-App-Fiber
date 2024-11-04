package roles

import (
	"context"

	roleModel "github.com/jumayevgadaym/tsu-toleg/internal/models/role"
)

// Service interface for performing role crud operations in this layer
type Service interface {
	AddRole(ctx context.Context, roleDTO *roleModel.DTO) (int, error)
	GetRole(ctx context.Context, roleID int) (*roleModel.DTO, error)
}
