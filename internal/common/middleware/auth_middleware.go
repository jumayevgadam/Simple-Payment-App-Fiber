package middleware

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// create dynamic roles.
var RoleMap = map[int]string{
	1: "SuperAdmin",
	2: "Admin",
	3: "Student",
}

// RoleBasedMiddleware takes needed middleware permissions.
func RoleBasedMiddleware(mw *MiddlewareManager, allowedRoles ...int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from cookie.
		accessToken := c.Cookies(os.Getenv("ACCESS_TOKEN_NAME"))
		refreshToken := c.Cookies(os.Getenv("REFRESH_TOKEN_NAME"))

		if accessToken == "" && refreshToken == "" {
			return errlst.NewUnauthorizedError("missing access and refresh tokens")
		}

		// get claims.
		claims, err := mw.ParseAccessToken(accessToken)
		if err != nil {
			if err == errlst.ErrTokenExpired && refreshToken != "" {
				newAccessToken, err := HandleNewAccessToken(c, mw, refreshToken)
				if err != nil {
					return errlst.NewBadRequestError("cannot create a new accessToken")
				}
				accessToken = newAccessToken
				claims, err = mw.ParseAccessToken(accessToken)
				if err != nil {
					return errlst.NewUnauthorizedError("invalid refreshed access token")
				}
			} else {
				return errlst.Response(c, err)
			}
		}

		for _, roleId := range allowedRoles {
			if claims.RoleID == roleId {
				// Attach user claims to the context for further usage.
				c.Locals("user_claims", claims)
				return c.Next()
			}
		}

		roleName, exists := RoleMap[claims.RoleID]
		if !exists {
			roleName = "Unknown role(" + strconv.Itoa(claims.RoleID) + ")"
		}

		// Proceed to the next middleware or handler
		return errlst.NewForbiddenError("access denied for role: " + roleName)
	}
}

// HandleNewAccessToken method generates a new access token and delete old refresh token from redisDB.
func HandleNewAccessToken(c *fiber.Ctx, mw *MiddlewareManager, refreshToken string) (string, error) {
	// verify refresh token from redisDB.

	// Retrieve user claims.

	// generate new tokens.

	// update cookies with new tokens

	// delete old refresh token from redis and store the new one.

	// return newAccessToken
	return "", nil
}
