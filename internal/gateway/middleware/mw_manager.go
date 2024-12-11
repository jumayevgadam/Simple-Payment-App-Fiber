package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/models/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// Ensure TokenOps implements the TokenGeneratorOps interface.
var _ TokenGeneratorOps = (*Manager)(nil)

// TokenGeneratorOps interface for generating tokens.
type TokenGeneratorOps interface {
	GenerateToken(userID, roleID int, username, rolename string, permissions []string) (string, error)
	ParseToken(accessToken string) (*token.AccessTokenClaims, error)
}

// MiddlewareManager struct takes all needed details for jwtToken from config.
type Manager struct {
	cfg    *config.Config
	Logger logger.Logger
}

// NewMiddlewareManager func creates and returns a new instance TokenOps.
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *Manager {
	return &Manager{cfg: cfg, Logger: logger}
}

// GenerateAccessToken method for creating access token.
func (mw *Manager) GenerateToken(userID, roleID int, username, role string, permissions []string) (string, error) {
	accessTokenclaims := token.AccessTokenClaims{
		ID:          userID,
		RoleID:      roleID,
		UserName:    username,
		Role:        role,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.TokenExpiryTime * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenclaims)

	tokenStr, err := token.SignedString([]byte(mw.cfg.JWT.TokenSecret))
	if err != nil {
		return "", errlst.NewUnauthorizedError("cannot get token string" + err.Error())
	}

	return tokenStr, nil
}

// ParseAccessToken method parses accessToken using claims.
func (mw *Manager) ParseToken(accessToken string) (*token.AccessTokenClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(
		accessToken, &token.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			// check jwt signing method.
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errlst.ErrInvalidJWTMethod
			}

			return []byte(mw.cfg.JWT.TokenSecret), nil
		})
	if err != nil {
		return nil, errlst.NewUnauthorizedError("invalid access token")
	}

	claims, ok := tokenStr.Claims.(*token.AccessTokenClaims)
	if !ok {
		mw.Logger.Info("error in tokenStr.Claims...")
		return nil, errlst.ErrInvalidJWTClaims
	}

	return claims, nil
}
