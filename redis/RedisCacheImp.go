package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

//Get TODO
func (c *Cache) Get(key string, obj interface{}) error {
	val, err := c.Cluster.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &obj)
	log.Printf("Get: %v", obj)
	if err != nil {
		return err
	}
	return nil
}

//Set TODO
func (c *Cache) Set(key string, obj interface{}, expiration time.Duration) error {
	log.Printf("Set: %v", obj)
	cacheEntry, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = c.Cluster.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//Delete TODO
func (c *Cache) Delete(key string) {
	err := c.Cluster.Del(key)

	if err != nil {
		fmt.Printf("Error deleting key: %v", key)
	}
}
