package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	groupOps "github.com/jumayevgadaym/tsu-toleg/internal/features/groups"
	groupModel "github.com/jumayevgadaym/tsu-toleg/internal/models/group"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadaym/tsu-toleg/pkg/reqvalidator"
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

// AddGroup handler creates a new group and returns id.
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
func (h *GroupHandler) ListGroups() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[ListGroups]")
		defer span.End()

		groups, err := h.service.ListGroups(ctx)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully listed groups")
		return c.Status(fiber.StatusOK).JSON(groups)
	}
}

// DeleteGroup handler deletes a group using identified id.
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
func (h *GroupHandler) UpdateGroup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[GroupHandler]").Start(c.Context(), "[UpdateGroup]")
		defer span.End()

		groupID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var groupReq groupModel.GroupDTO
		if err := reqvalidator.ReadRequest(c, &groupReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}
		groupReq.ID = groupID

		updateRes, err := h.service.UpdateGroup(ctx, groupReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully updated (edited) group params")
		return c.Status(fiber.StatusOK).JSON(updateRes)
	}
}
