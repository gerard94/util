package sets

import (
		"fmt"
	R	"math/rand"
		"testing"
)

func createSet(max int, prop float32, print bool) *Set {
	nb := int(prop * float32(max))
	s := NewSet()
	for i := 1; i <= nb; i++ {
		n := int(R.Int31n(int32(max)))
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

func printSet(s *Set) {
	i := s.Attach()
	
	for e, ok := i.FirstE(); ok; e, ok = i.NextE() {
		fmt.Println(e)
		
	}
	fmt.Println()
}

func TestCreate(t *testing.T) {
	const (
		max  = 10
		prop = 0.9
	)
	fmt.Println("TestCreate")
	printSet(createSet(max, prop, true))
}

func TestXorUnionInterDiff(t *testing.T) {
	const (
		max  = 100
		prop = 0.9
	)
	fmt.Println("TestXorUnionInterDiff")
	e1 := createSet(max, prop, false)
	e2 := createSet(max, prop, false)
	f := e1.Union(e2).Diff(e1.Inter(e2))
	g := e1.XOR(e2)
	if !f.Equal(g) {
		t.Fail()
	}
	fmt.Println("e1:")
	printSet(e1)
	fmt.Println("e2:")
	printSet(e2)
	fmt.Println("e1 XOR e2:")
	printSet(g)
}
