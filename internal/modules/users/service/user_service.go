package service

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
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

func (s *UserService) Login(ctx context.Context, loginRequest userModel.LoginRequest) (
	userModel.LoginResponseWithToken, error,
) {
	userDetails, err := s.repo.UsersRepo().Login(ctx, loginRequest.Username)
	if err != nil {
		return userModel.LoginResponseWithToken{}, errlst.ParseErrors(err)
	}

	err = utils.CheckAndComparePassword(loginRequest.Password, userDetails.Password)
	if err != nil {
		return userModel.LoginResponseWithToken{}, errlst.ParseErrors(err)
	}

	token, err := s.mw.GenerateToken(userDetails.UserID, userDetails.RoleID, userDetails.Username, userDetails.RoleType)
	if err != nil {
		return userModel.LoginResponseWithToken{}, errlst.ParseErrors(err)
	}

	resp := userModel.LoginResponseWithToken{
		LoginResponse: userDetails.ToServer(),
		Token:         token,
	}

	return resp, nil
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

func (s *UserService) UpdateAdmin(ctx context.Context, adminID int, updateRequest userModel.AdminUpdateRequest) (string, error) {
	var (
		updateRes string
		err       error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err = db.UsersRepo().GetAdmin(ctx, adminID)
		if err != nil {
			return errlst.NewNotFoundError("[userService][UpdateAdmin]: admin not found for updating")
		}

		if updateRequest.Password != "" {
			if len(updateRequest.Password) < constants.MinPasswordLength {
				return errlst.NewBadRequestError(constants.ErrMinPasswordLength)
			}

			hashedPass, err := utils.HashPassword(updateRequest.Password)
			if err != nil {
				return errlst.NewBadRequestError(err.Error())
			}

			updateRequest.Password = hashedPass
		}

		updateRes, err = db.UsersRepo().UpdateAdmin(ctx, updateRequest.ToPsqlDBStorage(adminID))
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

func (s *UserService) AdminFindStudent(ctx context.Context, filterStudent userModel.FilterStudent,
	paginationQuery abstract.PaginationQuery,
) (
	abstract.PaginatedResponse[*userModel.AllStudentDTO], error,
) {
	var (
		allStudentDataWithFilter  []*userModel.AllStudentData
		studentListResponseFilter abstract.PaginatedResponse[*userModel.AllStudentDTO]
		err                       error
	)

	fmt.Println(filterStudent)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		allStudentDataWithFilter, err = db.UsersRepo().AdminFindStudent(ctx, filterStudent, paginationQuery.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		studentListResponseFilter.TotalItems = len(allStudentDataWithFilter)

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*userModel.AllStudentDTO]{}, errlst.ParseErrors(err)
	}

	studentListWithFilter := lo.Map(
		allStudentDataWithFilter,
		func(item *userModel.AllStudentData, _ int) *userModel.AllStudentDTO {
			return item.ToServer()
		},
	)

	studentListResponseFilter.Items = studentListWithFilter
	studentListResponseFilter.CurrentPage = paginationQuery.CurrentPage
	studentListResponseFilter.Limit = paginationQuery.Limit
	studentListResponseFilter.ItemsInCurrentPage = len(studentListWithFilter)

	return studentListResponseFilter, nil
}
