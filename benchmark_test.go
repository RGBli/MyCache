package mycache

import (
	"strconv"
	"testing"
)

type bytes []byte

func (b bytes) Size() uint64 {
	return uint64(cap(b))
}

func BenchmarkGet(b *testing.B) {
	c := New(64 * 1024 * 1024)
	value := make(bytes, 1024)

	for i := 0; i < 1000; i++ {
		c.Set(strconv.Itoa(i), value)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v, _ := c.Get("500")
		if v == nil {
			panic("error")
		}
		_ = v
	}
}

func BenchmarkLRUCache_Set(b *testing.B) {
	c := New(64 * 1024 * 1024)
	value := make(bytes, 1024)

	for i := 0; i < b.N; i++ {
		c.Set(strconv.Itoa(i), value)
	}
}
