package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

// Ensure UserService implements the users.Service interface.
var (
	_ users.Service = (*UserService)(nil)
)

// Creating dynamic roleMap.
var RoleMap = map[string]int{
	"superadmin": 1,
	"admin":      2,
	"student":    3,
}

// UserService manages buisiness logic for modules/user part of application.
type UserService struct {
	mw   *middleware.MiddlewareManager
	repo database.DataStore
}

// NewUserService creates and returns a new instance of UserRepository.
func NewUserService(mw *middleware.MiddlewareManager, repo database.DataStore) *UserService {
	return &UserService{mw: mw, repo: repo}
}

// CreateUser service insert a user into db and returns its id.
func (s *UserService) CreateUser(ctx context.Context, request userModel.SignUpReq, role string) (int, error) {
	var (
		userID int
		err    error
	)
	err = s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// Check that role exist or not in database, do not try with dynamic methods.
		// Return id of that role
		roles, err := db.RolesRepo().GetRoleByRoleName(ctx, role)
		if err != nil {
			return errlst.ErrNoSuchRole
		}

		if roles.RoleName == "student" {
			if request.GroupID == nil {
				return errlst.NewBadRequestError("group id must need for student")
			}
		} else {
			if request.GroupID != nil {
				return errlst.NewBadRequestError("group id does not need for non-student roles")
			}
		}

		hashedPass, err := utils.HashPassword(request.Password)
		if err != nil {
			return errlst.ParseErrors(err)
		}
		request.Password = hashedPass

		userID, err = s.repo.UsersRepo().CreateUser(ctx, request.ToStorage(roles.ID))
		if err != nil {
			return errlst.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}

	return userID, nil
}

// Login service for login.
func (s *UserService) Login(ctx context.Context, loginReq userModel.LoginReq, role string) (userModel.UserWithTokens, error) {
	var token string
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// check role exist or not.
		// do not check dynamically because every time can be deleted role in db, but fixed names can be in db: superadmin, admin, student
		roles, err := db.RolesRepo().GetRoleByRoleName(ctx, role)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		userDetails, err := db.UsersRepo().GetUserByUsername(ctx, loginReq.Username)
		if err != nil {
			return errlst.ParseErrors(err)
		}

		if userDetails.RoleID != roles.ID {
			return errlst.NewConflictError("provided roleID does not match with taken roleID from db.")
		}

		// Compare passwords.
		if err := utils.CheckAndComparePassword(loginReq.Password, userDetails.Password); err != nil {
			return errlst.ParseErrors(err)
		}

		// Generate token.
		token, err = s.mw.GenerateToken(userDetails.ID, userDetails.RoleID, userDetails.Username, roles.RoleName)
		if err != nil {
			return errlst.NewUnauthorizedError("cannot generate token with this user details")
		}

		return nil
	})
	if err != nil {
		return userModel.UserWithTokens{}, errlst.ParseErrors(err)
	}

	return userModel.UserWithTokens{Token: token}, nil
}
