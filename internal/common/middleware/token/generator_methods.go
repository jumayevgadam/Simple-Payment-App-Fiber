package token

import (
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/cache"
)

// Ensure TokenOps implements the TokenGeneratorOps interface.
var _ TokenGeneratorOps = (*TokenOps)(nil)

// TokenGeneratorOps interface for generating tokens.
type TokenGeneratorOps interface {
	GenerateAccessToken(userID, roleID int, username string) (string, error)
	GenerateRefreshToken(userID, roleId int) (string, error)
	ParseAccessToken(accessToken string) (*AccessTokenClaims, error)
	ParseRefreshToken(refreshToken string) (*RefreshTokenClaims, error)
}

// JWTMaker struct takes all needed details for jwtToken from config.
type TokenOps struct {
	redisRepo cache.Store
	jwtOps config.JWTOps
}

// NewJWTMaker func creates and returns a new instance TokenOps.
func NewTokenOps(jwtOps config.JWTOps, redisRepo cache.Store) *TokenOps {
	return &TokenOps{jwtOps: jwtOps, redisRepo: redisRepo}
}
