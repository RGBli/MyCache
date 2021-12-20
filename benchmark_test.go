package mycache

import (
	"strconv"
	"testing"
)

func BenchmarkLRUCacheGet(b *testing.B) {
	dbName := "test"
	c := Default()
	db := c.Use(dbName)
	value := NewString("23")
	for i := 0; i < 1000; i++ {
		db.SetValue(strconv.Itoa(i), value)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v, _ := db.GetString("500")
		if v == nil {
			panic("error")
		}
		_ = v
	}
}

func BenchmarkLRUCacheSet(b *testing.B) {
	dbName := "test"
	c := Default()
	db := c.Use(dbName)
	value := NewString("23")

	for i := 0; i < b.N; i++ {
		db.SetValue(strconv.Itoa(i), value)
	}
}

func BenchmarkMapGet(b *testing.B) {
	m := make(map[string]string)
	value := "23"
	for i := 0; i < 1000; i++ {
		m[strconv.Itoa(i)] = value
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v, ok := m["500"]
		if !ok {
			panic("error")
		}
		_ = v
	}
}

func BenchmarkMapSet(b *testing.B) {
	m := make(map[string]string)
	value := "23"

	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = value
	}
}
