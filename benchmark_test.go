package mycache

import (
	"strconv"
	"testing"
)

type bytes []byte

func (b bytes) Size() uint64 {
	return uint64(cap(b))
}

func BenchmarkLRUCacheGet(b *testing.B) {
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

func BenchmarkLRUCacheSet(b *testing.B) {
	c := New(64 * 1024 * 1024)
	value := make(bytes, 1024)

	for i := 0; i < b.N; i++ {
		c.Set(strconv.Itoa(i), value)
	}
}

func BenchmarkMapGet(b *testing.B) {
	m := make(map[string]bytes)
	value := make(bytes, 1000)
	for i := 0; i < 1000; i++ {
		m[strconv.Itoa(i)] = value
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v, _ := m["500"]
		if v == nil {
			panic("error")
		}
		_ = v
	}
}

func BenchmarkMapSet(b *testing.B) {
	m := make(map[string]bytes)
	value := make(bytes, 1000)

	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = value
	}
}
