package constants

import "time"

const (
	// when user sign-up to application, it's default roleID is 3.
	DefaultRoleID = 3

	// token expiration time.
	TokenExpiryTime = 24 // 1day.

	// context timeout.
	CtxDefaultTimeOut = 10 * time.Second
)
