package orderedSets

import (
	
	A	"github.com/gerard94/util/avlT"
		"github.com/gerard94/util/alea"
		"fmt"
		"testing"

)

type comp int

// comp is a A.Comparer
func (i comp) Compare (j comp) A.Comp {
	switch {
	case i < j:
		return A.Lt
	case i > j:
		return A.Gt
	default:
		return A.Eq
	}
}

func createSet (max int, prop float64, print bool) *Set[comp] {
	nb := int(prop * float64(max))
	s := NewSet[comp]()
	for i := 1; i < nb; i++ {
		n := comp(alea.IntRand(0, int64(max)))
		if print {
			fmt.Println(n)
		}
		s.Incl(n)
	}
	if print {
		fmt.Println()
	}
	return s
}

func printSet (s *Set[comp]) {
	i := s.Attach()
	for e, ok := i.FirstE(); ok; e, ok = i.NextE() {
		fmt.Println(e)
	}
	fmt.Println()
}

func TestCreate (t *testing.T) {
	const (
		max  = 10
		prop = 0.9
	)
	fmt.Println("TestCreate")
	printSet(createSet(max, prop, true))
}

func TestXorUnionInterDiff (t *testing.T) {
	const (
		max  = 100
		prop = 0.9
	)
	fmt.Println("TestXorUnionInterDiff")
	s1 := createSet(max, prop, false)
	s2 := createSet(max, prop, false)
	f := s1.Union(s2).Diff(s1.Inter(s2))
	g := s1.XOR(s2)
	if !f.Equal(g) {
		t.Fail()
	}
	fmt.Println("s1:")
	printSet(s1)
	fmt.Println("s2:")
	printSet(s2)
	fmt.Println("s1 XOR s2:")
	printSet(g)
}

func TestSubset (t *testing.T) {
	const (
		max  = 200
		prop = 0.5
	)
	fmt.Println("TestSubset")
	s1 := createSet(max, prop, false)
	s2 := s1.Copy()
	i := s1.Attach()
	if s2.NbElems() > 0 {
		s2.Excl(s2.ElemOfRank(0))
	}
	for e, ok := i.FirstE(); ok; e, ok = i.NextE() {
		if alea.IntRand(0, 2) == 0 {
			s2.Excl(e)
		}
	}
	if !s1.Subset(s1) || !s2.Subset(s2) || !s2.Subset(s1) || s1.Subset(s2) || s2.In(comp(0)) {
		t.Fail()
	}
	fmt.Println("s1:")
	printSet(s1)
	fmt.Println("s2:")
	printSet(s2)
}
