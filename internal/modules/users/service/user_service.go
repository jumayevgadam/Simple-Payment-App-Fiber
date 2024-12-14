package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
	"github.com/samber/lo"
)

var _ users.Service = (*UserService)(nil)

type UserService struct {
	mw   *middleware.Manager
	repo database.DataStore
}

func NewUserService(mw *middleware.Manager, repo database.DataStore) *UserService {
	return &UserService{mw: mw, repo: repo}
}

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

func (s *UserService) AddAdmin(ctx context.Context, request userModel.AdminRequest) (int, error) {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return -1, errlst.NewBadRequestError(err.Error())
	}
	request.Password = hashedPass

	adminID, err := s.repo.UsersRepo().AddAdmin(ctx, request.ToPsqlDBStorage())
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return adminID, nil
}

func (s *UserService) ListAdmins(ctx context.Context, paginationQuery abstract.PaginationQuery) (
	abstract.PaginatedResponse[*userModel.Admin], error,
) {
	var (
		allAdminData      []*userModel.AdminData
		adminListResponse abstract.PaginatedResponse[*userModel.Admin]
		totalAdminCount   int
		err               error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		totalAdminCount, err = db.UsersRepo().CountAdmins(ctx)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		adminListResponse.TotalItems = totalAdminCount

		allAdminData, err = db.UsersRepo().ListAdmins(ctx, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*userModel.Admin]{}, errlst.ParseErrors(err)
	}

	adminList := lo.Map(
		allAdminData,
		func(item *userModel.AdminData, _ int) *userModel.Admin {
			return item.ToServer()
		},
	)

	adminListResponse.Items = adminList
	adminListResponse.CurrentPage = paginationQuery.CurrentPage
	adminListResponse.Limit = paginationQuery.Limit
	adminListResponse.ItemsInCurrentPage = len(adminList)

	return adminListResponse, nil
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

func (s *UserService) GetAdmin(ctx context.Context, adminID int) (*userModel.Admin, error) {
	adminData, err := s.repo.UsersRepo().GetAdmin(ctx, adminID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return adminData.ToServer(), nil
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

func (s *UserService) DeleteAdmin(ctx context.Context, adminID int) error {
	return s.repo.UsersRepo().DeleteAdmin(ctx, adminID)
}

func (s *UserService) DeleteStudent(ctx context.Context, studentID int) error {
	return s.repo.UsersRepo().DeleteStudent(ctx, studentID)
}
