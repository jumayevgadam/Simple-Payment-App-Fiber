package service

import (
	"context"

	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/middleware"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
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
