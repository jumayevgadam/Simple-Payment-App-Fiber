package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	groupModel "github.com/jumayevgadam/tsu-toleg/internal/models/group"
	groupOps "github.com/jumayevgadam/tsu-toleg/internal/modules/groups"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
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

// AddGroup godoc.
// @Summary Add a new group.
// @Description Creates a new group and returns its id.
// @Tags Groups
// @ID add-group
// @Accept multipart/form-data
// @Produce json
// @Param groupReq formData groupModel.GroupReq true "Group request payload"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /group/add [post]
func (h *GroupHandler) AddGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[AddGroup]")
		defer span.End()

		var groupReq groupModel.GroupReq
		if err := reqvalidator.ReadRequest(c, &groupReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		groupID, err := h.service.AddGroup(ctx, &groupReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully added group")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"groupID": groupID})
	}
}

// GetGroup handler fetches group by using identified id.
// @Summary Get one group by its id.
// @Description Retrieve a group using by identified id.
// @Tags Groups
// @ID get-group
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} groupModel.GroupDTO
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /group/{id} [get]
func (h *GroupHandler) GetGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[GetGroup]")
		defer span.End()

		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		group, err := h.service.GetGroup(ctx, groupID)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully got group")
		return c.Status(fiber.StatusOK).JSON(group)
	}
}

// ListGroups handler fetches a list of groups.
// @Summary List groups.
// @Description Listing groups by pagination.
// @Tags Groups
// @ID list-group
// @Accept multipart/form-data
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param limit query int false "number of elements per page" Format(limit)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} []groupModel.GroupDTO
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /group/get-all [get]
func (h *GroupHandler) ListGroups() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[ListGroups]")
		defer span.End()

		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		groups, err := h.service.ListGroups(ctx, paginationReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully listed groups")
		return c.Status(fiber.StatusOK).JSON(groups)
	}
}

// DeleteGroup handler deletes a group using identified id.
// @Summary Delete group.
// @Description Delete group by id
// @Tags Groups
// @ID delete-group
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /group/{id} [delete]
func (h *GroupHandler) DeleteGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[DeleteGroup]")
		defer span.End()

		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		if err := h.service.DeleteGroup(ctx, groupID); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully deleted group")
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"res": fmt.Sprintf("successfully deleted group with id: %d", groupID)})
	}
}

// UpdateGroup handler update group with a new data and identified id.
// @Summary Update a group.
// @Description update a group details using their fields.
// @Tags Groups
// @ID update-group
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "needed group id for update"
// @Param groupReq formData groupModel.UpdateGroupReq true "update group request"
// @Success 200 {string} string "updated group"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /group/{id} [put]
func (h *GroupHandler) UpdateGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[UpdateGroup]")
		defer span.End()

		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var groupReq groupModel.UpdateGroupReq
		if err := reqvalidator.ReadRequest(c, &groupReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		updateRes, err := h.service.UpdateGroup(ctx, groupID, &groupReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully updated (edited) group params")
		return c.Status(fiber.StatusOK).JSON(updateRes)
	}
}
