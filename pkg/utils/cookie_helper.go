package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
)

// ClearAccessTokenCookie
func ClearTokenCookie(c *fiber.Ctx, cfg *config.Config, accessToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   false,
		Domain:   "",
	})
}

// SetAuthCookies
func SetAuthCookies(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		Domain:   "",
	})
}
