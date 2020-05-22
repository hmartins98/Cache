package redis

import (
	"github.com/go-redis/redis"
)

//Cache struct
type Cache struct {
	cluster *redis.ClusterClient
}

//Initialize init
func (c *Cache) Initialize(addrs []string) {
	c.cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})
	if err := c.cluster.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
}

//Close TODO
func (c *Cache) Close() {
	c.cluster.Close()
}
