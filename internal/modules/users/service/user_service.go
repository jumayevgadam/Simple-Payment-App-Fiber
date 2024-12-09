package service

import (
	"context"

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

// UserService manages buisiness logic for modules/user part of application.
type UserService struct {
	repo database.DataStore
}

// NewUserService creates and returns a new instance of UserRepository.
func NewUserService(repo database.DataStore) *UserService {
	return &UserService{repo: repo}
}

// CreateUser service insert a user into db and returns its id.
func (s *UserService) Register(ctx context.Context, request userModel.SignUpReq) (int, error) {
	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		return -1, errlst.ParseErrors(err)
	}
	request.Password = hashedPass

	userID, err := s.repo.UsersRepo().CreateUser(ctx, request.ToStorage(3))
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
		_, err = db.RolesRepo().GetRole(ctx, user.RoleID)
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
