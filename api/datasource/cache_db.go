package datasource

import "time"

type CacheDB struct {
	DB ReaderWriter
}

func (c *CacheDB) Get(key string) string {
	return ""
}

func (c *CacheDB) Set(key string, value string, ttl time.Time) bool {
	return true
}