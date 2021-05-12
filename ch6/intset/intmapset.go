package intset

type IntMapSet struct {
	// Invariant: for any key in the map, its value is true.
	elems map[int]bool
}

func CreateIntMapSet() IntMapSet {
	return IntMapSet{make(map[int]bool)}
}

func (s *IntMapSet) Elems() []int {
	res := make([]int, len(s.elems))
	i := 0
	for k := range s.elems {
		res[i] = k
		i++
	}
	return res
}

func (s *IntMapSet) Add(x int) {
	s.elems[x] = true
}

func (s *IntMapSet) UnionWith(t *IntMapSet) {
	for k := range t.elems {
		s.elems[k] = true
	}
}

func (s *IntMapSet) IntersectWith(t *IntMapSet) {
	for k := range s.elems {
		_, ok := t.elems[k]
		if !ok {
			delete(s.elems, k)
		}
	}
}

func (s *IntMapSet) DifferenceWith(t *IntMapSet) {
	for k := range t.elems {
		_, ok := s.elems[k]
		if ok {
			delete(s.elems, k)
		}
	}
}

func (s *IntMapSet) SymmetricDifference(t *IntMapSet) {
	intersection := CreateIntMapSet()
	for k := range s.elems {
		_, ok := t.elems[k]
		if ok {
			intersection.Add(k)
		}
	}
	s.DifferenceWith(&intersection)
	t.DifferenceWith(&intersection)
	s.UnionWith(t)
}

func (s *IntMapSet) EqualToIntSet(t *IntSet) bool {
	if len(s.elems) != t.Len() {
		return false
	}
	for _, k := range t.Elems() {
		_, ok := s.elems[k]
		if !ok {
			return false
		}
	}
	return true
}
