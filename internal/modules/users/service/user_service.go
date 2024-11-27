package service

import (
	"context"
	"time"

	"github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/cache"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/models/abstract"
	userModel "github.com/jumayevgadam/tsu-toleg/internal/models/user"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/users"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst/tracing"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
	"go.opentelemetry.io/otel"
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
	jwtOps middleware.TokenGeneratorOps
	repo   database.DataStore
	cache  cache.Store
}

// NewUserService creates and returns a new instance of UserRepository.
func NewUserService(jwtOps middleware.TokenGeneratorOps, repo database.DataStore, cache cache.Store) *UserService {
	return &UserService{jwtOps: jwtOps, repo: repo, cache: cache}
}

// CreateUser service insert a user into db and returns its id.
func (s *UserService) CreateUser(ctx context.Context, request userModel.SignUpReq, role string) (int, error) {
	ctx, span := otel.Tracer("[UserService]").Start(ctx, "[CreateUser]")
	defer span.End()

	// check role that exist or not.
	roleID, exists := RoleMap[role]
	if !exists {
		return -1, errlst.NewBadRequestError("invalid role provided")
	}

	hashedPass, err := utils.HashPassword(request.Password)
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}
	request.Password = hashedPass

	userID, err := s.repo.UsersRepo().CreateUser(ctx, roleID, request.ToStorage())
	if err != nil {
		tracing.ErrorTracer(span, err)
		return -1, errlst.ParseErrors(err)
	}

	return userID, nil
}

// Login service for login.
func (s *UserService) Login(ctx context.Context, loginReq userModel.LoginReq, role string) (userModel.UserWithTokens, error) {
	ctx, span := otel.Tracer("[UserService]").Start(ctx, "[Login]")
	defer span.End()

	var userWithToken userModel.UserWithTokens
	err := s.repo.WithTransaction(ctx, func(db database.DataStore) error {
		// check role exist or not.
		roleID, exist := RoleMap[role]
		if !exist {
			return errlst.NewBadRequestError("invalid role provided")
		}

		userDAO, err := db.UsersRepo().GetUserByUsername(ctx, loginReq.Username)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		if userDAO.RoleID != roleID {
			return errlst.NewConflictError("provided roleID does not match with taken roleID from db.")
		}

		// Compare passwords
		if err := utils.CheckAndComparePassword(loginReq.Password, userDAO.Password); err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		// generate accessToken here
		accessToken, refreshToken, err := s.jwtOps.GenerateTokens(userDAO.ID, userDAO.RoleID, userDAO.Username)
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		// Put refresh token to the session
		err = s.cache.PutSession(ctx, abstract.SessionArgument{
			SessionPrefix: "refresh_token",
			RoleID:        userDAO.RoleID,
			UserID:        userDAO.ID,
			RefreshToken:  refreshToken,
			UserName:      userDAO.Username,
			ExpiresAt:     time.Duration(60 * time.Minute),
		})
		if err != nil {
			tracing.ErrorTracer(span, err)
			return errlst.ParseErrors(err)
		}

		// Putting all values to UserWithToken model
		userWithToken.AccessToken = accessToken
		userWithToken.RefreshToken = refreshToken
		userWithToken.User = userDAO.ToServer()

		return nil
	})
	if err != nil {
		tracing.ErrorTracer(span, err)
		return userModel.UserWithTokens{}, errlst.ParseErrors(err)
	}

	return userWithToken, nil
}
