package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

var _ times.Repository = (*TimeRepository)(nil)

type TimeRepository struct {
	psqlDB connection.DB
}

func NewTimeRepository(psqlDB connection.DB) *TimeRepository {
	return &TimeRepository{psqlDB: psqlDB}
}

func (t *TimeRepository) AddTime(ctx context.Context, res timeModel.Response) (int, error) {
	var timeID int

	err := t.psqlDB.QueryRow(
		ctx,
		"INSERT INTO times (start_year, end_year) VALUES ($1, $2) RETURNING id;",
		res.StartYear,
		res.EndYear,
	).Scan(&timeID)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return timeID, nil
}
