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

func (s *UserService) GetAdmin(ctx context.Context, adminID int) (*userModel.Admin, error) {
	adminData, err := s.repo.UsersRepo().GetAdmin(ctx, adminID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return adminData.ToServer(), nil
}

func (s *UserService) DeleteAdmin(ctx context.Context, adminID int) error {
	return s.repo.UsersRepo().DeleteAdmin(ctx, adminID)
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
