package mycache

import (
	"github.com/RGBli/MyCache/skiplist"
)

type Zset struct {
	list *skiplist.SkipList
}

func NewZset() *Zset {
	return &Zset{list: skiplist.New()}
}

func (z *Zset) Size() uint64 {
	return z.list.Size()
}

func (z *Zset) Len() int {
	return z.list.Len()
}

func (z *Zset) Type() string {
	return "Zset"
}

func (z *Zset) Add(score float64, value string) {
	z.list.Set(score, value)
}

func (z *Zset) Get(score float64) (string, bool) {
	value := z.list.Get(score)
	if value != nil {
		return value.Value(), true
	}
	return "", false
}

func (z *Zset) GetRange(start, end, step float64) []string {
	strs := make([]string, 0)
	for i := start; i < end; i += step {
		value := z.list.Get(i)
		if value != nil {
			strs = append(strs, value.Value())
		}
	}
	return strs
}

func (z *Zset) Remove(score float64) {
	z.list.Remove(score)
}
