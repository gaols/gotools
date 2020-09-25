package ds

type Intset struct {
	s map[int]bool
}

func (set *Intset) Size() int {
	return len(set.s)
}

func (set *Intset) IsEmpty() bool {
	return len(set.s) == 0
}

func (set *Intset) Contains(c int) bool {
	_, ok := set.s[c]
	return ok
}

func (set *Intset) toSlice() []int {
	keys := make([]int, 0, len(set.s))
	for k := range set.s {
		keys = append(keys, k)
	}
	return keys
}

func (set *Intset) Add(c int) {
	set.s[c] = true
}

func (set *Intset) Remove(c int) {
	delete(set.s, c)
}

func (set *Intset) RemoveAll(c []int) {
	for _, v := range c {
		delete(set.s, v)
	}
}

func (set *Intset) ContainsAll(c []int) bool {
	for _, v := range c {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set *Intset) Clear() {
	for k := range set.s {
		delete(set.s, k)
	}
}

func NewIntset() *Intset {
	s := make(map[int]bool)
	set := &Intset{
		s: s,
	}
	return set
}
