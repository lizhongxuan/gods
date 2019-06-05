package main

type IntSet map[uint64]struct{}

func NewIntSet() IntSet {
	return make(map[uint64]struct{})
}

func (set IntSet) Add(v uint64) {
	if _, ok := set[v]; ok {
		return
	}
	set[v] = struct{}{}
}

func (set IntSet) IsMember(v uint64) bool {
	if _, ok := set[v]; ok {
		return true
	}
	return false
}

func (set IntSet) Remove(v uint64) {
	if _, ok := set[v]; !ok {
		return
	}
	delete(set, v)
}

func (set IntSet) Clone() IntSet {
	n := make(map[uint64]struct{})
	for k, v := range set {
		n[k] = v
	}
	return n
}
