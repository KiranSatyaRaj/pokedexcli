package pokecache

import (
	"sync"
	"time"
	"github.com/KiranSatyaRaj/pokedexcli/pkg/results"
)

type Cache struct {
	mu sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	value       []results.LocationArea
}

var end   time.Duration

var cache Cache

func NewCache(interval time.Duration) Cache {
	end = time.Duration(interval)
	cache = Cache{}
	cache.cache = make(map[string]cacheEntry)
	cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []results.LocationArea) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	if !ok {
		c.cache[key] = cacheEntry{createdAt: time.Now(), value: val}
	}
}

func (c *Cache) Get(key string) ([]results.LocationArea, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	if !ok {
		return []results.LocationArea{}, ok
	}
	return c.cache[key].value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
 	ticker := time.NewTicker(interval)
	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		for t := range ticker.C {
			for key, entry := range c.cache {
				if t.Sub(entry.createdAt) >= interval {
					delete(c.cache, key)
				}
			}
		}
	}()
}

func (c *Cache) IsEmpty() bool {
	return c.cache == nil
}
