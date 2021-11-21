package mycache

import (
	"testing"
)

type myValue struct {
	age uint64
}

func (v *myValue) Size() uint64 {
	return 1
}

func TestInitialState(t *testing.T) {
	var capacity uint64 = 100
	var size uint64 = 0

	c := New(capacity)
	if c.Capacity() != capacity {
		t.Errorf("capacity = %v, want %v", c.Capacity(), capacity)
	}
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}
}

func TestSimpleCache(t *testing.T) {
	var size uint64 = 1
	key := "key"
	value := &myValue{23}

	c := New(100)
	if v, b := c.Get(key); b {
		t.Errorf("LRUCache has incorrect ele for key(%v): %v", key, v)
	}

	c.Set(key, value)
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}
	if v, _ := c.Get(key); v.(*myValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}

	size = 1
	value = &myValue{3}
	c.Set(key, value)
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}
	if v, _ := c.Get(key); v.(*myValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}
}

func TestCapacity(t *testing.T) {
	var capacity uint64 = 3
	c := New(capacity)
	value := &myValue{1}

	c.Set("key1", value)
	c.Set("key2", value)
	c.Set("key3", value)
	if c.Size() != capacity {
		t.Errorf("size = %v, want %v", c.Size(), capacity)
	}

	c.Set("key4", value)
	if c.Size() != capacity {
		t.Errorf("size = %v, want %v", c.Size(), capacity)
	}
	if _, b := c.Get("key1"); b {
		t.Errorf("key1 is not evicted")
	}
}

func TestDelete(t *testing.T) {
	var size uint64 = 1
	key := "key"
	value := &myValue{size}

	c := New(100)
	c.Set(key, value)
	if v, _ := c.Get(key); v.(*myValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}

	c.Delete("key2")
	if v, _ := c.Get(key); v.(*myValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}
	if c.Size() != size {
		t.Errorf("size = %v, want %v", c.Size(), size)
	}

	size = 0
	c.Delete(key)
	if _, b := c.Get(key); b {
		t.Errorf("failed to delete %v", key)
	}
	if c.Size() != 0 {
		t.Errorf("size = %v, want %v", c.Size(), 0)
	}
}

func TestFlush(t *testing.T) {
	key := "key"
	value := &myValue{20}

	c := New(100)
	c.Set(key, value)
	if v, _ := c.Get(key); v.(*myValue) != value {
		t.Errorf("%v = %v, want %v", key, v, value)
	}

	c.Flush()
	if _, b := c.Get(key); b {
		t.Errorf("failed to delete %v", key)
	}
	if c.Size() != 0 {
		t.Errorf("size = %v, want %v", c.Size(), 0)
	}
}
