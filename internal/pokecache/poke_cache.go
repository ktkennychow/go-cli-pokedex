package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mux sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
func NewCache(interval time.Duration) *Cache{
	newCache := &Cache{cacheEntries: make(map[string]cacheEntry)}
	go newCache.reaploop(interval)
	return newCache
}
func (c *Cache) Add(key string, val []byte){
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: 	time.Now(), 
		val: 				val,
	}
}
func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.RLock()
	defer c.mux.RUnlock()
	cacheEntry, exists := c.cacheEntries[key]
	if (!exists) {
		return nil, false
	}
	return cacheEntry.val, true
}
func (c *Cache) reaploop(interval time.Duration){
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for range ticker.C{
		c.reap(interval)
	}
}

func (c *Cache)reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	
	now := time.Now()
	for key, entry := range c.cacheEntries {
		if now.Sub(entry.createdAt) > interval {
			go delete(c.cacheEntries, key)
		}
	}
}