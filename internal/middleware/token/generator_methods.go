package token

import (
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
)

var _ TokenGeneratorOps = (*TokenOps)(nil)

// TokenGeneratorOps interface for generating tokens
type TokenGeneratorOps interface {
	GenerateAccessToken(userID, roleID int, username string) (string, error)
	GenerateRefreshToken(userID, roleId int) (string, error)
	ParseAccessToken(accessToken string) (*AccessTokenClaims, error)
	ParseRefreshToken(refreshToken string) (*RefreshTokenClaims, error)
}

// JWTMaker struct is
type TokenOps struct {
	jwtOps config.JWTOps
}

// NewJWTMaker func is
func NewTokenOps(jwtOps config.JWTOps) *TokenOps {
	return &TokenOps{jwtOps: jwtOps}
}
