package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	facultyModel "github.com/jumayevgadam/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/faculties"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Ensure FacultyRepository implements the faculties.Repository interface.
var (
	_ faculties.Repository = (*FacultyRepository)(nil)
)

// FacultyRepository performing database operations for faculties.
type FacultyRepository struct {
	psqlDB connection.DB
}

// NewFacultyRepository creates and returns a new instance of FacultyRepository.
func NewFacultyRepository(psqlDB connection.DB) *FacultyRepository {
	return &FacultyRepository{psqlDB: psqlDB}
}

// AddFaculty repo insert faculty data into db and returns id.
func (f *FacultyRepository) AddFaculty(ctx context.Context, facultyDAO *facultyModel.DAO) (int, error) {
	var facultyID int

	if err := f.psqlDB.QueryRow(
		ctx,
		addFacultyQuery,
		facultyDAO.Name,
		facultyDAO.Code,
	).Scan(&facultyID); err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return facultyID, nil
}

// GetFaculty repo fetches faculty by using identified id.
func (f *FacultyRepository) GetFaculty(ctx context.Context, facultyID int) (*facultyModel.FacultyData, error) {
	var facultyDAO facultyModel.FacultyData

	if err := f.psqlDB.Get(
		ctx,
		f.psqlDB,
		&facultyDAO,
		getFacultyQuery,
		facultyID,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &facultyDAO, nil
}

// ListFaculties repo fetches a list of faculties from DB.
func (f *FacultyRepository) ListFaculties(ctx context.Context, paginationData abstract.PaginationData) ([]*facultyModel.FacultyData, error) {
	var facultyDAOs []*facultyModel.FacultyData
	offset := (paginationData.Page - 1) * paginationData.Limit

	if err := f.psqlDB.Select(
		ctx,
		f.psqlDB,
		&facultyDAOs,
		listFacultiesQuery,
		offset,
		paginationData.Limit,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return facultyDAOs, nil
}

// CountFaculties give number of faculties.
func (f *FacultyRepository) CountFaculties(ctx context.Context) (int, error) {
	var totalCount int

	err := f.psqlDB.Get(
		ctx,
		f.psqlDB,
		&totalCount,
		countFacultiesQuery,
	)
	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

// DeleteFaculty repo deletes faculty from DB using identified faculty id.
func (f *FacultyRepository) DeleteFaculty(ctx context.Context, facultyID int) error {
	_, err := f.psqlDB.Exec(
		ctx,
		deleteFacultyQuery,
		facultyID,
	)
	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	return nil
}

// UpdateFaculty repo updates faculty data using a new data and identified faculty id.
func (f *FacultyRepository) UpdateFaculty(ctx context.Context, facultyDAO *facultyModel.FacultyData) (string, error) {
	var res string

	if err := f.psqlDB.QueryRow(
		ctx,
		updateFacultyQuery,
		facultyDAO.Name,
		facultyDAO.Code,
		facultyDAO.ID,
	).Scan(&res); err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return res, nil
}
