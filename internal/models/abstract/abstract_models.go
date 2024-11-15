package abstract

import (
	"strconv"

	redisModel "github.com/jumayevgadaym/tsu-toleg/internal/models/redis"
)

// CacheArgument for creating key:value pair in redisDB.
type CacheArgument struct {
	ObjectID   int
	ObjectType string
}

// ToCacheStorage for Sending CacheArgument into memory.
func (c *CacheArgument) ToCacheStorage() redisModel.CacheArgument {
	return redisModel.CacheArgument{
		ID:         strconv.Itoa(int(c.ObjectID)),
		ObjectType: c.ObjectType,
	}
}
