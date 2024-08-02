package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache{
	c := Cache{cache: make(map[string]cacheEntry), mux: &sync.RWMutex{},}

	go c.reaploop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte){
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: 	time.Now().UTC(), 
		val: 				val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.RLock()
	defer c.mux.RUnlock()
	cache, exist := c.cache[key]
	return cache.val, exist
}

func (c *Cache) reaploop(interval time.Duration){
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for range ticker.C{
		c.reap(time.Now().UTC(),interval)
	}
}

func (c *Cache)reap(now time.Time, interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	
	for key, entry := range c.cache {
		if now.Sub(entry.createdAt) > interval {
			go delete(c.cache, key)
		}
	}
}