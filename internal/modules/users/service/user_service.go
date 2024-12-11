package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/abstract"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
	"github.com/samber/lo"
)

// Ensure UserService implements the users.Service interface.
var (
	_ users.Service = (*UserService)(nil)
)

// UserService manages buisiness logic for modules/user part of application.
type UserService struct {
	mw   *middleware.Manager
	repo database.DataStore
}

// NewUserService creates and returns a new instance of UserRepository.
func NewUserService(mw *middleware.Manager, repo database.DataStore) *UserService {
	return &UserService{mw: mw, repo: repo}
}

// CreateUser service insert a user into db and returns its id.
func (s *UserService) Register(ctx context.Context, request userModel.SignUpReq) (int, error) {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}
	request.Password = hashedPass

	userID, err := s.repo.UsersRepo().CreateUser(ctx, request.ToStorage(constants.DefaultRoleID))
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return userID, nil
}

// Login service for login.
func (s *UserService) Login(ctx context.Context, loginReq userModel.LoginReq) (string, error) {
	var (
		token string
		err   error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// get user details by username.
		user, err := db.UsersRepo().GetUserByUsername(ctx, loginReq.Username)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		// Compare passwords.
		err = utils.CheckAndComparePassword(loginReq.Password, user.Password)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		// getRole by roleID.
		roleDAO, err := db.RolesRepo().GetRole(ctx, user.RoleID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		// getPermissions for that roleID.
		permissions, err := db.RolesRepo().GetPermissionsByRoleID(ctx, user.RoleID)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		// token generate.
		token, err = s.mw.GenerateToken(user.ID, user.RoleID, user.Username, roleDAO.RoleName, permissions)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return "", errlst.ParseErrors(err)
	}

	return token, nil
}

// GetUserByID service.
func (s *UserService) GetUserByID(ctx context.Context, userID int) (*userModel.AllUserDTO, error) {
	userRes, err := s.repo.UsersRepo().GetUserByID(ctx, userID)
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	return userRes.ToServer(), nil
}

// ChangeRoleOfUser service.
func (s *UserService) UpdateUser(ctx context.Context, userID int, updateReq *userModel.UpdateUserDetails) error {
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		_, err := db.UsersRepo().GetUserByID(ctx, userID)

		if err != nil {
			return errlst.NewNotFoundError(errlst.ErrNoSuchUser)
		}

		err = db.UsersRepo().UpdateUser(ctx, userID, updateReq.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}

// ListUsers service.
func (s *UserService) ListAllUsers(ctx context.Context, paginationRequest abstract.PaginationQuery) (
	abstract.PaginatedResponse[*userModel.AllUserDTO], error,
) {
	var (
		dataAllOfUsers    []*userModel.AllUserDAO
		usersListResponse abstract.PaginatedResponse[*userModel.AllUserDTO]
		totalUsersCount   int
		err               error
	)

	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		totalUsersCount, err = db.UsersRepo().CountAllUsers(ctx)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		usersListResponse.TotalItems = totalUsersCount

		dataAllOfUsers, err = db.UsersRepo().ListAllUsers(ctx, paginationRequest.ToPsqlDBStorage())
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})

	if err != nil {
		return abstract.PaginatedResponse[*userModel.AllUserDTO]{}, errlst.ParseErrors(err)
	}

	usersList := lo.Map(
		dataAllOfUsers,
		func(item *userModel.AllUserDAO, _ int) *userModel.AllUserDTO {
			return item.ToServer()
		},
	)

	usersListResponse.Items = usersList
	usersListResponse.Page = paginationRequest.Page
	usersListResponse.Limit = len(usersList)

	return usersListResponse, nil
}

// DeleteUser service.
func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	err := s.repo.UsersRepo().DeleteUser(ctx, userID)

	if err != nil {
		return errlst.ParseErrors(err)
	}

	return nil
}
