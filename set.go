package mycache

type Set struct {
	s map[string]struct{}
}

func NewEmptySet() *Set {
	return &Set{s: make(map[string]struct{})}
}

func NewSet(strs []string) *Set {
	set := &Set{s: make(map[string]struct{})}
	for _, str := range strs {
		set.Add(str)
	}
	return set
}

func (set *Set) Size() uint64 {
	var size uint64 = 0
	for k := range set.s {
		size += uint64(len(k))
	}
	return size
}

func (set *Set) Add(s string) {
	set.s[s] = struct{}{}
}

func (set *Set) Contains(s string) bool {
	_, ok := set.s[s]
	return ok
}

func (set *Set) Remove(s string) {
	delete(set.s, s)
}

func (set *Set) GetAll() []string {
	strs := make([]string, 0)
	for k := range set.s {
		strs = append(strs, k)
	}
	return strs
}

func (set *Set) Len() int {
	return len(set.s)
}

func (set *Set) Intersect(s Set) *Set {
	res := &Set{s: make(map[string]struct{})}
	for k1 := range set.s {
		for k2 := range s.s {
			if k1 == k2 {
				res.s[k1] = struct{}{}
			}
		}
	}
	return res
}

func (set *Set) Union(s Set) *Set {
	res := &Set{s: make(map[string]struct{})}
	for k1 := range set.s {
		res.s[k1] = struct{}{}
	}
	for k2 := range s.s {
		res.s[k2] = struct{}{}
	}
	return res
}
