package middleware

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/models/token"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadam/tsu-toleg/pkg/utils"
)

// create dynamic roles.
var RoleMap = map[int]string{
	1: "SuperAdmin",
	2: "Admin",
	3: "Student",
}

// Check this place. edit with silly codes.
// RoleBasedMiddleware takes needed middleware permissions.
func RoleBasedMiddleware(mw *MiddlewareManager, permission string, dataStore database.DataStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from cookie.
		accessToken := c.Cookies(os.Getenv("ACCESS_TOKEN_NAME"))
		refreshToken := c.Cookies(os.Getenv("REFRESH_TOKEN_NAME"))

		if accessToken == "" && refreshToken == "" {
			return errlst.NewUnauthorizedError("missing access and refresh tokens")
		}

		// get claims.
		var claims *token.AccessTokenClaims
		var err error
		claims, err = mw.ParseAccessToken(accessToken)
		if err != nil {
			mw.Logger.Error("error in 41") // error in here
			if errors.Is(err, errlst.ErrTokenExpired) && refreshToken != "" {
				// clear accesstoken from cookie
				utils.ClearAccessTokenCookie(c, mw.cfg, accessToken)

				claims, err = mw.ParseAccessToken(accessToken)
				if err != nil {
					mw.Logger.Error("error in 52")
					return errlst.NewUnauthorizedError("invalid refreshed access token")
				}
			} else {
				mw.Logger.Error("Token error not expired")
				return errlst.Response(c, err)
			}
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
			roleName, exists := RoleMap[claims.RoleID]
			if !exists {
				roleName = "Unknown role(" + strconv.Itoa(claims.RoleID) + ")"
			}
			return errlst.NewForbiddenError("access denied for role: " + roleName)
		}

		c.Locals("user_claims", claims)
		return c.Next()
	}
}
