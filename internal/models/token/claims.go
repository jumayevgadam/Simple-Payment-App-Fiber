package token

import "github.com/golang-jwt/jwt/v5"

// AccessTokenClaims model for access token.
type AccessTokenClaims struct {
	ID          int      `json:"user_id"`
	RoleID      int      `json:"role_id"`
	UserName    string   `json:"username"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}
