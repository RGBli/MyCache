package mycache

import (
	"container/list"
	"sync"
)

// Valuer is the interface that all the data types must implement to work with mycache.
type Valuer interface {
	Size() uint64
}

// ele is the data stored in list.
type ele struct {
	key   string
	value Valuer
}

// LRUCache is a thread-safe LRU cache that takes up specific space at most.
type LRUCache struct {
	mu       sync.RWMutex
	cache    map[string]*list.Element
	list     *list.List
	capacity uint64
	size     uint64
}

// New returns an initialized LRUCache.
func New(capacity uint64) *LRUCache {
	return &LRUCache{
		cache:    make(map[string]*list.Element),
		list:     list.New(),
		capacity: capacity,
		size: 0,
	}
}

// Get returns the ele for given key.
func (lru *LRUCache) Get(key string) (Valuer, bool) {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	e := lru.cache[key]
	if e == nil {
		return nil, false
	}

	lru.list.MoveToFront(e)
	return e.Value.(*ele).value, true
}

// Set stores ele for given key.
func (lru *LRUCache) Set(key string, value Valuer) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if e := lru.cache[key]; e != nil {
		lru.list.MoveToFront(e)
		e.Value.(*ele).value = value
		lru.size += value.Size() - e.Value.(*ele).value.Size()
	} else {
		e := &ele{
			key:   key,
			value: value,
		}
		element := lru.list.PushFront(e)
		lru.cache[key] = element
		lru.size += value.Size()
	}

	for lru.size > lru.capacity {
		e := lru.list.Back()
		delete(lru.cache, e.Value.(*ele).key)
		lru.list.Remove(e)
		lru.size -= e.Value.(*ele).value.Size()
	}
}

// Delete removes given key and its value.
func (lru *LRUCache) Delete(key string) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	e := lru.cache[key]
	if e == nil {
		return
	}
	delete(lru.cache, key)
	lru.list.Remove(e)
	lru.size -= e.Value.(*ele).value.Size()
}

// Flush removes all the K-V pairs.
func (lru *LRUCache) Flush() {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	lru.cache = make(map[string]*list.Element)
	lru.list.Init()
	lru.size = 0
}

// Capacity returns the capacity of the cache.
func (lru *LRUCache) Capacity() uint64 {
	lru.mu.RLock()
	defer lru.mu.RUnlock()
	return lru.capacity
}

// Size returns the current space of mycache.
func (lru *LRUCache) Size() uint64 {
	lru.mu.RLock()
	defer lru.mu.RUnlock()
	return lru.size
}
