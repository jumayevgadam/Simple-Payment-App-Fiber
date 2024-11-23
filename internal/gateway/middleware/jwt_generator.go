package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// GenerateAccessToken method for creating access token.
func (mw *MiddlewareManager) GenerateTokens(userID, roleID int, username string) (string, string, error) {
	accessTokenExTime := time.Now().Add(time.Duration(mw.cfg.JWT.AccessTokenExpiryTime) * time.Minute)
	refreshTokenExTime := time.Now().Add(time.Duration(mw.cfg.JWT.RefreshTokenExpiryTime) * time.Minute)
	accessTokenclaims := token.AccessTokenClaims{
		ID:       userID,
		RoleID:   roleID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExTime),
		},
	}
	refreshTokenClaims := token.RefreshTokenClaims{
		ID:       userID,
		RoleID:   roleID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExTime),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenclaims).SignedString([]byte(mw.cfg.JWT.AccessTokenName))
	if err != nil {
		return "", "", errlst.NewInternalServerError("error creating accessToken")
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(mw.cfg.JWT.RefreshTokenSecret))
	if err != nil {
		return "", "", errlst.NewInternalServerError("error creating refreshToken")
	}

	return accessToken, refreshToken, nil
}
