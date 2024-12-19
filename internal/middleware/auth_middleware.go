package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

func (mw *Manager) RoleBasedMiddleware(allowedRoles []string, allowedRoleIDs []int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from the Authorization header.
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return errlst.Response(c, errlst.NewForbiddenError(errlst.ErrAuthorizationHeaderNotProvided.Error()))
		}

		// Parse the token to extract claims.
		claims, err := mw.ParseToken(tokenString)
		if err != nil {
			return errlst.Response(c, err)
		}

		// Validate role and role ID.
		isRoleAllowed := containsString(allowedRoles, claims.Role)
		isRoleIDAllowed := containsInt(allowedRoleIDs, claims.RoleID)

		if !isRoleAllowed || !isRoleIDAllowed {
			return errlst.NewForbiddenError(fmt.Sprintf(
				"access denied: required roles %v and role IDs %v, but got role [%s] and role ID [%d]",
				allowedRoles, allowedRoleIDs, claims.Role, claims.RoleID,
			))
		}

		// Add user details to the context for downstream handlers.
		c.Locals("user_id", claims.ID)
		c.Locals("role_id", claims.RoleID)
		c.Locals("role_type", claims.Role)
		c.Locals("username", claims.UserName)

		return c.Next()
	}
}

// Helper function to check if a string exists in a slice.
func containsString(slice []string, item string) bool {
	for _, str := range slice {
		if strings.EqualFold(str, item) { // Case-insensitive comparison.
			return true
		}
	}

	return false
}

// Helper function to check if an int exists in a slice.
func containsInt(slice []int, item int) bool {
	for _, num := range slice {
		if num == item {
			return true
		}
	}

	return false
}

// GetStudentIDFromContext extracts the student ID from the Fiber context.
func GetStudentIDFromFiberContext(ctx *fiber.Ctx) (int, error) {
	// Extract role type from context.
	roleType, ok := ctx.Locals("role_type").(string)
	if !ok {
		return 0, errlst.NewUnauthorizedError("role_type not found in context")
	}

	// Only allow superadmins and students to retrieve the student ID.
	if roleType != "superadmin" && roleType != "student" {
		return 0, errlst.NewForbiddenError(fmt.Sprintf("access denied for role_type: %s", roleType))
	}

	// Extract user ID from context (assumed to be student ID for students).
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return 0, errlst.NewUnauthorizedError("user_id not found in context or invalid format")
	}

	return userID, nil
}
