package mycache

import (
	"container/list"
)

type MyCache struct {
	databases map[string]*database
	capacity  uint64
	size      uint64
}

// New returns an initialized MyCache.
func New(capacity uint64) *MyCache {
	return &MyCache{
		databases: make(map[string]*database),
		capacity:  capacity,
		size:      0,
	}
}

func (c *MyCache) Use(name string) *database {
	if db, ok := c.databases[name]; ok {
		return db
	}

	return &database{
		mycache: c,
		dbName:  name,
		cache:   make(map[string]*list.Element),
		list:    list.New(),
	}
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
