package types

type Set map[string]struct{}

func NewEmptySet() Set {
	return Set(make(map[string]struct{}))
}

func NewSet(strs []string) Set {
	set := Set(make(map[string]struct{}))
	for _, str := range strs {
		set.Add(str)
	}
	return set
}

func (set Set) Size() uint64 {
	var size uint64 = 0
	for k := range map[string]struct{}(set) {
		size += uint64(len(k))
	}
	return size
}

func (set Set) Add(s string) {
	set[s] = struct{}{}
}

func (set Set) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

func (set Set) Remove(s string) {
	delete(set, s)
}

func (set Set) GetAll() []string {
	strs := make([]string, len(set))
	for k, _ := range set {
		strs = append(strs, k)
	}
	return strs
}

func (set Set) Len() int {
	return len(set)
}

func (set1 Set) Intersect(set2 Set) Set {
	set := Set(make(map[string]struct{}))
	for k1, _ := range set1 {
		for k2, _ := range set2 {
			if k1 == k2 {
				set[k1] = struct{}{}
			}
		}
	}
	return set
}

func (set1 Set) Union(set2 Set) Set {
	set := Set(make(map[string]struct{}))
	for k1, _ := range set1 {
		set[k1] = struct{}{}
	}
	for k2, _ := range set2 {
		set[k2] = struct{}{}
	}
	return set
}