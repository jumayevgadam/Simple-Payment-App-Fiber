package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/common/middleware/token"
	"github.com/jumayevgadaym/tsu-toleg/internal/config"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
)

// create dynamic roles.
var RoleMap = map[int]string{
	1: "SuperAdmin",
	2: "Admin",
	3: "Student",
}

// RoleBasedMiddleware takes needed middleware permissions.
func RoleBasedMiddleware(cfg config.JWTOps, allowedRoles ...int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the JWT token from cookie.
		accessToken := c.Cookies(cfg.AccessTokenName)
		if accessToken == "" {
			return errlst.NewUnauthorizedError("missing access token in cookies")
		}

		// get claims.
		var tokenOps *token.TokenOps
		claims, err := tokenOps.ParseAccessToken(accessToken)
		if err != nil {
			return errlst.Response(c, err)
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

// // GetSuperAdminMiddleware for superadmin.
// func GetSuperAdminMiddleware(cfg config.JWTOps, tokenOps *token.TokenOps) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// read the token from cookie.
// 		accessToken := c.Cookies(cfg.AccessTokenName)
// 		if accessToken == "" {
// 			return c.Status(fiber.StatusNotFound).JSON("access token is not found in cookie")
// 		}

// 		// get claims.
// 		claims, err := tokenOps.ParseAccessToken(accessToken)
// 		if err != nil {
// 			return errlst.Response(c, err)
// 		}

// 		if claims.RoleID != 1 {
// 			return c.Status(fiber.StatusForbidden).JSON("user is not SuperAdmin")
// 		}

// 		// pass the payload/claims down the fiber context.
// 		c.Locals("SuperAdmin-claims", claims)
// 		return c.Next()
// 	}
// }

// // GetAdminMiddleware for admin middleware.
// func GetAdminMiddleware(cfg config.JWTOps, tokenOps *token.TokenOps) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Read the token from cookie.
// 		accessToken := c.Cookies(cfg.AccessTokenName)
// 		if accessToken == "" {
// 			return c.Status(fiber.StatusNotFound).JSON("access token is not found in cookie")
// 		}

// 		// get claims.
// 		claims, err := tokenOps.ParseAccessToken(accessToken)
// 		if err != nil {
// 			return errlst.Response(c, err)
// 		}

// 		if claims.RoleID != 2 {
// 			return c.Status(fiber.StatusForbidden).JSON("user is not admin")
// 		}

// 		// pass the payload/claims down the fiber context.
// 		c.Locals("Admin-claims", claims)
// 		return c.Next()
// 	}
// }

// // GetStudentMiddleware for student middleware.
// func GetStudentMiddleware(cfg config.JWTOps, tokenOps *token.TokenOps) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Read the token from cookie.
// 		accessToken := c.Cookies(cfg.AccessTokenName)
// 		if accessToken == "" {
// 			return c.Status(fiber.StatusNotFound).JSON("access token is not found in cookie")
// 		}

// 		// get claims.
// 		claims, err := tokenOps.ParseAccessToken(accessToken)
// 		if err != nil {
// 			return errlst.Response(c, err)
// 		}

// 		if claims.RoleID != 3 {
// 			return c.Status(fiber.StatusForbidden).JSON("user is not a student")
// 		}

// 		// pass the payload/claims down the fiber context.
// 		c.Locals("Student-claims", claims)
// 		return c.Next()
// 	}
// }
