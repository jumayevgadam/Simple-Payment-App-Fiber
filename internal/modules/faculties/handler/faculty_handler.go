package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	facultyOps "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
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
func (h *FacultyHandler) GetFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != constants.SuperAdmin {
			return errlst.NewUnauthorizedError("only superadmin can get faculty")
		}

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

// ListGroupsByFacultyID handler.
func (h *FacultyHandler) ListGroupsByFacultyID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role_type").(string)
		if !ok || role != "superadmin" {
			return errlst.NewUnauthorizedError("only superadmin can list groups by faculty id")
		}

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
