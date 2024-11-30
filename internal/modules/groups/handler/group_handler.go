package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	groupOps "github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"

	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// Ensure GroupHandler implements the groupOps.Handler.
var (
	_ groupOps.Handler = (*GroupHandler)(nil)
)

// GroupHandler performs http request actions and call methods from service.
type GroupHandler struct {
	service groupOps.Service
}

// NewGroupHandler creates and returns a new instance of GroupHandler.
func NewGroupHandler(service groupOps.Service) *GroupHandler {
	return &GroupHandler{service: service}
}

// AddGroup handler.
func (h *GroupHandler) AddGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var groupReq groupModel.GroupReq
		if err := reqvalidator.ReadRequest(c, &groupReq); err != nil {
			return errlst.Response(c, err)
		}

		groupID, err := h.service.AddGroup(c.Context(), &groupReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"groupID": groupID})
	}
}

// GetGroup handler fetches group by using identified id.
func (h *GroupHandler) GetGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		group, err := h.service.GetGroup(c.Context(), groupID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(group)
	}
}

// ListGroups handler fetches a list of groups.
func (h *GroupHandler) ListGroups() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		groups, err := h.service.ListGroups(c.Context(), paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(groups)
	}
}

// DeleteGroup handler deletes a group using identified id.
func (h *GroupHandler) DeleteGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.DeleteGroup(c.Context(), groupID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"res": fmt.Sprintf("successfully deleted group with id: %d", groupID)})
	}
}

// UpdateGroup handler update group with a new data and identified id.
func (h *GroupHandler) UpdateGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var groupReq groupModel.UpdateGroupReq
		if err := reqvalidator.ReadRequest(c, &groupReq); err != nil {
			updateRes, err := groupReq.Validate()
			if err == nil {
				return c.Status(fiber.StatusOK).JSON(updateRes)
			}

			return errlst.Response(c, err)
		}

		updateRes, err := h.service.UpdateGroup(c.Context(), groupID, &groupReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(updateRes)
	}
}
