package redis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
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
	if err != nil {
		return err
	}
	return nil
}

//Set TODO
func (c *Cache) Set(key string, obj interface{}, expiration time.Duration) error {
	b := new(bytes.Buffer)

	if err := gob.NewEncoder(b).Encode(obj); err != nil {
		return err
	}

	if err := c.Cluster.Set(key, b.Bytes(), expiration).Err(); err != nil {
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
