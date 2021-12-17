package mycache

import (
	"container/list"
	"github.com/RGBli/MyCache/types"
	"sync"
	"time"
)

// entry is the data stored in list.
type entry struct {
	key        string
	value      Valuer
	expireTime time.Time
}

// MyCache is a thread-safe LRU cache that takes up specific space at most.
type MyCache struct {
	mu       sync.RWMutex
	cache    map[string]*list.Element
	list     *list.List
	capacity uint64
	size     uint64
}

// New returns an initialized MyCache.
func New(capacity uint64) *MyCache {
	return &MyCache{
		cache:    make(map[string]*list.Element),
		list:     list.New(),
		capacity: capacity,
		size:     0,
	}
}

// get returns the entry for given key.
func (c *MyCache) get(key string) (Valuer, bool) {
	e := c.cache[key]
	if e == nil {
		return nil, false
	}

	// only get alive entry
	if !e.Value.(*entry).expireTime.IsZero() && e.Value.(*entry).expireTime.Before(time.Now()) {
		c.remove(key)
		return nil, false
	}

	c.list.MoveToFront(e)
	return e.Value.(*entry).value, true
}

func (c *MyCache) GetString(key string) (*types.String, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.get(key)
	if !ok {
		return nil, false
	}

	s, ok := v.(*types.String)
	if !ok {
		return nil, false
	}

	return s, true
}

func (c *MyCache) GetList(key string) (*types.List, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.get(key)
	if !ok {
		return nil, false
	}

	l, ok := v.(*types.List)
	if !ok {
		return nil, false
	}

	return l, true
}

func (c *MyCache) GetHash(key string) (*types.Hash, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.get(key)
	if !ok {
		return nil, false
	}

	h, ok := v.(*types.Hash)
	if !ok {
		return nil, false
	}

	return h, true
}

func (c *MyCache) GetSet(key string) (*types.Set, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.get(key)
	if !ok {
		return nil, false
	}

	set, ok := v.(*types.Set)
	if !ok {
		return nil, false
	}

	return set, true
}

func (c *MyCache) GetZset(key string) (*types.Zset, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.get(key)
	if !ok {
		return nil, false
	}

	zset, ok := v.(*types.Zset)
	if !ok {
		return nil, false
	}

	return zset, true
}

// GetExpireTime returns the expire time and whether this entry exists in cache
func (c *MyCache) GetExpireTime(key string) (time.Time, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e := c.cache[key]
	if e == nil {
		return time.Unix(0, 0), false
	}

	c.list.MoveToFront(e)
	return e.Value.(*entry).expireTime, true
}

// Set stores entry for given key.
func (c *MyCache) Set(key string, value Valuer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*entry).value = value
		c.size += value.Size() - e.Value.(*entry).value.Size()
	} else {
		e := &entry{
			key:   key,
			value: value,
		}
		element := c.list.PushFront(e)
		c.cache[key] = element
		c.size += value.Size()
	}

	for c.size > c.capacity {
		e := c.list.Back()
		delete(c.cache, e.Value.(*entry).key)
		c.list.Remove(e)
		c.size -= e.Value.(*entry).value.Size()
	}
}

// SetExpireTime updates expire time for an entry
func (c *MyCache) SetExpireTime(key string, expireTime time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*entry).expireTime = expireTime
	}
}

// SetValueAndExpireTime sets or updates value and expire time for an entry
func (c *MyCache) SetValueAndExpireTime(key string, value Valuer, expireTime time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		e.Value.(*entry).value = value
		e.Value.(*entry).expireTime = expireTime
		c.size += value.Size() - e.Value.(*entry).value.Size()
	} else {
		e := &entry{
			key:        key,
			value:      value,
			expireTime: expireTime,
		}
		element := c.list.PushFront(e)
		c.cache[key] = element
		c.size += value.Size()
	}

	for c.size > c.capacity {
		e := c.list.Back()
		delete(c.cache, e.Value.(*entry).key)
		c.list.Remove(e)
		c.size -= e.Value.(*entry).value.Size()
	}
}

// Remove deletes a single entry with lock
func (c *MyCache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.remove(key)
}

// remove deletes a single entry
func (c *MyCache) remove(key string) {
	e := c.cache[key]
	if e == nil {
		return
	}
	delete(c.cache, key)
	c.list.Remove(e)
	c.size -= e.Value.(*entry).value.Size()
}

// removeExpired deletes all the expired entries
func (c *MyCache) removeExpired() {
	for key, e := range c.cache {
		if !e.Value.(*entry).expireTime.IsZero() && e.Value.(*entry).expireTime.Before(time.Now()) {
			c.remove(key)
		}
	}
}

// Flush deletes all the entries
func (c *MyCache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[string]*list.Element)
	c.list.Init()
	c.size = 0
}

// Capacity returns the capacity of the cache
func (c *MyCache) Capacity() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.capacity
}

// Size returns the current space of mycache
func (c *MyCache) Size() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.removeExpired()
	return c.size
}

// Exist returns true if key exists in cache
func (c *MyCache) Contains(key string) bool {
	_, ok := c.cache[key]
	return ok
}
