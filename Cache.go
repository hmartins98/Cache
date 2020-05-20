package cache

type ICache interface {
	OpenCacheConnection()
	CloseCacheConnection()
	Get(key string, obj interface{})
	Set(key string, obj interface{})
	Delete(key string)
}
