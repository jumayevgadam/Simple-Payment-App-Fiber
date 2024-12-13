package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
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
