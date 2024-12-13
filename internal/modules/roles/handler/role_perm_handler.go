package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	rolePermModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// AddRolePermission handler adds a new role permission.
func (h *RoleHandler) AddRolePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req rolePermModel.RolePermissionReq
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			if req.PermissionID == 0 || req.RoleID == 0 {
				return c.Status(fiber.StatusBadRequest).JSON("error: roleID or permissionID must be implement")
			}

			return errlst.Response(c, err)
		}

		res, err := h.service.RoleService().AddRolePermission(c.Context(), req)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}

// GetPermissionsByRole handler method retrieve all permissions of roles using identified role id.
func (h *RoleHandler) GetPermissionsByRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("role_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		response, err := h.service.RoleService().GetPermissionsByRole(c.Context(), roleID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
}

// GetRolesByPermission handler retrieve all roles by identified permission.
func (h *RoleHandler) GetRolesByPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissionID, err := strconv.Atoi(c.Params("permission_id"))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		response, err := h.service.RoleService().GetRolesByPermission(c.Context(), permissionID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
}

// DeleteRolePermission handler deletes role with that permission.
func (h *RoleHandler) DeleteRolePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("role_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		permissionID, err := strconv.Atoi(c.Params("permission_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.RoleService().DeleteRolePermission(c.Context(), roleID, permissionID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"message": fmt.Sprintf(
					"successfully removed role-permission with roleID: %d and permissionID: %d",
					roleID, permissionID),
			},
		)
	}
}
