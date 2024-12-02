package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/models/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
)

// Ensure TokenOps implements the TokenGeneratorOps interface.
var _ TokenGeneratorOps = (*MiddlewareManager)(nil)

// TokenGeneratorOps interface for generating tokens.
type TokenGeneratorOps interface {
	GenerateToken(userID, roleID int, username string, expirationTime time.Duration) (string, error)
	ParseToken(accessToken string) (*token.AccessTokenClaims, error)
}

// MiddlewareManager struct takes all needed details for jwtToken from config.
type MiddlewareManager struct {
	cfg    *config.Config
	Logger logger.Logger
}

// NewMiddlewareManager func creates and returns a new instance TokenOps.
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, Logger: logger}
}

// GenerateAccessToken method for creating access token.
func (mw *MiddlewareManager) GenerateToken(userID, roleID int, username string, expirationTime time.Duration) (string, error) {
	accessTokenclaims := token.AccessTokenClaims{
		ID:       userID,
		RoleID:   roleID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime * time.Minute)),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenclaims).SignedString([]byte(mw.cfg.JWT.AccessTokenName))
	if err != nil {
		return "", errlst.NewInternalServerError("error creating accessToken")
	}

	return accessToken, nil
}

// ParseAccessToken method parses accessToken using claims.
func (mw *MiddlewareManager) ParseToken(accessToken string) (*token.AccessTokenClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(accessToken, &token.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check jwt signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errlst.ErrInvalidJWTMethod
		}

		return []byte(mw.cfg.JWT.AccessTokenSecret), nil
	})
	mw.Logger.Info(tokenStr)
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
