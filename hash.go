package mycache

type Hash struct {
	h map[string]string
}

func NewHash() *Hash {
	return &Hash{h: make(map[string]string)}
}

func (h *Hash) Size() uint64 {
	var size uint64 = 0
	for k, v := range h.h {
		size += uint64(len(k) + len(v))
	}
	return size
}

func (h *Hash) Len() int {
	return len(h.h)
}

func (h *Hash) Type() string {
	return "Hash"
}

func (h *Hash) Put(key string, value string) {
	h.h[key] = value
}

func (h *Hash) Get(key string) (value string, ok bool) {
	value, ok = h.h[key]
	return
}

func (h *Hash) Remove(key string) {
	delete(h.h, key)
}
