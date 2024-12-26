package repository

import (
	"context"

	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

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

func (r *UserRepository) CountStudentsByGroupID(ctx context.Context, groupID int) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		countStudentsByGroupIDQuery,
		groupID,
	)

	if err != nil {
		return 0, errlst.ParseSQLErrors(err)
	}

	return totalCount, nil
}

func (r *UserRepository) ListStudentsByGroupID(ctx context.Context, groupID int, paginationData abstract.PaginationData) (
	[]*userModel.StudentDataByGroupID, error,
) {
	var studentDataByGroupID []*userModel.StudentDataByGroupID
	offset := (paginationData.CurrentPage - 1) * paginationData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&studentDataByGroupID,
		listStudentsByGroupIDQuery,
		groupID,
		offset,
		paginationData.Limit,
	)

	if err != nil {
		return nil, errlst.ParseSQLErrors(err)
	}

	return studentDataByGroupID, nil
}

func (r *UserRepository) UpdateStudent(ctx context.Context, updateData userModel.StudentUpdateData) (string, error) {
	var updateRes string

	err := r.psqlDB.QueryRow(
		ctx,
		updateStudentQuery,
		updateData.GroupID,
		updateData.Name,
		updateData.Surname,
		updateData.Username,
		updateData.Password,
		updateData.StudentID,
	).Scan(&updateRes)

	if err != nil {
		return "", errlst.ParseSQLErrors(err)
	}

	return updateRes, nil
}
