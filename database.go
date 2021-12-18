package mycache

import (
	"container/list"
	"sync"
	"time"
)

type database struct {
	mu      sync.RWMutex
	mycache *MyCache
	dbName  string
	size    uint64
	cache   map[string]*list.Element
	list    *list.List
}

// get returns the entry for given key.
func (db *database) get(key string) (Valuer, bool) {
	e := db.cache[key]
	if e == nil {
		return nil, false
	}

	// only get alive entry
	if !e.Value.(*entry).expireTime.IsZero() && e.Value.(*entry).expireTime.Before(time.Now()) {
		db.remove(key)
		return nil, false
	}

	db.list.MoveToFront(e)
	return e.Value.(*entry).value, true
}

func (db *database) GetString(key string) (*String, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	v, ok := db.get(key)
	if !ok {
		return nil, false
	}

	s, ok := v.(*String)
	if !ok {
		return nil, false
	}

	return s, true
}

func (db *database) GetList(key string) (*List, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	v, ok := db.get(key)
	if !ok {
		return nil, false
	}

	l, ok := v.(*List)
	if !ok {
		return nil, false
	}

	return l, true
}

func (db *database) GetHash(key string) (*Hash, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	v, ok := db.get(key)
	if !ok {
		return nil, false
	}

	h, ok := v.(*Hash)
	if !ok {
		return nil, false
	}

	return h, true
}

func (db *database) GetSet(key string) (*Set, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	v, ok := db.get(key)
	if !ok {
		return nil, false
	}

	set, ok := v.(*Set)
	if !ok {
		return nil, false
	}

	return set, true
}

func (db *database) GetZset(key string) (*Zset, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	v, ok := db.get(key)
	if !ok {
		return nil, false
	}

	zset, ok := v.(*Zset)
	if !ok {
		return nil, false
	}

	return zset, true
}

// GetExpireTime returns the expire time and whether this entry exists in cache
func (db *database) GetExpireTime(key string) (time.Time, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	e := db.cache[key]
	if e == nil {
		return time.Unix(0, 0), false
	}

	db.list.MoveToFront(e)
	return e.Value.(*entry).expireTime, true
}

// Set stores entry for given key.
func (db *database) Set(key string, value Valuer) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if e, ok := db.cache[key]; ok {
		db.list.MoveToFront(e)
		e.Value.(*entry).value = value
		db.size += value.Size() - e.Value.(*entry).value.Size()
	} else {
		e := &entry{
			key:   key,
			value: value,
		}
		element := db.list.PushFront(e)
		db.cache[key] = element
		db.size += value.Size()
	}

	for db.size > db.mycache.capacity {
		e := db.list.Back()
		delete(db.cache, e.Value.(*entry).key)
		db.list.Remove(e)
		db.size -= e.Value.(*entry).value.Size()
	}
}

// SetExpireTime updates expire time for an entry
func (db *database) SetExpireTime(key string, expireTime time.Time) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if e, ok := db.cache[key]; ok {
		db.list.MoveToFront(e)
		e.Value.(*entry).expireTime = expireTime
	}
}

// SetValueAndExpireTime sets or updates value and expire time for an entry
func (db *database) SetValueAndExpireTime(key string, value Valuer, expireTime time.Time) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if e, ok := db.cache[key]; ok {
		db.list.MoveToFront(e)
		e.Value.(*entry).value = value
		e.Value.(*entry).expireTime = expireTime
		db.size += value.Size() - e.Value.(*entry).value.Size()
	} else {
		e := &entry{
			key:        key,
			value:      value,
			expireTime: expireTime,
		}
		element := db.list.PushFront(e)
		db.cache[key] = element
		db.size += value.Size()
	}

	for db.size > db.mycache.capacity {
		e := db.list.Back()
		delete(db.cache, e.Value.(*entry).key)
		db.list.Remove(e)
		db.size -= e.Value.(*entry).value.Size()
	}
}

// Remove deletes a single entry with lock
func (db *database) Remove(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.remove(key)
}

// remove deletes a single entry
func (db *database) remove(key string) {
	e := db.cache[key]
	if e == nil {
		return
	}
	delete(db.cache, key)
	db.list.Remove(e)
	db.size -= e.Value.(*entry).value.Size()
}

// removeExpired deletes all the expired entries
func (db *database) removeExpired() {
	for key, e := range db.cache {
		if !e.Value.(*entry).expireTime.IsZero() && e.Value.(*entry).expireTime.Before(time.Now()) {
			db.remove(key)
		}
	}
}

// Flush deletes all the entries
func (db *database) Flush(name string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.cache = make(map[string]*list.Element)
	db.list.Init()
	db.size = 0
}

// Exist returns true if key exists in cache
func (db *database) Contains(key string) bool {
	_, ok := db.cache[key]
	return ok
}

func (db *database) getSize() uint64 {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.removeExpired()
	return db.size
}
