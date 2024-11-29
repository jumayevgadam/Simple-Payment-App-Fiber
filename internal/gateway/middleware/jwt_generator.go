package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/models/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// GenerateAccessToken method for creating access token.
func (mw *MiddlewareManager) GenerateTokens(userID, roleID int, username string) (string, string, error) {
	accessTokenclaims := token.AccessTokenClaims{
		ID:               userID,
		RoleID:           roleID,
		UserName:         username,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	refreshTokenClaims := token.RefreshTokenClaims{
		ID:               userID,
		RoleID:           roleID,
		UserName:         username,
		RegisteredClaims: jwt.RegisteredClaims{},
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
