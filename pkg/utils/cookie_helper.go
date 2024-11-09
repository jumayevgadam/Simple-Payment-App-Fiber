package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
)

// TokenCookies model is
type TokenCookies struct {
	AccessTokenCookie  *fiber.Cookie
	RefreshTokenCookie *fiber.Cookie
}

// SetAuthCookies
func SetAuthCookies(cfg *config.JWTOps, accessToken, refreshToken string) TokenCookies {
	accessTokenCookie := &fiber.Cookie{
		Name:     os.Getenv("ACCESS_TOKEN_NAME"),
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(cfg.AccessTokenExpiryTime) * time.Minute),
		Secure:   false,
		HTTPOnly: true,
	}

	refreshTokenCookie := &fiber.Cookie{
		Name:     os.Getenv("REFRESH_TOKEN_NAME"),
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(cfg.RefreshTokenExpiryTime) * time.Minute),
		Secure:   false,
		HTTPOnly: true,
	}

	return TokenCookies{
		AccessTokenCookie:  accessTokenCookie,
		RefreshTokenCookie: refreshTokenCookie,
	}
}
