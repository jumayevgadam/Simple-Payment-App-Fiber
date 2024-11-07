package token

import "github.com/golang-jwt/jwt/v5"

// AccessTokenClaims model is
type AccessTokenClaims struct {
	ID       int    `json:"user_id"`
	RoleID   int    `json:"role_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// RefreshTokenClaims model is
type RefreshTokenClaims struct {
	ID     int `json:"user_id"`
	RoleID int `json:"role_id"`
	jwt.RegisteredClaims
}
