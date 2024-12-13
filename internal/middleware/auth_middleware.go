package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (mw *Manager) RoleBasedMiddleware(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from header.
		tokenString := c.Get("Authorization")

		// Validate the token.
		if tokenString == "" {
			return errlst.NewForbiddenError("authorization header not provided")
		}

		claims, err := mw.ParseToken(tokenString)
		if err != nil {
			return errlst.NewUnauthorizedError("can not parse token: err[mw.ParseToken]")
		}

		// Check for required permission.
		hasPermission := false
		for _, perm := range claims.Permissions {
			if perm == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return errlst.NewUnauthorizedError("access denied for this permission")
		}

		// Store user details in context.
		c.Locals("user_id", claims.ID)
		c.Locals("role_id", claims.RoleID)
		c.Locals("role_type", claims.Role)
		c.Locals("username", claims.UserName)

		return c.Next()
	}
}
