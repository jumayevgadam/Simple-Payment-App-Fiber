package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	roleModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	roleOps "github.com/jumayevgadam/tsu-toleg/internal/modules/roles"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	_ roleOps.Handlers = (*RoleHandler)(nil)
)

// RoleHandler is for calling methods from service.
type RoleHandler struct {
	service roleOps.Service
}

// NewRoleHandler creates and returns a new instance of RoleHandler.
func NewRoleHandler(service roleOps.Service) *RoleHandler {
	return &RoleHandler{service: service}
}

// AddRole handleris method adds a new role to the system and returns the created role's id.
func (h *RoleHandler) AddRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler][AddRole]").Start(c.Context(), "[RoleHandler]")
		defer span.End()
		var roleReq roleModel.DTO

		if err := reqvalidator.ReadRequest(c, &roleReq); err != nil {
			tracing.EventErrorTracer(span, err, "[AddRole][Handler]")
			return errlst.Response(c, err)
		}

		roleID, err := h.service.AddRole(ctx, roleReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, "[AddRole][Handler]")
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"role":   roleReq.RoleName,
			"roleID": roleID,
		})
	}
}

// GetRole handler method fetches a role by its id and returns its details.
func (h *RoleHandler) GetRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[GetRole]")
		defer span.End()

		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			tracing.EventErrorTracer(span, err, "getting error in roleID taking")
			return errlst.Response(c, err)
		}

		role, err := h.service.GetRole(ctx, roleID)
		if err != nil {
			tracing.EventErrorTracer(span, err, "service_error")
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully got role by ID")
		return c.Status(fiber.StatusOK).JSON(role)
	}
}

// GetRoles handler method fetches and returns a list of all roles.
func (h *RoleHandler) GetRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "GetRoles")
		defer span.End()

		roles, err := h.service.GetRoles(ctx)
		if err != nil {
			log.Println(err)
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully got roles")
		return c.Status(200).JSON(roles)
	}
}

// DeleteRole handler method deletes a role from the system identified by the given id.
func (h *RoleHandler) DeleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[DeleteRole]")
		defer span.End()

		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrBadQueryParams.Error())
			return errlst.Response(c, err)
		}

		if err := h.service.DeleteRole(ctx, roleID); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully deleted role")
		return c.Status(fiber.StatusNoContent).JSON("Successfully deleted role")
	}
}

// UpdateRole handler method updates an existing role based on the provided id and new role data.
func (h *RoleHandler) UpdateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[UpdateRole]")
		defer span.End()

		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrBadRequest.Error())
			return errlst.Response(c, err)
		}

		var roleReq roleModel.DTO
		if err := reqvalidator.ReadRequest(c, &roleReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}
		roleReq.ID = roleID

		res, err := h.service.UpdateRole(ctx, roleReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully updated role")
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
