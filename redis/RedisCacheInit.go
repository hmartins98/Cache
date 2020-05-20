package redis

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

//Cache TODO
type Cache struct {
	cache *cache.Codec
	ring  *redis.Ring
}

//OpenCacheConnection TODO
func (c *Cache) OpenCacheConnection() {
	c.ring = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"redis": "6379",
			// "server2": "6380",
		},
	})

	c.cache = &cache.Codec{
		Redis: c.ring,

		Marshal:   msgpack.Marshal,
		Unmarshal: msgpack.Unmarshal,
	}
}

//CloseCacheConnection TODO
func (c *Cache) CloseCacheConnection() {
	c.ring.Close()
}
