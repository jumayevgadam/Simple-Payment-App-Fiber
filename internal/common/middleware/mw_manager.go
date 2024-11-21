package middleware

import (
	"github.com/jumayevgadaym/tsu-toleg/internal/common/middleware/token"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/infrastructure/cache"
)

// Ensure TokenOps implements the TokenGeneratorOps interface.
var _ TokenGeneratorOps = (*MiddlewareManager)(nil)

// TokenGeneratorOps interface for generating tokens.
type TokenGeneratorOps interface {
	GenerateTokens(userID, roleID int, username string) (string, string, error)
	ParseAccessToken(accessToken string) (*token.AccessTokenClaims, error)
	ParseRefreshToken(refreshToken string) (*token.RefreshTokenClaims, error)
}

// MiddlewareManager struct takes all needed details for jwtToken from config.
type MiddlewareManager struct {
	redisRepo cache.Store
	cfg       *config.Config
}

// NewMiddlewareManager func creates and returns a new instance TokenOps.
func NewMiddlewareManager(cfg *config.Config, redisRepo cache.Store) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, redisRepo: redisRepo}
}
