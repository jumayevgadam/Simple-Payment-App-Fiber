package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	facultyOps "github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
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

// AddFaculty handler.
func (h *FacultyHandler) AddFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var facultyReq facultyModel.DTO
		if err := reqvalidator.ReadRequest(c, &facultyReq); err != nil {
			return errlst.Response(c, err)
		}

		resID, err := h.service.AddFaculty(c.Context(), &facultyReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(resID)
	}
}

// GetFaculty handler fetches faculty using identified id.
func (h *FacultyHandler) GetFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facultyID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		faculty, err := h.service.GetFaculty(c.Context(), facultyID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(faculty)
	}
}

// ListFaculties handler fetches a list of faculties.
func (h *FacultyHandler) ListFaculties() fiber.Handler {
	return func(c *fiber.Ctx) error {
		faculties, err := h.service.ListFaculties(c.Context())
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

		if err := h.service.DeleteFaculty(c.Context(), facultyID); err != nil {
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

		res, err := h.service.UpdateFaculty(c.Context(), facultyID, &inputReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}
