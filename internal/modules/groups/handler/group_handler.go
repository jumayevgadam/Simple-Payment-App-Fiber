package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	groupOps "github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"

	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// Ensure GroupHandler implements the groupOps.Handler.
var (
	_ groupOps.Handlers = (*GroupHandler)(nil)
)

// GroupHandler performs http request actions and call methods from service.
type GroupHandler struct {
	service services.DataService
}

// NewGroupHandler creates and returns a new instance of GroupHandler.
func NewGroupHandler(service services.DataService) *GroupHandler {
	return &GroupHandler{service: service}
}

// AddGroup handler.
func (h *GroupHandler) AddGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var groupReq groupModel.Req
		if err := reqvalidator.ReadRequest(c, &groupReq); err != nil {
			return errlst.Response(c, err)
		}

		groupID, err := h.service.GroupService().AddGroup(c.Context(), &groupReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"groupID": groupID,
			"message": "successfully group created",
		})
	}
}

// GetGroup handler fetches group by using identified id.
func (h *GroupHandler) GetGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		group, err := h.service.GroupService().GetGroup(c.Context(), groupID)
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

		groups, err := h.service.GroupService().ListGroups(c.Context(), paginationReq)
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

		err = h.service.GroupService().DeleteGroup(c.Context(), groupID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusNoContent).JSON(
			fiber.Map{"res": fmt.Sprintf(
				"successfully deleted group with id: %d", groupID,
			)})
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

		updateRes, err := h.service.GroupService().UpdateGroup(c.Context(), groupID, &groupReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(updateRes)
	}
}

// ListGroupsByFacultyID handler.
func (h *GroupHandler) ListGroupsByFacultyID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facultyID, err := strconv.Atoi(c.Params("faculty_id"))
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err.Error())
		}

		groupListResponse, err := h.service.GroupService().ListGroupsByFacultyID(c.Context(), facultyID, paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(groupListResponse)
	}
}

func (h *GroupHandler) ListStudentsByGroupID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		groupIDStr := c.Query("group-id")
		if groupIDStr == "" {
			return errlst.NewBadRequestError("[groupHandler][ListStudentsByGroupID]: group-id query param can not be empty")
		}

		groupID, err := strconv.Atoi(groupIDStr)
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		paginationQuery, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		studentListByGroupID, err := h.service.UserService().ListStudentsByGroupID(c.Context(), groupID, paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(studentListByGroupID)
	}
}
