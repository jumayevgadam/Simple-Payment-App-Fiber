package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// RoleBasedMiddleware middleware system built with permissions.
func (mw *MiddlewareManager) RoleBasedMiddleware(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from header.
		bearerHeader := c.Get("Authorization")

		mw.Logger.Infof("auth middleware bearerHeader %s", bearerHeader)

		if bearerHeader != "" {
			headerParts := strings.Split(bearerHeader, " ")
			if len(headerParts) != 2 {
				mw.Logger.Error("[authRoleBasedMiddleware], headerParts, len(headerParts) != 2")
				return errlst.Response(c, errlst.NewUnauthorizedError(errlst.ErrUnauthorized))
			}

			tokenString := headerParts[1]

			claims, err := mw.ParseToken(tokenString)
			if err != nil {
				mw.Logger.Errorf("middleware ParseToken: %v", err.Error())
				return errlst.Response(c, err)
			}

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

			c.Locals("user_id", claims.ID)
			c.Locals("role_id", claims.RoleID)
			c.Locals("role_type", claims.Role)
			c.Locals("username", claims.UserName)
			return c.Next()
		}

		return errlst.NewForbiddenError("authorization header not provided")
	}
}
