package roles

import (
	"context"

	roleModel "github.com/jumayevgadaym/tsu-toleg/internal/models/role"
)

// Repository interface for performing crud ops in this layer
type Repository interface {
	AddRole(ctx context.Context, roleDAO *roleModel.DAO) (int, error)
	GetRole(ctx context.Context, roleID int) (*roleModel.DAO, error)
	GetRoles(ctx context.Context) ([]*roleModel.DAO, error)
	DeleteRole(ctx context.Context, roleID int) error
	UpdateRole(ctx context.Context, roleDAO *roleModel.DAO) (string, error)
}

// RoleOps inteface is
