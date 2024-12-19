package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/services"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/reqvalidator"
)

var _ times.Handlers = (*TimeHandler)(nil)

type TimeHandler struct {
	service services.DataService
}

func NewTimeHandler(service services.DataService) *TimeHandler {
	return &TimeHandler{service: service}
}

func (t *TimeHandler) AddTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req timeModel.Request

		err := reqvalidator.ReadRequest(c, &req)
		if err != nil {
			return errlst.NewBadRequestError(err.Error())
		}

		timeID, err := t.service.TimeService().AddTime(c.Context(), req)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(timeID)
	}
}

func (t *TimeHandler) GetTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		timeID, err := strconv.Atoi(c.Params("time_id"))
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		time, err := t.service.TimeService().GetTime(c.Context(), timeID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(time)
	}
}

func (t *TimeHandler) ListTimes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginationQuery, err := abstract.GetPaginationFromFiberCtx(c)
		if err != nil {
			return errlst.NewBadQueryParamsError(err)
		}

		listOfTimes, err := t.service.TimeService().ListTimes(c.Context(), paginationQuery)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(listOfTimes)
	}
}

func (t *TimeHandler) DeleteTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		timeID, err := strconv.Atoi(c.Params("time_id"))
		if err != nil {
			return errlst.NewBadRequestError(err)
		}

		err = t.service.TimeService().DeleteTime(c.Context(), timeID)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("successfully deleted time with id: %d", timeID),
		})
	}
}

func (t *TimeHandler) UpdateTime() fiber.Handler {
	return func(c *fiber.Ctx) error {
		timeID, err := strconv.Atoi(c.Params("time_id"))
		if err != nil {
			return errlst.Response(c, err)
		}

		var updateRequest timeModel.UpdateRequest

		err = reqvalidator.ReadRequest(c, &updateRequest)
		if err != nil {
			return errlst.Response(c, err)
		}

		// Debugging: Log the parsed updateRequest.
		fmt.Printf("Parsed UpdateRequest: %+v\n", updateRequest)

		updateRes, err := t.service.TimeService().UpdateTime(c.Context(), timeID, &updateRequest)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  fmt.Sprintf("updated timeID: %d", timeID),
			"response": updateRes,
		})
	}
}

func (t *TimeHandler) SelectActiveYear() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activeYear, err := t.service.TimeService().SelectActiveYear(c.Context())
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.Status(fiber.StatusOK).JSON(activeYear)
	}
}
