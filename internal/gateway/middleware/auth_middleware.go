package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/models/abstract"
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
				newAccessToken, _, err := HandleNewAccessToken(c, mw, refreshToken)
				if err != nil {
					return errlst.NewBadRequestError("cannot create a new accessToken")
				}
				accessToken = newAccessToken

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

// HandleNewAccessToken method generates a new access token and delete old refresh token from redisDB.
func HandleNewAccessToken(c *fiber.Ctx, mw *MiddlewareManager, refreshToken string) (string, string, error) {
	refreshTokenClaims, err := mw.ParseRefreshToken(refreshToken)
	if err != nil {
		mw.Logger.Error(err)
		return "", "", errlst.NewUnauthorizedError("error verifying refresh token")
	}
	mw.Logger.Info(refreshTokenClaims.UserName)

	sessionArgument := abstract.SessionArgument{
		SessionPrefix: "refresh_token",
		UserID:        refreshTokenClaims.ID,
		UserName:      refreshTokenClaims.UserName,
		RoleID:        refreshTokenClaims.RoleID,
	}

	// verify refresh token from redisDB.
	userID, err := mw.redisRepo.GetSession(c.Context(), sessionArgument)
	if err != nil {
		mw.Logger.Error(err)
		return "", "", errlst.NewUnauthorizedError("invalid or expired refresh token")
	}

	var sessionClaims abstract.SessionArgument
	if err := json.Unmarshal(userID, &sessionClaims); err != nil {
		mw.Logger.Error(err)
		return "", "", errlst.NewInternalServerError("failed to parse session claims")
	}

	// generate new tokens.
	newAccessToken, newRefreshToken, err := mw.GenerateTokens(sessionClaims.UserID, sessionClaims.RoleID, sessionClaims.UserName)
	if err != nil {
		mw.Logger.Error(err)
		return "", "", errlst.NewInternalServerError("error creating access, refresh token in HandleNewAccessToken")
	}

	// delete old refresh token from redis and store the new one.
	err = mw.redisRepo.DelSession(c.Context(), abstract.SessionArgument{
		SessionPrefix: "refresh_token",
		RefreshToken:  refreshToken,
		UserID:        sessionClaims.UserID,
	})
	if err != nil {
		mw.Logger.Error(err)
		return "", "", errlst.NewInternalServerError("error deleting old refresh token key from redisDB")
	}
	// Clear refresh token from cookie
	utils.ClearRefreshTokenCookie(c, mw.cfg, refreshToken)

	// store new refresh token in redis.
	err = mw.redisRepo.PutSession(c.Context(), abstract.SessionArgument{
		SessionPrefix: "refresh_token",
		UserID:        sessionClaims.UserID,
		RoleID:        sessionClaims.RoleID,
		RefreshToken:  newRefreshToken,
		ExpiresAt:     time.Duration(24 * time.Hour),
	})
	if err != nil {
		mw.Logger.Error(err)
		return "", "", errlst.NewInternalServerError("error setting new refresh token in redisDB")
	}

	// update cookies with new values
	utils.SetAuthCookies(c, newAccessToken, newRefreshToken)

	// return newAccessToken
	return newAccessToken, newRefreshToken, nil
}
