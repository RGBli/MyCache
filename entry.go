package mycache

import "time"

// Valuer is the interface that all the data skiplist must implement to work with mycache.
type Valuer interface {
	Size() uint64
	Len() int
}

// entry is the data stored in list.
type entry struct {
	key        string
	value      Valuer
	expireTime time.Time
}
