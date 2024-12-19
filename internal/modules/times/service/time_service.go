package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/samber/lo"
)

var _ times.Service = (*TimeService)(nil)

type TimeService struct {
	repo database.DataStore
}

func NewTimeService(repo database.DataStore) *TimeService {
	return &TimeService{repo: repo}
}

func (t *TimeService) AddTime(ctx context.Context, request timeModel.Request) (int, error) {
	timeID, err := t.repo.TimesRepo().AddTime(ctx, request.ToPsqlDBStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return timeID, nil
}

func (t *TimeService) GetTime(ctx context.Context, timeID int) (timeModel.DTO, error) {
	timeRes, err := t.repo.TimesRepo().GetTime(ctx, timeID)
	if err != nil {
		return timeModel.DTO{}, errlst.ParseErrors(err)
	}

	return timeRes.ToServer(), nil
}

func (t *TimeService) ListTimes(ctx context.Context, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[timeModel.DTO], error,
) {
	var (
		allTimeData       []timeModel.DAO
		timesListResponse abstract.PaginatedResponse[timeModel.DTO]
		totalCount        int
		err               error
	)

	err = t.repo.WithTransaction(ctx, func(db database.DataStore) error {
		totalCount, err = db.TimesRepo().CountOfTimes(ctx)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		timesListResponse.TotalItems = totalCount

		allTimeData, err = db.TimesRepo().ListTimes(ctx, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[timeModel.DTO]{}, errlst.ParseErrors(err)
	}

	timesList := lo.Map(
		allTimeData,
		func(item timeModel.DAO, _ int) timeModel.DTO {
			return item.ToServer()
		},
	)

	timesListResponse.Items = timesList
	timesListResponse.CurrentPage = paginationQuery.CurrentPage
	timesListResponse.Limit = paginationQuery.Limit
	timesListResponse.ItemsInCurrentPage = len(timesList)

	return timesListResponse, nil
}

func (t *TimeService) DeleteTime(ctx context.Context, timeID int) error {
	return t.repo.TimesRepo().DeleteTime(ctx, timeID)
}

func (t *TimeService) UpdateTime(ctx context.Context, timeID int, updateRequest *timeModel.UpdateRequest) (string, error) {
	var (
		updateRes string
		err       error
	)

	err = t.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err = db.TimesRepo().GetTime(ctx, timeID)
		if err != nil {
			return errlst.NewNotFoundError("[timeService][UpdateTime]: time not found for this timeID")
		}

		updateRes, err = db.TimesRepo().UpdateTime(ctx, updateRequest.ToPsqlDBStorage(timeID))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return updateRes, nil
}

func (t *TimeService) SelectActiveYear(ctx context.Context) (timeModel.DTO, error) {
	timeRes, err := t.repo.TimesRepo().SelectActiveYear(ctx)
	if err != nil {
		return timeModel.DTO{}, errlst.ParseErrors(err)
	}

	return timeRes.ToServer(), nil
}
