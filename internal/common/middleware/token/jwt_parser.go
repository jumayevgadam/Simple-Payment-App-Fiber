package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// ParseAccessToken method parses accessToken using claims.
func (tp *TokenOps) ParseAccessToken(accessToken string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check jwt signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errlst.ErrInvalidJWTMethod
		}
		return []byte(tp.jwtOps.AccessTokenSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errlst.NewUnauthorizedError("invalid access token")
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return nil, errors.New("error in type assertion for AccessTokenClaims")
	}

	return claims, nil
}

// ParseRefreshToken method parses refresh token taking claims.
func (tp *TokenOps) ParseRefreshToken(refreshToken string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check jwt signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errlst.ErrInvalidJWTMethod
		}

		return []byte(tp.jwtOps.RefreshTokenSecret), nil
	})
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return nil, errors.New("error in type assertion in this place")
	}

	return claims, nil
}
