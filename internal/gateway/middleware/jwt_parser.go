package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jumayevgadam/tsu-toleg/internal/models/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// ParseAccessToken method parses accessToken using claims.
func (mw *MiddlewareManager) ParseAccessToken(accessToken string) (*token.AccessTokenClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(accessToken, &token.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check jwt signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errlst.ErrInvalidJWTMethod
		}
		return []byte(mw.cfg.JWT.AccessTokenSecret), nil
	})
	if err != nil || !tokenStr.Valid {
		return nil, errlst.NewUnauthorizedError("invalid access token")
	}

	claims, ok := tokenStr.Claims.(*token.AccessTokenClaims)
	if !ok {
		return nil, errlst.ErrInvalidJWTClaims
	}

	return claims, nil
}

// ParseRefreshToken method parses refresh token taking claims.
func (mw *MiddlewareManager) ParseRefreshToken(refreshToken string) (*token.RefreshTokenClaims, error) {
	tokenStr, err := jwt.ParseWithClaims(refreshToken, &token.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check jwt signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errlst.ErrInvalidJWTMethod
		}

		return []byte(mw.cfg.JWT.RefreshTokenSecret), nil
	})
	if err != nil {
		return nil, errlst.ParseErrors(err)
	}

	claims, ok := tokenStr.Claims.(*token.RefreshTokenClaims)
	if !ok {
		return nil, errlst.ErrInvalidJWTClaims
	}

	return claims, nil
}
