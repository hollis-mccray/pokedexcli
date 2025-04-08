package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) Cache {
	data := Cache{}
	data.cache = make(map[string]cacheEntry)
	go data.reapLoop(interval)
	return data
}

func (c Cache) Add(key string, val []byte) {
    c.mux.Lock()
    defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return
}

func (c Cache) Get(key string) ([]byte, bool) {
    c.mux.RLock()
    defer c.mux.RUnlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mux.Lock()
    defer c.mux.Unlock()
    for k, v := range c.cache {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.cache, k)
        }
    }
}