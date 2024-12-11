package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	facultyOps "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

// Ensure FacultyHandler implements the facultyOps.Handlers interface.
var (
	_ facultyOps.Handlers = (*FacultyHandler)(nil)
)

// FacultyHandler for performing http request in handler layer calling methods from service.
type FacultyHandler struct {
	service services.DataService
}

// NewFacultyHandler creates and returns a new instance of FacultyHandler.
func NewFacultyHandler(service services.DataService) *FacultyHandler {
	return &FacultyHandler{service: service}
}

// AddFaculty handler.
// @Summary Add a new faculty.
// @Description Creates a new faculty and returns its id.
// @Tags Faculties
// @ID add-faculty
// @Accept multipart/form-data
// @Produce json
// @Param facultyReq formData facultyModel.Req true "faculty request for adding new one"
// @Success 200 {integer} integer 1
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/create [post].
func (h *FacultyHandler) AddFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var facultyReq facultyModel.Req
		if err := reqvalidator.ReadRequest(c, &facultyReq); err != nil {
			return errlst.Response(c, err)
		}

		resID, err := h.service.FacultyService().AddFaculty(c.Context(), &facultyReq)
		if err != nil {
			return errlst.Response(c, err)
		}

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
// @Success 200 {object} facultyModel.DTO
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/{id} [get].
func (h *FacultyHandler) GetFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		faculty, err := h.service.FacultyService().GetFaculty(c.Context(), facultyID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(faculty)
	}
}

// ListFaculties handler fetches a list of faculties.
// @Summary List faculties.
// @Description list faculties, pagination not setted.
// @Tags Faculties
// @ID list-faculties
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param limit query int false "number of elements per page" Format(limit)
// @Param orderBy query string false "filter name" Format(orderBy)
// @Success 200 {object} []facultyModel.DTO
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /faculty/get-all [get].
func (h *FacultyHandler) ListFaculties() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationReq, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.Response(c, err)
		}

		faculties, err := h.service.FacultyService().ListFaculties(c.Context(), paginationReq)
		if err != nil {
			return errlst.Response(c, err)
		}

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
// @Router /faculty/{id} [delete].
func (h *FacultyHandler) DeleteFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		err = h.service.FacultyService().DeleteFaculty(c.Context(), facultyID)
		if err != nil {
			return errlst.Response(c, err)
		}

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
// @Router /faculty/{id} [put].
func (h *FacultyHandler) UpdateFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		// Parse the request body into `UpdateInputReq`.
		var inputReq facultyModel.UpdateInputReq
		if err := reqvalidator.ReadRequest(c, &inputReq); err != nil {
			res, err := inputReq.Validate()
			if err == nil {
				return c.JSON(res)
			}

			return errlst.Response(c, err)
		}

		res, err := h.service.FacultyService().UpdateFaculty(c.Context(), facultyID, &inputReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}
