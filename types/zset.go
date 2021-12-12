package types

import "github.com/RGBli/MyCache/types/skiplist"

type Zset struct {
	skiplist *skiplist.SkipList
}

func NewZset() *Zset {
	return &Zset{skiplist: skiplist.New()}
}

func (z *Zset) Size() uint64 {
	var size uint64 = 0
	node := z.skiplist.Front()
	for node != nil {
		size += uint64(len(node.Value().(string)) + node.NextNodesNum() * 8 + 8)
		node = node.Next()
	}
	return size
}

func (z *Zset) Add(score float64, value string) {
	z.skiplist.Set(score, value)
}

func (z *Zset) Get(score float64) (string, bool) {
	value := z.skiplist.Get(score).Value()
	if value != nil {
		return value.(string), true
	}
	return "", false
}

func (z *Zset) GetRange(start, end, step float64) []string {
	strs := make([]string, 0)
	for i := start; i < end; i+=step {
		value := z.skiplist.Get(i).Value()
		if value != nil {
			strs = append(strs, value.(string))
		}
	}
	return strs
}

func (z *Zset) Remove(score float64) {
	z.skiplist.Remove(score)
}

func (z *Zset) Len() int {
	return z.skiplist.Length
}