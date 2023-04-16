package smallSets

import (
	
	M	"git.duniter.org/gerard94/util/misc"

)

const (
	
	MinSet = 0
	MaxSet = 63

)

type (

	Set uint64

)

func NewSet () Set {
	return Set(0)
}

func Full () Set {
	return Set(M.MaxUint64)
}

func MakeSet (values ... int) Set {
	set := Set(0)
	for _, val := range values {
		M.Assert(MinSet <= val && val <= MaxSet, val, 20)
		set |= 1 << uint(val)
	}
	return set
}

func (s Set) IsEmpty () bool {
	return s == 0
}

func (s Set) NbElems () int {
	n := 0
	curr := Set(1)
	for i := MinSet; i <= MaxSet; i++ {
		if curr & s != 0 {
			n++
		}
		curr <<= 1
	}
	return n
}

func (set1 Set) Union (set2 Set) Set {
	return set1 | set2
}

func (set1 Set) Inter (set2 Set) Set {
	return set1 & set2
}

func (set1 Set) Diff (set2 Set) Set {
	return set1 &^ set2
}

func (set1 Set) XOR (set2 Set) Set {
	return set1 ^ set2
}

func (set *Set) Incl (value int) {
	M.Assert(MinSet <= value && value <= MaxSet, value, 20)
	*set |= 1 << uint(value)
}

func (set *Set) Excl (value int) {
	M.Assert(MinSet <= value && value <= MaxSet, value, 20)
	*set &^= 1 << uint(value)
}

func (set Set) In (value int) bool {
	M.Assert(MinSet <= value && value <= MaxSet, value, 20)
	return 1 << uint(value) & set != 0
}

func (s1 Set) Equal (s2 Set) bool {
	return s1.XOR(s2).IsEmpty()
}

func (s1 Set) Subset (s2 Set) bool {
	return s1.Diff(s2).IsEmpty()
}

func (s Set) FirstE () (int, bool) {
	curr := Set(1)
	for i := MinSet; i <= MaxSet; i++ {
		if curr & s != 0 {
			return i, true
		}
		curr <<= 1
	}
	return 0, false
}

func (s Set) NextE (e int) (int, bool) {
	curr := MakeSet(e + 1)
	for i := e + 1; i <= MaxSet; i++ {
		if curr & s != 0 {
			return i, true
		}
		curr <<= 1
	}
	return 0, false
}
