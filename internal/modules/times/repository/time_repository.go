package repository

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	timeModel "github.com/jumayevgadam/tsu-toleg/internal/models/time"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/times"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
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

func (t *TimeRepository) GetTime(ctx context.Context, timeID int) (timeModel.DAO, error) {
	var timeDAO timeModel.DAO

	err := t.psqlDB.Get(
		ctx,
		t.psqlDB,
		&timeDAO,
		"SELECT id, start_year, end_year, is_active FROM times WHERE id = $1;",
		timeID,
	)

	if err != nil {
		return timeModel.DAO{}, errlst.ParseSQLErrors(err)
	}

	return timeDAO, nil
}

func (t *TimeRepository) CountOfTimes(ctx context.Context) (int, error) {
	var totalCount int

	err := t.psqlDB.Get(
		ctx,
		t.psqlDB,
		&totalCount,
		"SELECT COUNT(id) FROM times;",
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

func (t *TimeRepository) ListTimes(ctx context.Context, paginationData abstract.PaginationData) ([]timeModel.DAO, error) {
	var times []timeModel.DAO
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := t.psqlDB.Select(
		ctx,
		t.psqlDB,
		&times,
		"SELECT id, start_year, end_year, is_active FROM times ORDER BY id DESC OFFSET $1 LIMIT $2;",
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return times, nil
}

func (t *TimeRepository) DeleteTime(ctx context.Context, timeID int) error {
	result, err := t.psqlDB.Exec(
		ctx,
		"DELETE FROM times WHERE id = $1;",
		timeID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	if result.RowsAffected() == 0 {
		return errlst.NewNotFoundError("[timeRepository][DeleteTime]: time not found for deleting")
	}

	return nil
}

func (t *TimeRepository) UpdateTime(ctx context.Context, timeDAO timeModel.DAO) (string, error) {
	var updateRes string

	err := t.psqlDB.QueryRow(
		ctx,
		`UPDATE times 
			SET start_year = COALESCE(NULLIF($1, 0), start_year),
				end_year = COALESCE(NULLIF($2, 0), end_year),
				is_active = $3
		WHERE id = $4
		RETURNING 'successfully updated times';`,
		timeDAO.StartYear,
		timeDAO.EndYear,
		timeDAO.IsActive,
		timeDAO.ID,
	).Scan(&updateRes)

	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	fmt.Printf("Updating TimeDAO: %+v\n", timeDAO)

	return updateRes, nil
}

func (t *TimeRepository) SelectActiveYear(ctx context.Context) (timeModel.DAO, error) {
	var timeDAO timeModel.DAO

	err := t.psqlDB.Get(
		ctx,
		t.psqlDB,
		&timeDAO,
		"SELECT id, start_year, end_year, is_active FROM times WHERE is_active IS TRUE",
	)

	if err != nil {
		return timeModel.DAO{}, errlst.ParseSQLErrors(err)
	}

	return timeDAO, nil
}
