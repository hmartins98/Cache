package redis

import (
	"github.com/go-redis/redis"
)

//Cache struct
type Cache struct {
	cluster *redis.ClusterClient
}

//GetClient get the redis client
func (c *Cache) initialize(addrs []string) {
	c.cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})
	if err := c.cluster.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
}

//CloseCacheConnection TODO
func (c *Cache) CloseCacheConnection() {
	c.cluster.Close()
}
