package abstract

import (
	"strconv"

	redisModel "github.com/jumayevgadaym/tsu-toleg/internal/models/redis"
)

// CacheArgument is
type CacheArgument struct {
	ObjectID   int
	ObjectType string
}

// ToCacheStorage is
func (c *CacheArgument) ToCacheStorage() redisModel.CacheArgument {
	return redisModel.CacheArgument{
		ID:         strconv.Itoa(int(c.ObjectID)),
		ObjectType: c.ObjectType,
	}
}
