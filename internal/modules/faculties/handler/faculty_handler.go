package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	facultyOps "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// Ensure FacultyHandler implements the facultyOps.Handlers interface.
var (
	_ facultyOps.Handlers = (*FacultyHandler)(nil)
)

// FacultyHandler for performing http request in handler layer calling methods from service.
type FacultyHandler struct {
	service facultyOps.Service
}

// NewFacultyHandler creates and returns a new instance of FacultyHandler.
func NewFacultyHandler(service facultyOps.Service) *FacultyHandler {
	return &FacultyHandler{service: service}
}

// AddFaculty handler processes requests and returns faculty's id.
// @Summary Add a new faculty.
// @Description Creates a new faculty and returns its id.
// @Tags Faculties
// @ID add-faculty
// @Accept multipart/form-data
// @Produce json
// @Param facultyReq formData facultyModel.DTO true "faculty request for adding new one"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/create [post]
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

// GetFaculty handler fetches faculty using identified id.
// @Summary get-faculty.
// @Description retrieve faculty using its id.
// @Tags Faculties
// @ID get-faculty
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} facultyModel.Faculty
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/{id} [get]
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

// ListFaculties handler fetches a list of faculties.
// @Summary List faculties.
// @Description list faculties, pagination not setted.
// @Tags Faculties
// @ID list-faculties
// @Produce json
// @Success 200 {object} []facultyModel.Faculty
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/get-all [get]
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

// DeleteFaculty handler deletes a faculty using identified id.
// @Summary delete-faculty
// @Description delete faculty using identified id.
// @Tags Faculties
// @ID delete-faculty
// @Produce json
// @Param id path int true "id"
// @Success 200 {string} string "successfully deleted faculty"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/{id} [delete]
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

// UpdateFaculty handler updates a faculty using a new faculty data and identified id.
// @Summary Update Faculty
// @Description update faculty fields using identified faculty id
// @Tags Faculties
// @ID update-faculty
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "needed faculty id for update"
// @Param inputReq formData facultyModel.UpdateInputReq true "update faculty request"
// @Success 200 {string} string "successfully updated faculty ops"
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/{id} [put]
func (h *FacultyHandler) UpdateFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, span := otel.Tracer("[FacultyHandler]").Start(c.Context(), "[UpdateFaculty]")
		defer span.End()
		var res string

		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var inputReq facultyModel.UpdateInputReq
		if err := inputReq.Validate(); err != nil {
			res = "update structure has no value"
			return c.JSON(res)
		} else {
			err = reqvalidator.ReadRequest(c, &inputReq)
			if err != nil {
				log.Println(err)
				return errlst.Response(c, err)
			}
		}

		res, err = h.service.UpdateFaculty(ctx, facultyID, &inputReq)
		if err != nil {
			tracing.EventErrorTracer(span, err, errlst.ErrInternalServer.Error())
			return errlst.Response(c, err)
		}

		span.SetStatus(codes.Ok, "successfully updated faculty")
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
