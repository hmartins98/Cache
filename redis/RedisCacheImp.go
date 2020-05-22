package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//Get TODO
func (c *Cache) Get(key string, obj interface{}) error {
	val, err := c.cluster.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &obj)
	if err != nil {
		return err
	}
	return nil
}

//Set TODO
func (c *Cache) Set(key string, obj interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = c.cluster.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//Delete TODO
func (c *Cache) Delete(key string) {
	err := c.cluster.Del(key)

	if err != nil {
		fmt.Printf("Error deleting key: %v", key)
	}
}
