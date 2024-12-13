package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	permissionModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// AddPermission method adds a new permission.
func (h *RoleHandler) AddPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req permissionModel.PermissionReq
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			return errlst.Response(c, err)
		}

		resID, err := h.service.RoleService().AddPermission(c.Context(), req)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(resID)
	}
}

// GetPermission handler retrieve permission using by identified id.
func (h *RoleHandler) GetPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissionID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		permission, err := h.service.RoleService().GetPermission(c.Context(), permissionID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(permission)
	}
}

// ListPermissions handler retrieve all permissions.
func (h *RoleHandler) ListPermissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err.Error())
		}

		permissions, err := h.service.RoleService().ListPermissions(c.Context(), paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(permissions)
	}
}

// DeletePermission handler deletes a permission using identified id.
func (h *RoleHandler) DeletePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissionID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.RoleService().DeletePermission(c.Context(), permissionID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("successfully deleted permission with ID: %d", permissionID),
		})
	}
}

// UpdatePermission handler edits permission using  by identified id.
func (h *RoleHandler) UpdatePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissionID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var updateReq permissionModel.PermissionReq
		if err := reqvalidator.ReadRequest(c, &updateReq); err != nil {
			if updateReq.PermissionType == "" {
				return c.Status(fiber.StatusOK).JSON("update structure has no value")
			}

			return errlst.Response(c, err)
		}

		res, err := h.service.RoleService().UpdatePermission(c.Context(), permissionID, updateReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}
