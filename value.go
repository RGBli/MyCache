package mycache

// Valuer is the interface that all the data types must implement to work with mycache.
type Valuer interface {
	Size() uint64
	Len() int
}
