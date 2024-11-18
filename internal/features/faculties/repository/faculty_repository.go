package repository

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/connection"
	"github.com/jumayevgadaym/tsu-toleg/internal/features/faculties"
	facultyModel "github.com/jumayevgadaym/tsu-toleg/internal/models/faculty"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
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
func (f *FacultyRepository) AddFaculty(ctx context.Context, facultyDAO facultyModel.DAO) (int, error) {
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
func (f *FacultyRepository) GetFaculty(ctx context.Context, facultyID int) (facultyModel.DAO, error) {
	var facultyDAO facultyModel.DAO

	if err := f.psqlDB.Get(
		ctx,
		f.psqlDB,
		&facultyDAO,
		getFacultyQuery,
		facultyID,
	); err != nil {
		return facultyModel.DAO{}, errlst.ParseSQLErrors(err)
	}

	return facultyDAO, nil
}

// ListFaculties repo fetches a list of faculties from DB.
func (f *FacultyRepository) ListFaculties(ctx context.Context) ([]facultyModel.DAO, error) {
	var facultyDAOs []facultyModel.DAO

	if err := f.psqlDB.Select(
		ctx,
		f.psqlDB,
		&facultyDAOs,
		listFacultiesQuery,
	); err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return facultyDAOs, nil
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
func (f *FacultyRepository) UpdateFaculty(ctx context.Context, facultyDAO facultyModel.DAO) (string, error) {
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
