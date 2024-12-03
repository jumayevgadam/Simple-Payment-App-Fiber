package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
)

// ClearAccessTokenCookie
func ClearTokenCookie(c *fiber.Ctx, cfg *config.Config, role, accessToken string) {
	CookieName := role + "_" + "token"
	c.Cookie(&fiber.Cookie{
		Name:     CookieName,
		Value:    accessToken,
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   false,
		Domain:   "",
	})
}

// SetAuthCookies
func SetAuthCookies(c *fiber.Ctx, role, token string) {
	CookieName := role + "_" + "token"
	c.Cookie(&fiber.Cookie{
		Name:     CookieName,
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
		Domain:   "",
	})
}
