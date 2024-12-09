package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// AddRole handleris method adds a new role to the system and returns the created role's id.
// @Summary Add-Role.
// @Description creates a new role and returns its id.
// @Tags Roles
// @ID add-role
// @Accept multipart/form-data
// @Produce json
// @Param roleReq formData roleModel.DTO true "role request payload"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role/create [post]
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
// @Summary Get-Role
// @Description retrieve role by identified id.
// @Tags Roles
// @ID get-role
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} roleModel.DTO
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role/{id} [get]
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
// @Summary List-Roles
// @Description list roles
// @Tags Roles
// @ID list-roles
// @Produce json
// @Success 200 {object} []roleModel.DTO
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role/get-all [get]
func (h *RoleHandler) GetRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, err := h.service.RoleService().GetRoles(c.Context())
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(200).JSON(roles)
	}
}

// DeleteRole handler method deletes a role from the system identified by the given id.
// @Summary Delete-Role
// @Description deletes role using identified role id.
// @Tags Roles
// @ID delete-role
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string "role successfully deleted"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role/{id} [delete]
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
// @Summary Update-Role
// @Description updating roles only giving rolename and do not use roleID, only can change roleName by given id.
// @Tags Roles
// @ID update-role
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Param roleReq formData roleModel.DTO true "update request for roles"
// @Success 200 {string} string "role successfully updated"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /role/{id} [put]
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
