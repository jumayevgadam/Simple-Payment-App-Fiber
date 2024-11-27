package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
)

// ClearRefreshTokenCookie
func ClearRefreshTokenCookie(c *fiber.Ctx, cfg *config.Config, refreshToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     cfg.JWT.RefreshTokenName,
		Value:    refreshToken,
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		Domain:   "localhost",
	})
}

// ClearAccessTokenCookie
func ClearAccessTokenCookie(c *fiber.Ctx, cfg *config.Config, accessToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     cfg.JWT.AccessTokenName,
		Value:    accessToken,
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		Domain:   "localhost",
	})
}

// SetAuthCookies
func SetAuthCookies(c *fiber.Ctx, cfg *config.Config, accessToken, refreshToken string) {
	accessTokenCookie := &fiber.Cookie{
		Name:     cfg.JWT.AccessTokenName,
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(cfg.JWT.AccessTokenExpiryTime) * time.Minute),
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	}

	refreshTokenCookie := &fiber.Cookie{
		Name:     cfg.JWT.RefreshTokenName,
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(cfg.JWT.RefreshTokenExpiryTime) * time.Minute),
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	}

	c.Cookie(accessTokenCookie)
	c.Cookie(refreshTokenCookie)
}
