package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	m map[string]cacheEntry
	mu sync.RWMutex
	interval time.Duration
}

//constructor
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{m: make(map[string]cacheEntry), interval: interval}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.m[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	if entry, ok := c.m[key]; ok {
		if time.Now().Sub(entry.createdAt) > c.interval {
			c.mu.RUnlock()
			return nil, false
		}
		c.mu.RUnlock()
		return entry.val, true
	}
	c.mu.RUnlock()
	return nil, false
}

func (c *Cache) reapLoop() {
	for {
		time.Sleep(c.interval)
		c.mu.Lock()
		if len(c.m) == 0 {
			c.mu.Unlock()
			continue
		}
		for k := range c.m {
			entry:= c.m[k]
			if time.Now().Sub(entry.createdAt) > c.interval {
				delete(c.m, k)
			}
		}
		c.mu.Unlock()
	}
}