package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	permissionModel "github.com/jumayevgadam/tsu-toleg/internal/models/role"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// AddPermission method adds a new permission.
// @Summary AddPermission.
// @Description add a new permission for roles.
// @Tags Permissions
// @ID add-permission
// @Accept multipart/form-data
// @Produce json
// @Param req formData permissionModel.PermissionReq true "permission request payload"
// @Success 200 {integer} integer
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /permission/add [post]
func (h *RoleHandler) AddPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[AddPermission]")
		defer span.End()

		var req permissionModel.PermissionReq
		if err := reqvalidator.ReadRequest(c, &req); err != nil {
			return errlst.Response(c, err)
		}

		resID, err := h.service.AddPermission(ctx, req)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully added permission")
		return c.Status(fiber.StatusOK).JSON(resID)
	}
}

// GetPermission handler retrieve permission using by identified id.
// @Summary Get-PermissionBYID
// @Description retrieve permission by using identified id.
// @Tags Permissions
// @ID get-permission
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} permissionModel.Permission
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /permission/{id} [get]
func (h *RoleHandler) GetPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[GetPermission]")
		defer span.End()

		permissionID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		permission, err := h.service.GetPermission(ctx, permissionID)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully got permission")
		return c.Status(fiber.StatusOK).JSON(permission)
	}
}

// ListPermissions handler retrieve all permissions.
// @Summary List-Permissions
// @Description list permissions by pagination
// @Tags Permissions
// @ID list-permission
// @Accept multipart/form-data
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param limit query int false "number of objects per page" Format(limit)
// @Param orderBy query string false "filter name" Format(orderBy)
// @Success 200 {object} []permissionModel.Permission
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /permission/list-all [get]
func (h *RoleHandler) ListPermissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[ListPermissions]")
		defer span.End()

		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		permissions, err := h.service.ListPermissions(ctx, paginationReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully listed permissions")
		return c.Status(fiber.StatusOK).JSON(permissions)
	}
}

// DeletePermission handler deletes a permission using identified id.
// @Summary Delete-Permission.
// @Description deletes a permission using identified id.
// @Tags Permissions
// @ID delete-permission
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string "permission deleted successfully"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /permission/{id} [delete]
func (h *RoleHandler) DeletePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[DeletePermission]")
		defer span.End()

		permissionID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.DeletePermission(ctx, permissionID)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully deleted permission")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("successfully deleted permission with ID: %d", permissionID),
		})
	}
}

// UpdatePermission handler edits permission using  by identified id.
// @Summary Update-Permission
// @Description editing a permissions by using id.
// @Tags Permissions
// @ID update-permission
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Param updateReq formData permissionModel.PermissionReq true "update permission request"
// @Success 200 {string} string "successfully edited permission"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /permission/{id} [put]
func (h *RoleHandler) UpdatePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[RoleHandler]").Start(c.Context(), "[UpdatePermission]")
		defer span.End()

		permissionID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var updateReq permissionModel.PermissionReq
		if err := reqvalidator.ReadRequest(c, &updateReq); err != nil {
			if updateReq.PermissionType == "" {
				return c.Status(fiber.StatusOK).JSON("update structure has no value")
			}
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())

			return errlst.Response(c, err)
		}

		res, err := h.service.UpdatePermission(ctx, permissionID, updateReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully edited permission")
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
