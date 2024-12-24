package handler

import (
	"github.com/gofiber/fiber/v2"
	statisticsModel "github.com/jumayevgadam/tsu-toleg/internal/models/statistics"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

func (h *PaymentHandler) AdminGetStatisticsAboutYear() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var academicYear statisticsModel.AcademicYear

		err := reqvalidator.ReadRequest(c, &academicYear)
		if err != nil {
			academicYear.EndYear = academicYear.StartYear + 1

			return errlst.Response(c, err)
		}

		return nil
	}
}

func (h *PaymentHandler) AdminGetStatisticsAboutFaculty() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
