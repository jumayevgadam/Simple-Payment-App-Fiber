package roles

import "github.com/gofiber/fiber/v2"

// Handler interface for performing roles crud in this layer.
type Handlers interface {
	RoleOps
	PermissionOps
	RolePermissions
}

// RoleOps for handling roles in handler.
type RoleOps interface {
	AddRole() fiber.Handler
	GetRole() fiber.Handler
	GetRoles() fiber.Handler
	UpdateRole() fiber.Handler
	DeleteRole() fiber.Handler
}

// PermissionOps interface for Permissions.
type PermissionOps interface {
	AddPermission() fiber.Handler
	GetPermission() fiber.Handler
	ListPermissions() fiber.Handler
	DeletePermission() fiber.Handler
	UpdatePermission() fiber.Handler
}

// RolePermissions interface for role_permissions.
type RolePermissions interface {
	AddRolePermission() fiber.Handler
	GetPermissionsByRole() fiber.Handler
	GetRolesByPermission() fiber.Handler
}
