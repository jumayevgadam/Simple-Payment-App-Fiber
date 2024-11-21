package middleware

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadaym/tsu-toleg/internal/models/abstract"
	"github.com/jumayevgadaym/tsu-toleg/pkg/errlst"
	"github.com/jumayevgadaym/tsu-toleg/pkg/utils"
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
	userID, err := mw.redisRepo.GetSession(c.Context(), abstract.SessionArgument{
		SessionPrefix: "refresh_token",
		RefreshToken:  refreshToken,
		UserID:        0,
	})
	if err != nil {
		return "", errlst.NewUnauthorizedError("invalid or expired refresh token")
	}

	var sessionClaims abstract.SessionArgument
	if err := json.Unmarshal(userID, &sessionClaims); err != nil {
		return "", errlst.NewInternalServerError("failed to parse session claims")
	}

	// generate new tokens.
	newAccessToken, newRefreshToken, err := mw.GenerateTokens(sessionClaims.UserID, sessionClaims.RoleID, sessionClaims.UserName)
	if err != nil {
		return "", errlst.NewInternalServerError("error creating access, refresh token in HandleNewAccessToken")
	}

	// delete old refresh token from redis and store the new one.
	err = mw.redisRepo.DelSession(c.Context(), abstract.SessionArgument{
		SessionPrefix: "refresh_token",
		RefreshToken:  refreshToken,
		UserID:        sessionClaims.UserID,
	})
	if err != nil {
		return "", errlst.NewInternalServerError("error deleting old refresh token key from redisDB")
	}

	// store new refresh token in redis.
	err = mw.redisRepo.PutSession(c.Context(), abstract.SessionArgument{
		SessionPrefix: "refresh_token",
		UserID:        sessionClaims.UserID,
		RefreshToken:  sessionClaims.RefreshToken,
		ExpiresAt:     time.Duration(mw.cfg.JWT.RefreshTokenExpiryTime),
	})
	if err != nil {
		return "", errlst.NewInternalServerError("error setting new refresh token in redisDB")
	}

	// update cookies with new values
	utils.SetAuthCookies(c, mw.cfg, newAccessToken, newRefreshToken)

	// return newAccessToken
	return newAccessToken, nil
}
