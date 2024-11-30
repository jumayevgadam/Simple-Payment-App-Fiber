package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
)

// ClearAccessTokenCookie
func ClearTokenCookie(c *fiber.Ctx, cfg *config.Config, accessToken string) {
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
		Secure:   false,
		HTTPOnly: false,
		Domain:   "localhost",
	})
}
