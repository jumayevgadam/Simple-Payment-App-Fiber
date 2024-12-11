package roles

import (
	"context"

	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
)

// Repository interface for performing crud ops in this layer.
type Repository interface {
	RoleRepoOps
	PermissionRepoOps
	RolePermRepoOps // Role Permissions repository ops.
}

// RoleOps interface handles repo methods for roles.
type RoleRepoOps interface {
	AddRole(ctx context.Context, roleDAO roleModel.DAO) (int, error)
	GetRole(ctx context.Context, roleID int) (roleModel.DAO, error)
	GetRoleByRoleName(ctx context.Context, role string) (roleModel.DAO, error)
	GetRoles(ctx context.Context) ([]roleModel.DAO, error)
	DeleteRole(ctx context.Context, roleID int) error
	UpdateRole(ctx context.Context, roleDAO roleModel.DAO) (string, error)
}

// PermissionRepoOps interface handles repo methods for permissions.
type PermissionRepoOps interface {
	AddPermission(ctx context.Context, res roleModel.PermissionRes) (int, error)
	GetPermission(ctx context.Context, permissionID int) (*roleModel.PermissionData, error)
	ListPermissions(ctx context.Context, paginationOps abstract.PaginationData) ([]*roleModel.PermissionData, error)
	DeletePermission(ctx context.Context, permissionID int) error
	UpdatePermission(ctx context.Context, permissionID int, updateRes roleModel.PermissionRes) (string, error)
}

// RolePermRepoOps interface handles repo methods for role_permissions.
type RolePermRepoOps interface {
	AddRolePermission(ctx context.Context, data roleModel.RolePermissionRes) (string, error)
	GetPermissionsByRole(ctx context.Context, roleID int) ([]roleModel.RolePermissionRes, error)
	GetRolesByPermissionID(ctx context.Context, permissionID int) ([]roleModel.RolePermissionRes, error)
	GetRolesByPermission(ctx context.Context, permissionType string) ([]roleModel.DAO, error) // for middleware need.
	DeleteRolePermission(ctx context.Context, roleID, permissionID int) error
	GetPermissionsByRoleID(ctx context.Context, roleID int) ([]string, error)
}
