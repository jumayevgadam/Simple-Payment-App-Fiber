package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/config"
)

// SetAuthCookies
func SetAuthCookies(c *fiber.Ctx, cfg *config.Config, accessToken, refreshToken string) {
	accessTokenCookie := &fiber.Cookie{
		Name:     os.Getenv("ACCESS_TOKEN_NAME"),
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(cfg.JWT.AccessTokenExpiryTime) * time.Minute),
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	}

	refreshTokenCookie := &fiber.Cookie{
		Name:     os.Getenv("REFRESH_TOKEN_NAME"),
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
