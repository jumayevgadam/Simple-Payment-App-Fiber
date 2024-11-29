package middleware

import (
	"github.com/jumayevgadam/tsu-toleg/internal/config"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/cache"
	"github.com/jumayevgadam/tsu-toleg/internal/models/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/logger"
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
	cfg    *config.Config
	Logger logger.Logger
}

// NewMiddlewareManager func creates and returns a new instance TokenOps.
func NewMiddlewareManager(cfg *config.Config, redisRepo cache.Store, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, Logger: logger}
}
