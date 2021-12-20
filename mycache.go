package mycache

import (
	"container/list"
	"sync"
	"time"
)

const (
	DefaultPersistPath   = "log"
	DefaultCleanInterval = time.Minute
	DefaultCapacity      = 10 * 1024 * 1024
)

type MyCache struct {
	mu            sync.RWMutex
	databases     map[string]*database
	capacity      uint64
	size          uint64
	cleanInterval time.Duration
	persistPath   string
}

// New returns an initialized MyCache
func New(capacity uint64, cleanInterval time.Duration, persistPath string) *MyCache {
	return &MyCache{
		databases:     make(map[string]*database),
		capacity:      capacity,
		size:          0,
		cleanInterval: cleanInterval,
		persistPath:   persistPath,
	}
}

// Default returns a mycache instance initialized with default parameters
func Default() *MyCache {
	return &MyCache{
		databases:     make(map[string]*database),
		capacity:      DefaultCapacity,
		size:          0,
		cleanInterval: DefaultCleanInterval,
		persistPath:   DefaultPersistPath,
	}
}

// Use select or create a database
func (c *MyCache) Use(name string) *database {
	c.mu.Lock()
	defer c.mu.Unlock()

	db, ok := c.databases[name]
	if !ok {
		db = &database{
			mycache: c,
			dbName:  name,
			cache:   make(map[string]*list.Element),
			list:    list.New(),
		}
		c.databases[name] = db
	}

	if c.cleanInterval > 0 {
		go func() {
			for {
				time.Sleep(c.cleanInterval)
				db.RemoveExpired()
			}
		}()
	}
	return db
}

// Capacity returns the capacity of the cache
func (c *MyCache) Capacity() uint64 {
	return c.capacity
}

// Size returns the current space of mycache
func (c *MyCache) Size() uint64 {
	var size uint64 = 0
	for _, db := range c.databases {
		size += db.getSize()
	}
	return size
}

func (c *MyCache) SetCapacity(capacity uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.capacity = capacity
}

func (c *MyCache) SetPersistPath(persistPath string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.persistPath = persistPath
}
