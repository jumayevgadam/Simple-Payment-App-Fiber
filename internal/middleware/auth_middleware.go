package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (mw *Manager) RoleBasedMiddleware(requiredRole string, requiredRoleID int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from header.
		tokenString := c.Get("Authorization")

		// Validate the token.
		if tokenString == "" {
			return errlst.NewForbiddenError("authorization header not provided")
		}

		claims, err := mw.ParseToken(tokenString)
		if err != nil {
			errlst.NewUnauthorizedError("cannot parse token: [mw.ParseToken]")
		}

		if claims.Role == "superadmin" || claims.RoleID == 1 {
			// Log the bypass for auditing
			mw.Logger.Info(fmt.Sprintf("Superadmin access granted: user_id [%d], username [%s]", claims.ID, claims.UserName))

			c.Locals("user_id", claims.ID)
			c.Locals("role_id", claims.RoleID)
			c.Locals("role_type", claims.Role)
			c.Locals("username", claims.UserName)

			return c.Next()
		}

		if claims.Role != requiredRole || claims.RoleID != requiredRoleID {
			return errlst.NewForbiddenError(fmt.Sprintf(
				"access denied: required role [%s] and role_id [%d]",
				requiredRole, requiredRoleID,
			))
		}

		c.Locals("user_id", claims.ID)
		c.Locals("role_id", claims.RoleID)
		c.Locals("role_type", claims.Role)
		c.Locals("username", claims.UserName)

		return c.Next()
	}
}

func GetUserIDFromFiberContext(ctx *fiber.Ctx) (int, error) {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return 0, errlst.NewUnauthorizedError("user_id not found in context")
	}
	return userID, nil
}

func GetStudentIDFromFiberContext(ctx *fiber.Ctx) (int, error) {
	roleType, _ := ctx.Locals("role_type").(string)
	if roleType == "superadmin" || roleType == "student" {
		return GetUserIDFromFiberContext(ctx)
	}

	return 0, errlst.NewUnauthorizedError("only students and superadmin can perform this action")
}

func GetAdminIDFromFiberContext(ctx *fiber.Ctx) (int, error) {
	roleType, _ := ctx.Locals("role_type").(string)
	if roleType == "superadmin" || roleType == "admin" {
		return GetUserIDFromFiberContext(ctx)
	}

	return 0, errlst.NewUnauthorizedError("only admins and superadmin can perform this action")
}
