package skiplist

import (
	"testing"
)

func TestNewLen(t *testing.T) {
	list := New()
	len := list.Len()
	if len != 0 {
		t.Errorf("got %d, expect 0", len)
	}
}

func TestSet(t *testing.T) {
	list := New()
	list.Set(1.0, "lbw")
	len := list.Len()
	if len != 1 {
		t.Errorf("got %d, expect 1", len)
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Set(1.0, "lbw")
	v := list.Get(1.0).Value()
	if v != "lbw" {
		t.Errorf("got %s, expect lbw", v)
	}
}

func TestRemoveContains(t *testing.T) {
	list := New()
	list.Set(1.0, "lbw")
	exist := list.Contains(1.0)
	if !exist {
		t.Errorf("got %t, expect true", exist)
	}

	list.Remove(1.0)
	exist = list.Contains(1.0)
	if exist {
		t.Errorf("got %t, expect false", exist)
	}
}
