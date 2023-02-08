package cache_ttl

import (
	"errors"
	"sync"
	"time"
)

// ErrKeyNoFound - used when key no found.
var ErrKeyNoFound = errors.New("key not found")

// Cache - used as cache entity.
type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}

// Item - used as concrete item in cache.
type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

// NewCache - used for create cache instance.
func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {

	// инициализируем карту(map) в паре ключ(string)/значение(Item)
	items := make(map[string]Item)

	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	// Если интервал очистки больше 0, запускаем GC (удаление устаревших элементов)
	if cleanupInterval > 0 {
		cache.StartGC()
	}

	return &cache
}

// Set - used for setup value.
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64

	// Если продолжительность жизни равна 0 - используется значение по-умолчанию
	if duration == 0 {
		duration = c.defaultExpiration
	}

	// Устанавливаем время истечения кеша
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}

// Get - used to get the value from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}

	return item.Value, true
}

// Delete - used for delete value from cache.
func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}

// StartGC - used to start the cleaning process.
func (c *Cache) StartGC() {
	go c.GC()
}

// GC - used as cleaning process.
func (c *Cache) GC() {
	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

// expiredKeys - used for return expired keys.
func (c *Cache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

// clearItems - used for delete keys form this list.
func (c *Cache) clearItems(keys []string) {
	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
