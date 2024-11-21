package redis

import "time"

// Argument model is.
type Argument struct {
	CacheArgument
	SessionArgument
}

// CacheKey model need for redisDB.
type CacheArgument struct {
	ID         string
	ObjectType string
}

// SessionArgument model is.
type SessionArgument struct {
	UserID        string
	RoleID        string
	UserName      string
	SessionPrefix string
	RefreshToken  string
	ExpiresAt     time.Duration
}
