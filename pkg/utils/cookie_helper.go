package utils

import (
	"os"
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
		HTTPOnly: false,
		Secure:   false,
		Domain:   "",
	})
}

// ClearAccessTokenCookie
func ClearAccessTokenCookie(c *fiber.Ctx, cfg *config.Config, accessToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     cfg.JWT.AccessTokenName,
		Value:    accessToken,
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: false,
		Secure:   false,
		Domain:   "",
	})
}

// SetAuthCookies
func SetAuthCookies(c *fiber.Ctx, accessToken, refreshToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("ACCESS_TOKEN_NAME"),
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(5 * time.Minute),
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})

	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("REFRESH_TOKEN_NAME"),
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})
}
