package roles

import (
	"context"

	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Service interface for performing role crud operations in this layer.
type Service interface {
	RoleServOps
	PermissionServOps
	RolePermServOps
}

// RoleServOps interface handles service methods for roles.
type RoleServOps interface {
	AddRole(ctx context.Context, roleDTO roleModel.DTO) (int, error)
	GetRole(ctx context.Context, roleID int) (roleModel.DTO, error)
	GetRoles(ctx context.Context) ([]roleModel.DTO, error)
	DeleteRole(ctx context.Context, roleID int) error
	UpdateRole(ctx context.Context, roleDTO roleModel.DTO) (string, error)
}

// PermissionServOps interface handles service methods for permissions.
type PermissionServOps interface {
	AddPermission(ctx context.Context, req roleModel.PermissionReq) (int, error)
	GetPermission(ctx context.Context, permissionID int) (*roleModel.Permission, error)
	ListPermissions(ctx context.Context, paginationOps abstract.PaginationQuery) ([]*roleModel.Permission, error)
	DeletePermission(ctx context.Context, permissionID int) error
	UpdatePermission(ctx context.Context, permissionID int, updateReq roleModel.PermissionReq) (string, error)
}

// RolePermServOps interface handles service methods for role permissions.
type RolePermServOps interface {
	AddRolePermission(ctx context.Context, req roleModel.RolePermissionReq) (string, error)
	GetPermissionsByRole(ctx context.Context, roleID int) ([]roleModel.RolePermissionReq, error)
	GetRolesByPermission(ctx context.Context, permissionID int) ([]roleModel.RolePermissionReq, error)
	DeleteRolePermission(ctx context.Context, roleID, permissionID int) error
}
