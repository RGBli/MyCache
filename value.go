package mycache

// Valuer is the interface that all the data type must implement to work with mycache.
type Valuer interface {
	Size() uint64
	Len() int
	Type() string
}
