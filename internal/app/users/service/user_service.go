package service

import (
	"context"

	"github.com/jumayevgadaym/tsu-toleg/internal/app/users"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadaym/tsu-toleg/internal/middleware/token"
	userModel "github.com/jumayevgadaym/tsu-toleg/internal/models/user"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadaym/tsu-toleg/pkg/utils"
	"go.opentelemetry.io/otel"
)

// Ensure UserService implements the users.Service interface.
var (
	_ users.Service = (*UserService)(nil)
)

// UserService manages buisiness logic for app/user part of application.
type UserService struct {
	jwtOps token.TokenGeneratorOps
	repo   database.DataStore
}

// NewUserService creates and returns a new instance of UserRepository.
func NewUserService(jwtOps token.TokenGeneratorOps, repo database.DataStore) *UserService {
	return &UserService{jwtOps: jwtOps, repo: repo}
}

// CreateUser service insert a user into db and returns its id.
func (s *UserService) CreateUser(ctx context.Context, request userModel.SignUpReq) (int, error) {
	ctx, span := otel.Tracer("[UserService]").Start(ctx, "[CreateUser]")
	defer span.End()

	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}
	request.Password = hashedPass

	userID, err := s.repo.UsersRepo().CreateUser(ctx, request.ToStorage())
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}

	return userID, nil
}

// Login service for login.
func (s *UserService) Login(ctx context.Context, loginReq userModel.LoginReq) (userModel.UserWithTokens, error) {
	ctx, span := otel.Tracer("[UserService]").Start(ctx, "[Login]")
	defer span.End()

	var userWithToken userModel.UserWithTokens
	if err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		userDAO, err := db.UsersRepo().GetUserByUsername(ctx, loginReq.Username)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		// Compare passwords
		if err := utils.CheckAndComparePassword(loginReq.Password, userDAO.Password); err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		// generate accessToken here
		accessToken, err := s.jwtOps.GenerateAccessToken(userDAO.ID, userDAO.RoleID, userDAO.Username)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}
		// generate refresh token here
		refreshToken, err := s.jwtOps.GenerateRefreshToken(userDAO.ID, userDAO.RoleID)
		if err != nil {
			tracing.ErrorTracer(span, err)
		}

		// Putting all values to UserWithToken model
		userWithToken.AccessToken = accessToken
		userWithToken.RefreshToken = refreshToken
		userWithToken.User = userDAO.ToServer()

		return nil
	}); err != nil {
		tracing.ErrorTracer(span, err)
		return userModel.UserWithTokens{}, errlst.ParseErrors(err)
	}

	return userWithToken, nil
}
