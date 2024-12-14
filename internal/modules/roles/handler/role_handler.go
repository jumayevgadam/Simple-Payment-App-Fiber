package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	roleOps "github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// Ensuring RoleHandler implements methods of roleOps.Handlers.
var (
	_ roleOps.Handlers = (*RoleHandler)(nil)
)

// RoleHandler is for calling methods from service.
type RoleHandler struct {
	service services.DataService
}

// NewRoleHandler creates and returns a new instance of RoleHandler.
func NewRoleHandler(service services.DataService) *RoleHandler {
	return &RoleHandler{service: service}
}

// AddRole handleris method adds a new role to the system and returns the created role's id.
func (h *RoleHandler) AddRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var roleReq roleModel.DTO
		if err := reqvalidator.ReadRequest(c, &roleReq); err != nil {
			return errlst.Response(c, err)
		}

		roleID, err := h.service.RoleService().AddRole(c.Context(), roleReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"roleID": roleID,
		})
	}
}

// GetRole handler method fetches a role by its id and returns its details.
func (h *RoleHandler) GetRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		role, err := h.service.RoleService().GetRole(c.Context(), roleID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(role)
	}
}

// GetRoles handler method fetches and returns a list of all roles.
func (h *RoleHandler) GetRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, err := h.service.RoleService().GetRoles(c.Context())
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(roles)
	}
}

// DeleteRole handler method deletes a role from the system identified by the given id.
func (h *RoleHandler) DeleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		if err := h.service.RoleService().DeleteRole(c.Context(), roleID); err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusNoContent).JSON("Successfully deleted role")
	}
}

// UpdateRole handler method updates an existing role based on the provided id and new role data.
func (h *RoleHandler) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var roleReq roleModel.DTO
		if err := reqvalidator.ReadRequest(c, &roleReq); err != nil {
			if roleReq.RoleName == "" {
				return c.Status(fiber.StatusOK).JSON("update structure has no value")
			}

			return errlst.Response(c, err)
		}
		roleReq.ID = roleID

		res, err := h.service.RoleService().UpdateRole(c.Context(), roleReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}
