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
// @Summary Add Role-Permission.
// @Description adds a new role-permission twice
// @Tags RolePermissions
// @ID add-role-permission
// @Accept multipart/form-data
// @Produce json
// @Param req formData rolePermModel.RolePermissionReq true "role-permission request payload"
// @Success 200 {string} string "successfully added role-permission"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role-permission/create [post]
func (h *RoleHandler) AddRolePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req rolePermModel.RolePermissionReq
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			if req.PermissionID == 0 || req.RoleID == 0 {
				return c.Status(fiber.StatusBadRequest).JSON("error: roleID or permissionID must be implement")
			}

			return errlst.Response(c, err)
		}

		res, err := h.service.AddRolePermission(c.Context(), req)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}

// GetPermissionsByRole handler method retrieve all permissions of roles using identified role id.
// @Summary GetPermissions By Role
// @Description retrieve all permissions for that role.
// @Tags RolePermissions
// @ID get-permissions-by-role
// @Accept multipart/form-data
// @Produce json
// @Param role_id path int true "role_id"
// @Success 200 {object} rolePermModel.RolePermissionReq
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role-permission/{role_id} [get]
func (h *RoleHandler) GetPermissionsByRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("role_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		response, err := h.service.GetPermissionsByRole(c.Context(), roleID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
}

// GetRolesByPermission handler retrieve all roles by identified permission.
// @Summary GetRoles By Permissions
// @Description retrieve all roles for identified that permission.
// @Tags RolePermissions
// @ID get-roles-by-permission
// @Accept multipart/form-data
// @Produce json
// @Param permission_id path int true "permission_id"
// @Success 200 {object} rolePermModel.RolePermissionReq
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role-permission/{permission_id} [get]
func (h *RoleHandler) GetRolesByPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissionID, err := strconv.Atoi(c.Params("permission_id"))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		response, err := h.service.GetRolesByPermission(c.Context(), permissionID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
}

// DeleteRolePermission handler deletes role with that permission.
// @Summary Delete Role-Permission
// @Description delete role-permission with using role_id and permission_id.
// @Tags RolePermissions
// @ID delete-role-permission
// @Accept multipart/form-data
// @Produce json
// @Param role_id path int true "role_id"
// @Param permission_id path int true "permission_id"
// @Success 200 {string} string "successfully removed role-permission"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role-permission/{role_id}/and/{permission_id} [delete]
func (h *RoleHandler) DeleteRolePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role_id, err := strconv.Atoi(c.Params("role_id"))
		if err != nil {
			return errlst.Response(c, err)
		}
		permission_id, err := strconv.Atoi(c.Params("permission_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.DeleteRolePermission(c.Context(), role_id, permission_id)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"message": fmt.Sprintf("successfully removed role-permission with roleID: %d and permissionID: %d", role_id, permission_id),
			},
		)
	}
}
