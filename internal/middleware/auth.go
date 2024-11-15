package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/internal/middleware/token"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// GetSuperAdminMiddleware for superadmin.
func GetSuperAdminMiddleware(cfg config.JWTOps, tokenOps *token.TokenOps) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// read the token from cookie
		accessToken := c.Cookies(cfg.AccessTokenName)
		if accessToken == "" {
			return c.Status(fiber.StatusNotFound).JSON("access token is not found in cookie")
		}

		// get claims
		claims, err := tokenOps.ParseAccessToken(accessToken)
		if err != nil {
			return errlst.Response(c, err)
		}

		if claims.RoleID != 1 {
			return c.Status(fiber.StatusForbidden).JSON("user is not SuperAdmin")
		}

		// pass the payload/claims down the fiber context
		c.Locals("SuperAdmin-claims", claims)
		return c.Next()
	}
}

// GetAdminMiddleware for admin middleware.
func GetAdminMiddleware(cfg config.JWTOps, tokenOps *token.TokenOps) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Read the token from cookie
		accessToken := c.Cookies(cfg.AccessTokenName)
		if accessToken == "" {
			return c.Status(fiber.StatusNotFound).JSON("access token is not found in cookie")
		}

		// get claims
		claims, err := tokenOps.ParseAccessToken(accessToken)
		if err != nil {
			return errlst.Response(c, err)
		}

		if claims.RoleID != 2 {
			return c.Status(fiber.StatusForbidden).JSON("user is not admin")
		}

		// pass the payload/claims down the fiber context
		c.Locals("Admin-claims", claims)
		return c.Next()
	}
}

// GetStudentMiddleware for student middleware.
func GetStudentMiddleware(cfg config.JWTOps, tokenOps *token.TokenOps) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Read the token from cookie
		accessToken := c.Cookies(cfg.AccessTokenName)
		if accessToken == "" {
			return c.Status(fiber.StatusNotFound).JSON("access token is not found in cookie")
		}

		// get claims
		claims, err := tokenOps.ParseAccessToken(accessToken)
		if err != nil {
			return errlst.Response(c, err)
		}

		if claims.RoleID != 3 {
			return c.Status(fiber.StatusForbidden).JSON("user is not a student")
		}

		// pass the payload/claims down the fiber context
		c.Locals("Student-claims", claims)
		return c.Next()
	}
}
