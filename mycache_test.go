package mycache

import (
	"fmt"
	"testing"
	"time"

	"github.com/RGBli/MyCache/types"
)

func TestInitialState(t *testing.T) {
	var capacity uint64 = 1 * 1024 * 1024
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
	value := types.NewString("23")
	c := New(1 * 1024 * 1024)
	c.Set(key, value)
	v, ok := c.GetString(key)
	if !ok {
		t.Errorf("ok should be true")
	}
	if v.ToString() != "23" {
		t.Errorf("get %s, expect 23", v.ToString())
	}
}

func TestExpireTime(t *testing.T) {
	key := "lbw"
	value := types.NewString("23")
	c := New(1 * 1024 * 1024)

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
	if v, _ := c.GetString(key); v != nil {
		t.Errorf("%v = %v, expect nil", key, v)
	}
}

func TestList(t *testing.T) {
	c := New(1 * 1024 * 1024)
	key := "lbw"
	list := types.NewList([]string{"foo", "bar"})
	list.Add("test")
	if list.Len() != 3 {
		t.Errorf("list length is %d, expect 3", list.Len())
	}

	_ = list.Remove(2)
	if list.Len() != 2 {
		t.Errorf("list length is %d, expect 2", list.Len())
	}

	c.Set(key, list)
	v, _ := c.GetList(key)
	s, _ := v.Get(1)
	if s != "bar" {
		t.Errorf("got %s, expect bar", s)
	}
}

func TestHash(t *testing.T) {
	c := New(1 * 1024 * 1024)
	key := "lbw"
	hash := types.NewHash()
	hash.Put("age", "23")
	hash.Put("gender", "male")
	c.Set(key, hash)
	v, _ := c.GetHash(key)
	s, _ := v.Get("age")
	if s != "23" {
		t.Errorf("got %s, expect 23", s)
	}
}

func TestSet(t *testing.T) {
	c := New(1 * 1024 * 1024)
	key := "lbw"
	set := types.NewSet([]string{"foo", "foo", "bar"})
	c.Set(key, set)
	v, _ := c.GetSet(key)
	len := v.Len()
	if len != 2 {
		t.Errorf("got %d, expect 2", len)
	}

	strs := v.GetAll()
	for _, str := range strs {
		fmt.Println(str)
	}
	if strs[1] != "bar" {
		t.Errorf("got %s, expect bar", strs[1])
	}
}

func TestZset(t *testing.T) {
	c := New(1 * 1024 * 1024)
	key := "lbw"
	zset := types.NewZset()
	zset.Add(1.0, "23")
	zset.Add(0.0, "25")
	c.Set(key, zset)
	if zset.Len() != 2 {
		t.Errorf("got %d, expect 2", zset.Len())
	}
	v, _ := c.GetZset(key)
	strs := v.GetRange(0, 2, 1)
	if strs[1] != "23" {
		t.Errorf("got %s, expect 23", strs[1])
	}
}
