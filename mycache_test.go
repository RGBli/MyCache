package mycache

import (
	"testing"
	"time"
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
	value := NewString("23")
	dbName := "test"
	c := New(1 * 1024 * 1024)
	db := c.Use(dbName)
	db.Set(key, value)
	v, ok := db.GetString(key)
	if !ok {
		t.Errorf("ok should be true")
	}
	if v.ToString() != "23" {
		t.Errorf("get %s, expect 23", v.ToString())
	}
}

func TestExpireTime(t *testing.T) {
	key := "lbw"
	value := NewString("23")
	dbName := "test"
	c := New(1 * 1024 * 1024)
	db := c.Use(dbName)

	start := time.Now()
	db.SetValueAndExpireTime(key, value, start.Add(time.Second*10))
	expireTime, ok := db.GetExpireTime(key)
	if expireTime.Before(start.Add(time.Second*9)) || expireTime.After(start.Add(time.Second*11)) {
		t.Errorf("get %v, expect %v", expireTime, start.Add(time.Second*55))
	}

	if !ok {
		t.Errorf("get false, expect true")
	}

	time.Sleep(time.Second * 10)
	if v, _ := db.GetString(key); v != nil {
		t.Errorf("%v = %v, expect nil", key, v)
	}
}

func TestDatabase(t *testing.T) {
	c := New(1 * 1024 * 1024)
	dbName1 := "test1"
	db1 := c.Use(dbName1)
	db1.Set("lbw", NewString("23"))
	v1, _ := db1.GetString("lbw")
	if v1.ToString() != "23" {
		t.Errorf("got %s, expect 23", v1.ToString())
	}

	dbName2 := "test2"
	db2 := c.Use(dbName2)
	db2.Set("lbw", NewString("3"))
	v2, _ := db2.GetString("lbw")
	if v2.ToString() != "3" {
		t.Errorf("got %s, expect 3", v1.ToString())
	}
}

func TestList(t *testing.T) {
	dbName := "test"
	c := New(1 * 1024 * 1024)
	db := c.Use(dbName)
	key := "lbw"
	list := NewList([]string{"foo", "bar"})
	list.Add("test")
	if list.Len() != 3 {
		t.Errorf("list length is %d, expect 3", list.Len())
	}

	_ = list.Remove(2)
	if list.Len() != 2 {
		t.Errorf("list length is %d, expect 2", list.Len())
	}

	db.Set(key, list)
	v, _ := db.GetList(key)
	s, _ := v.Get(1)
	if s != "bar" {
		t.Errorf("got %s, expect bar", s)
	}
}

func TestHash(t *testing.T) {
	dbName := "test"
	c := New(1 * 1024 * 1024)
	db := c.Use(dbName)
	key := "lbw"
	hash := NewHash()
	hash.Put("age", "23")
	hash.Put("gender", "male")
	db.Set(key, hash)
	v, _ := db.GetHash(key)
	s, _ := v.Get("age")
	if s != "23" {
		t.Errorf("got %s, expect 23", s)
	}
}

func TestSet(t *testing.T) {
	dbName := "test"
	c := New(1 * 1024 * 1024)
	db := c.Use(dbName)
	key := "lbw"
	set := NewSet([]string{"foo", "foo", "bar"})
	db.Set(key, set)
	v, _ := db.GetSet(key)
	len := v.Len()
	if len != 2 {
		t.Errorf("got %d, expect 2", len)
	}

	strs := v.GetAll()
	if strs[1] != "bar" {
		t.Errorf("got %s, expect bar", strs[1])
	}
}

func TestZset(t *testing.T) {
	dbName := "test"
	c := New(1 * 1024 * 1024)
	db := c.Use(dbName)
	key := "lbw"
	zset := NewZset()
	zset.Add(1.0, "23")
	zset.Add(0.0, "25")
	db.Set(key, zset)
	if zset.Len() != 2 {
		t.Errorf("got %d, expect 2", zset.Len())
	}
	v, _ := db.GetZset(key)
	strs := v.GetRange(0, 2, 1)
	if strs[1] != "23" {
		t.Errorf("got %s, expect 23", strs[1])
	}
}
