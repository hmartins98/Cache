package redis

import (
	"fmt"

	"github.com/go-redis/cache"
)

//Get TODO
func (c *Cache) Get(key string, obj interface{}) {
	err := c.cache.Get(key, obj)

	if err != nil {
		fmt.Printf("Error reading key: %v", key)
	}
}

//Set TODO
func (c *Cache) Set(key string, obj interface{}) {
	i := &cache.Item{
		Key:    key,
		Object: obj,
	}
	err := c.cache.Set(i)

	if err != nil {
		fmt.Printf("Error setting key: %v", key)
	}
}

//Delete TODO
func (c *Cache) Delete(key string) {
	err := c.cache.Delete(key)

	if err != nil {
		fmt.Printf("Error deleting key: %v", key)
	}
}
