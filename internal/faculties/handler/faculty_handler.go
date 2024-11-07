package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	facultyOps "github.com/jumayevgadaym/tsu-toleg/internal/faculties"
	facultyModel "github.com/jumayevgadaym/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadaym/tsu-toleg/pkg/reqvalidator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var (
	_ facultyOps.Handlers = (*FacultyHandler)(nil)
)

// FacultyHandler is
type FacultyHandler struct {
	service facultyOps.Service
}

// NewFacultyHandler is
func NewFacultyHandler(service facultyOps.Service) *FacultyHandler {
	return &FacultyHandler{service: service}
}

// AddFaculty handler is
func (h *FacultyHandler) AddFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[FacultyHandler]").Start(c.Context(), "[AddFaculty]")
		defer span.End()

		var facultyReq facultyModel.DTO
		if err := reqvalidator.ReadRequest(c, &facultyReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}

		resID, err := h.service.AddFaculty(ctx, &facultyReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully added faculty")
		return c.Status(fiber.StatusOK).JSON(resID)
	}
}

// GetFaculty handler is
func (h *FacultyHandler) GetFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[FacultyHandler]").Start(c.Context(), "[GetFaculty]")
		defer span.End()

		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrBadRequest.Error())
			return errlst.Response(c, err)
		}

		faculty, err := h.service.GetFaculty(ctx, facultyID)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "Successfully got faculty")
		return c.Status(fiber.StatusOK).JSON(faculty)
	}
}

// ListFaculties handler is
func (h *FacultyHandler) ListFaculties() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[FacultyHandler]").Start(c.Context(), "[ListFaculties]")
		defer span.End()

		faculties, err := h.service.ListFaculties(ctx)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully got all faculties")
		return c.Status(fiber.StatusOK).JSON(faculties)
	}
}

// DeleteFaculty handler is
func (h *FacultyHandler) DeleteFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[FacultyHandler]").Start(c.Context(), "[DeleteFaculty]")
		defer span.End()

		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrBadQueryParams.Error())
			return errlst.Response(c, err)
		}

		if err := h.service.DeleteFaculty(ctx, facultyID); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully deleted faculty")
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"response": "successfully deleted faculty",
		})
	}
}

// UpdateFaculty handler is
func (h *FacultyHandler) UpdateFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[FacultyHandler]").Start(c.Context(), "[UpdateFaculty]")
		defer span.End()

		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var facultyReq facultyModel.DTO
		if err := reqvalidator.ReadRequest(c, &facultyReq); err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrFieldValidation.Error())
			return errlst.Response(c, err)
		}
		facultyReq.ID = facultyID

		res, err := h.service.UpdateFaculty(ctx, &facultyReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully updated faculty")
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
