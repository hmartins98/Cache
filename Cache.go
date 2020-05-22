package cache

import "time"

//ICache cache
type ICache interface {
	Initialize(addrs []string)
	Close()
	Get(key string, obj interface{}) error
	Set(key string, obj interface{}, expiration time.Duration) error
	Delete(key string)
}
