package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

// RoleBasedMiddleware takes needed middleware permissions.
func RoleBasedMiddleware(mw *MiddlewareManager, permission string, dataStore database.DataStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// retrieve roles by permission
		roles, err := dataStore.RolesRepo().GetRolesByPermission(context.Background(), permission)
		if err != nil {
			return errlst.NewBadRequestError("can not find roles decided that permission")
		}

		fmt.Println("roles fetched with permission: ", permission, roles)

		roleMatched := false
		for _, role := range roles {
			fmt.Println(role.RoleName)
			Token := c.Cookies(role.RoleName + "_" + "token")
			fmt.Println("Tokens and roles:: ", Token, role.RoleName)
			if Token == "" {
				utils.ClearTokenCookie(c, mw.cfg, role.RoleName, Token)
				continue
			}

			// get token claims.
			claims, err := mw.ParseToken(Token)
			if err != nil {
				utils.ClearTokenCookie(c, mw.cfg, role.RoleName, Token)
				fmt.Println("invalid token for role:", role.RoleName, "Error", err)
				continue
			}

			if claims.RoleID == role.ID {
				c.Locals("userRoleID", claims.RoleID)
				c.Locals("userID", claims.ID)
				c.Locals("username", claims.UserName)
				c.Locals("role", claims.Role)
				fmt.Println("Middleware set locals: ", claims)

				roleMatched = true
				break
			}
		}

		if roleMatched {
			return c.Next()
		}

		return errlst.NewForbiddenError("access denied for role this permission")
	}
}

func AuthMiddleware(mw *MiddlewareManager, dataStore database.DataStore, permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		return nil
	}
}
