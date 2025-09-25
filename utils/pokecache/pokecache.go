package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu *sync.Mutex
}

func (c *Cache) Add(key string, value []byte) *Cache {
	c.mu.Lock()
	var cEntry = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
	c.cacheMap[key] = cEntry
	c.mu.Unlock()
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cache, exists := c.cacheMap[key]
	if !exists {
		c.mu.Unlock()
		return nil, false
	}
	c.mu.Unlock()
	return cache.val, true
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheMap {
			age := time.Since(entry.createdAt)
			if age > interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	newCacheEntry := make(map[string]cacheEntry)
	mutex := &sync.Mutex{}
	var newCache = &Cache{
		cacheMap: newCacheEntry,
		mu: mutex,
	}
	go newCache.reaploop(interval)
 	return newCache
}