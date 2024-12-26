package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (h *PaymentHandler) AdminGetStatisticsAboutYear() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var academicYearReq timeModel.AcademicYearRequest

		startYearStr := c.Query("start-year")
		if startYearStr == "" {
			return errlst.Response(c, errors.New("start-year can not be empty"))
		}

		startYear, err := strconv.Atoi(startYearStr)
		if err != nil {
			return errlst.Response(c, err)
		}
		academicYearReq.StartYear = startYear

		if academicYearReq.EndYear == 0 {
			academicYearReq.EndYear = academicYearReq.StartYear + 1
		}

		stcByTSU, err := h.service.PaymentService().AdminGetStatisticsAboutYear(c.Context(), academicYearReq)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(stcByTSU)
	}
}

func (h *PaymentHandler) AdminGetStatisticsAboutFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		facultyID := c.Query("faculty-id")
		if facultyID == "" {
			return errlst.Response(c, errors.New("error: faculty-id can not be empty"))
		}

		return nil
	}
}
