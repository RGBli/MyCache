package mycache

import (
	"testing"
	"time"
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
		t.Errorf("capacity = %v, expect %v", c.Capacity(), capacity)
	}
	if c.Size() != size {
		t.Errorf("size = %v, expect %v", c.Size(), size)
	}
}

func TestGetSet(t *testing.T) {
	key := "lbw"
	value := &myValue{23}
	c := New(100)
	c.Set(key, value)
	v, _ := c.Get(key)
	if v.(*myValue).age != 23 {
		t.Errorf("get %d, expect 23", v.(*myValue).age)
	}
}

func TestExpireTime(t *testing.T) {
	key := "lbw"
	value := &myValue{23}
	c := New(100)

	start := time.Now()
	c.SetValueAndExpireTime(key, value, start.Add(time.Second*10))
	expireTime, ok := c.GetExpireTime(key)
	if expireTime.Before(start.Add(time.Second*9)) || expireTime.After(start.Add(time.Second*11)) {
		t.Errorf("get %v, expect %v", expireTime, start.Add(time.Second*55))
	}

	if !ok {
		t.Errorf("get false, expect true")
	}

	time.Sleep(time.Second * 10)
	if v, _ := c.Get(key); v != nil {
		t.Errorf("%v = %v, expect nil", key, v)
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
		t.Errorf("size = %v, expect %v", c.Size(), capacity)
	}

	c.Set("key4", value)
	if c.Size() != capacity {
		t.Errorf("size = %v, expect %v", c.Size(), capacity)
	}
	if _, b := c.Get("key1"); b {
		t.Errorf("key1 is not evicted")
	}
}
