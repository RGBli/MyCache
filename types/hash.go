package types

type Hash map[string]string

func NewHash() Hash {
	return make(map[string]string)
}

func (h Hash) Size() uint64 {
	var size uint64 = 0
	for k, v := range h {
		size += uint64(len(k) + len(v))
	}
	return size
}

func (h Hash) Put(key string, value string) {
	h[key] = value
}

func (h Hash) Get(key string) (value string, ok bool) {
	value, ok = h[key]
	return
}

func (h Hash) Remove(key string) {
	delete(h, key)
}

func (h Hash) Len() int {
	return len(h)
}