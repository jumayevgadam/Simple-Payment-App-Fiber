package abstract

import (
	"strconv"
	"time"

	redisModel "github.com/jumayevgadaym/tsu-toleg/internal/models/redis"
)

// CacheArgument for creating key:value pair in redisDB.
type CacheArgument struct {
	ObjectID   int
	ObjectType string
}

type SessionArgument struct {
	SessionPrefix string
	UserID        int
	RoleID        int
	UserName      string
	RefreshToken  string
	ExpiresAt     time.Duration
}

// ToCacheStorage for Sending CacheArgument into memory.
func (c *CacheArgument) ToCacheStorage() redisModel.CacheArgument {
	return redisModel.CacheArgument{
		ID:         strconv.Itoa(int(c.ObjectID)),
		ObjectType: c.ObjectType,
	}
}

// ToSessionStorage for sending SessionArgument into memory.
func (s *SessionArgument) ToSessionStorage() redisModel.SessionArgument {
	return redisModel.SessionArgument{
		UserID:        strconv.Itoa(int(s.UserID)),
		RoleID:        strconv.Itoa(int(s.RoleID)),
		UserName:      s.UserName,
		SessionPrefix: s.SessionPrefix,
		RefreshToken:  s.RefreshToken,
		ExpiresAt:     s.ExpiresAt,
	}
}
