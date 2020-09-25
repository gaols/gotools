package ds

type Stringset struct {
	s map[string]bool
}

func (set *Stringset) Size() int {
	return len(set.s)
}

func (set *Stringset) IsEmpty() bool {
	return len(set.s) == 0
}

func (set *Stringset) Contains(c string) bool {
	_, ok := set.s[c]
	return ok
}

func (set *Stringset) toSlice() []string {
	keys := make([]string, 0, len(set.s))
	for k := range set.s {
		keys = append(keys, k)
	}
	return keys
}

func (set *Stringset) Add(c string) {
	set.s[c] = true
}

func (set *Stringset) Remove(c string) {
	delete(set.s, c)
}

func (set *Stringset) RemoveAll(c []string) {
	for _, v := range c {
		delete(set.s, v)
	}
}

func (set *Stringset) ContainsAll(c []string) bool {
	for _, v := range c {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set *Stringset) Clear() {
	for k := range set.s {
		delete(set.s, k)
	}
}

func NewStringset() *Stringset {
	s := make(map[string]bool)
	set := &Stringset{
		s: s,
	}
	return set
}
