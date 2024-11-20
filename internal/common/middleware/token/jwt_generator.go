package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// GenerateAccessToken method for creating access token.
func (tg *TokenOps) GenerateAccessToken(userID, roleID int, username string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(tg.jwtOps.AccessTokenExpiryTime) * time.Minute)
	claims := AccessTokenClaims{
		ID:       userID,
		RoleID:   roleID,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenStr, err := accessToken.SignedString([]byte(tg.jwtOps.AccessTokenSecret))
	if err != nil {
		return "", errlst.NewUnauthorizedError("error creating accessTokenStr")
	}

	return accessTokenStr, nil
}

// GenerateRefreshToken method for creating refresh token.
func (tg *TokenOps) GenerateRefreshToken(userID, roleID int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(tg.jwtOps.RefreshTokenExpiryTime) * time.Minute)
	claims := RefreshTokenClaims{
		ID:     userID,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// create refresh token with claims.
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenStr, err := refreshToken.SignedString([]byte(tg.jwtOps.AccessTokenSecret))
	if err != nil {
		return "", errors.New("error in taking refresh token string")
	}

	return refreshTokenStr, nil
}
