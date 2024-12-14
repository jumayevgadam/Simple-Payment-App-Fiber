package repository

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/connection"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

var _ users.Repository = (*UserRepository)(nil)

type UserRepository struct {
	psqlDB connection.DB
}

func NewUserRepository(psqlDB connection.DB) *UserRepository {
	return &UserRepository{psqlDB: psqlDB}
}

func (r *UserRepository) AddStudent(ctx context.Context, res userModel.Response) (int, error) {
	var userID int

	err := r.psqlDB.QueryRow(
		ctx,
		addStudentQuery,
		res.RoleID,
		res.GroupID,
		res.Name,
		res.Surname,
		res.Username,
		res.Password,
	).Scan(&userID)

	if err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return userID, nil
}

func (r *UserRepository) AddAdmin(ctx context.Context, res userModel.AdminResponse) (int, error) {
	var adminID int

	err := r.psqlDB.QueryRow(
		ctx,
		addAdminQuery,
		res.Name,
		res.Surname,
		res.Username,
		res.Password,
	).Scan(&adminID)

	if err != nil {
		return -1, errlst.ParseSQLErrors(err)
	}

	return adminID, nil
}

func (r *UserRepository) ListAdmins(ctx context.Context, paginationData abstract.PaginationData) (
	[]*userModel.AdminData, error,
) {
	var adminDatas []*userModel.AdminData
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&adminDatas,
		listAdminsQuery,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return adminDatas, nil
}

func (r *UserRepository) CountAdmins(ctx context.Context) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		totalAdminCountQuery,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

func (r *UserRepository) CountStudents(ctx context.Context) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		totalStudentCountQuery,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

func (r *UserRepository) ListStudents(ctx context.Context, paginationData abstract.PaginationData) (
	[]*userModel.StudentData, error,
) {
	var studentDatas []*userModel.StudentData
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&studentDatas,
		listStudentsQuery,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return studentDatas, nil
}

func (r *UserRepository) GetAdmin(ctx context.Context, adminID int) (*userModel.AdminData, error) {
	var adminData userModel.AdminData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&adminData,
		getAdminQuery,
		adminID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &adminData, nil
}

func (r *UserRepository) GetStudent(ctx context.Context, studentID int) (*userModel.StudentData, error) {
	var studentData userModel.StudentData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&studentData,
		getStudentQuery,
		studentID,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return &studentData, nil
}

func (r *UserRepository) DeleteAdmin(ctx context.Context, adminID int) error {
	result, err := r.psqlDB.Exec(
		ctx,
		deleteAdminQuery,
		adminID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	if result.RowsAffected() == 0 {
		return errlst.NewNotFoundError("[userRepository][DeleteAdmin]: admin not found")
	}

	return nil
}

func (r *UserRepository) DeleteStudent(ctx context.Context, studentID int) error {
	result, err := r.psqlDB.Exec(
		ctx,
		deleteStudentQuery,
		studentID,
	)

	if err != nil {
		return errlst.ParseSQLErrors(err)
	}

	if result.RowsAffected() == 0 {
		return errlst.NewNotFoundError("[userRepository][DeleteStudent]: student not found")
	}

	return nil
}
