package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/middleware/token"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
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
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenclaims)

	accessTokenStr, err := accessToken.SignedString([]byte(mw.cfg.JWT.AccessTokenName))
	if err != nil {
		return "", "", errlst.NewUnauthorizedError("error creating accessTokenStr")
	}

	refreshTokenClaims := token.RefreshTokenClaims{
		ID:       userID,
		RoleID:   roleID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExTime),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenStr, err := refreshToken.SignedString([]byte(mw.cfg.JWT.RefreshTokenSecret))
	if err != nil {
		return "", "", errlst.NewUnauthorizedError("error creating refreshTokenStr")
	}

	return accessTokenStr, refreshTokenStr, nil
}
