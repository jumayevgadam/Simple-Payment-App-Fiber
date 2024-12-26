package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
	"github.com/samber/lo"
)

func (s *UserService) AddStudent(ctx context.Context, request userModel.Request) (int, error) {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}
	request.Password = hashedPass

	userID, err := s.repo.UsersRepo().AddStudent(ctx, request.ToPsqlDBStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return userID, nil
}

func (s *UserService) ListStudents(ctx context.Context, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[*userModel.Student], error,
) {
	var (
		allStudentData      []*userModel.StudentData
		studentListResponse abstract.PaginatedResponse[*userModel.Student]
		totalStudentCount   int
		err                 error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		totalStudentCount, err = db.UsersRepo().CountStudents(ctx)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		studentListResponse.TotalItems = totalStudentCount

		allStudentData, err = db.UsersRepo().ListStudents(ctx, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*userModel.Student]{}, errlst.ParseErrors(err)
	}

	studentList := lo.Map(
		allStudentData,
		func(item *userModel.StudentData, _ int) *userModel.Student {
			return item.ToServer()
		},
	)

	studentListResponse.Items = studentList
	studentListResponse.CurrentPage = paginationQuery.CurrentPage
	studentListResponse.Limit = paginationQuery.Limit
	studentListResponse.ItemsInCurrentPage = len(studentList)

	return studentListResponse, nil
}

func (s *UserService) GetStudent(ctx context.Context, studentID int) (*userModel.Student, error) {
	studentData, err := s.repo.UsersRepo().GetStudent(ctx, studentID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return studentData.ToServer(), nil
}

func (s *UserService) ListStudentsByGroupID(ctx context.Context, groupID int, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[*userModel.StudentResGroupID], error,
) {
	var (
		allStudentDataByGroupID      []*userModel.StudentDataByGroupID
		studentListResponseByGroupID abstract.PaginatedResponse[*userModel.StudentResGroupID]
		totalStudentCountByGroupID   int
		err                          error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err = db.GroupsRepo().GetGroup(ctx, groupID)
		if err != nil {
			return errlst.NewNotFoundError("[userService][ListStudentsByGroupID]: group not found")
		}

		totalStudentCountByGroupID, err = db.UsersRepo().CountStudentsByGroupID(ctx, groupID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		studentListResponseByGroupID.TotalItems = totalStudentCountByGroupID

		allStudentDataByGroupID, err = db.UsersRepo().ListStudentsByGroupID(ctx, groupID, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*userModel.StudentResGroupID]{}, errlst.ParseErrors(err)
	}

	studentListResponse := lo.Map(
		allStudentDataByGroupID,
		func(item *userModel.StudentDataByGroupID, _ int) *userModel.StudentResGroupID {
			return item.ToServer()
		},
	)

	studentListResponseByGroupID.Items = studentListResponse
	studentListResponseByGroupID.CurrentPage = paginationQuery.CurrentPage
	studentListResponseByGroupID.Limit = paginationQuery.Limit
	studentListResponseByGroupID.ItemsInCurrentPage = len(studentListResponse)

	return studentListResponseByGroupID, nil
}

func (s *UserService) DeleteStudent(ctx context.Context, studentID int) error {
	return s.repo.UsersRepo().DeleteStudent(ctx, studentID)
}

func (s *UserService) UpdateStudent(ctx context.Context, studentID int, updateRequest userModel.StudentUpdateRequest) (string, error) {
	var (
		updateRes string
		err       error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err = db.UsersRepo().GetStudent(ctx, studentID)
		if err != nil {
			return errlst.NewNotFoundError("[userService][UpdateStudent]: student not found for updating")
		}

		if updateRequest.Password != "" {
			if len(updateRequest.Password) < constants.MinPasswordLength {
				return errlst.NewBadRequestError(constants.ErrMinPasswordLength)
			}

			hashedPass, err := utils.HashPassword(updateRequest.Password)
			if err != nil {
				return errlst.ParseErrors(err)
			}

			updateRequest.Password = hashedPass
		}

		updateRes, err = db.UsersRepo().UpdateStudent(ctx, updateRequest.ToPsqlDBStorage(studentID))
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
