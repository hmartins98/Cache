package redis

import (
	"github.com/go-redis/redis"
)

//Cache struct
type Cache struct {
	Cluster *redis.ClusterClient
	Addrs   []string
}

//Initialize init
func (c *Cache) Initialize() {
	c.Cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: c.Addrs,
	})
	// if err := c.Cluster.Ping().Err(); err != nil {
	// 	panic("Unable to connect to redis " + err.Error())
	// }
}

//Close TODO
func (c *Cache) Close() {
	c.Cluster.Close()
}
