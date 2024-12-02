package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// Check this place. edit with silly codes.
// RoleBasedMiddleware takes needed middleware permissions.
func RoleBasedMiddleware(mw *MiddlewareManager, permission string, dataStore database.DataStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from cookie.
		accessToken := c.Cookies("access_token")
		if accessToken == "" {
			return errlst.NewUnauthorizedError("missing access and refresh tokens")
		}
		fmt.Println("accessToken gelman yatyr diyyanmi " + accessToken)

		// get token claims.
		claims, err := mw.ParseToken(accessToken)
		if err != nil {
			// error occured in this place
			return errlst.NewUnauthorizedError(err.Error())
		}

		// Fetch permissions for the user's role
		roleIDs, err := dataStore.RolesRepo().GetRolesByPermission(context.Background(), permission)
		if err != nil {
			return errlst.ErrNoSuchRole
		}

		hasPermission := false
		for _, roleID := range roleIDs {
			if claims.RoleID == roleID {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return errlst.NewForbiddenError("access denied for role  with those permissions")
		}

		c.Locals("userRole", claims.RoleID)
		c.Locals("userID", claims.ID)
		c.Locals("username", claims.UserName)
		return c.Next()
	}
}
