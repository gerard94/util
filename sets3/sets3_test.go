package sets3

import (
	
	"fmt"
	"github.com/gerard94/util/alea"
	"testing"

)

const seed = 5473

var r = alea.New()

func createSet (nb int, print bool) *Set {
	
	var (
		
		list = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	
	)
	
	s := NewSet()
	for i := 1; i <= nb; i++ {
		st := list[int(r.IntRand(0, int64(len(list))))]
		if print {
			fmt.Println(st)
		}
		s.Incl(st)
	}
	if print {
		fmt.Println()
	}
	return s
}

func printSet (s *Set) {
	i := s.Attach()
	for e, ok := i.FirstE(); ok; e, ok = i.NextE() {
		fmt.Println(e)
	}
	fmt.Println()
}

func TestCreate (t *testing.T) {
	const (
		
		nb  = 10
	
	)
	
	fmt.Println("TestCreate")
	printSet(createSet(nb, true))
}

func TestXorUnionInterDiffWithAdd (t *testing.T) {
	const (
		
		repeat = 100
		nb  = 8
	
	)
	fmt.Println("TestXorUnionInterDiffWithAdd")
	for k := 1; k <= repeat; k++ {
		e1 := createSet(nb, false)
		fmt.Println("e1:")
		printSet(e1)
		e2 := createSet(nb, false)
		fmt.Println("e2:")
		printSet(e2)
		g := e1.XOR(e2)
		fmt.Println("e1 XOR e2:")
		printSet(g)
		f := e1.Inter(e2)
		h := e1.Union(e2)
		f = h.Diff(f)
		if !f.Equal(g) {
			t.Fail()
		}
		fmt.Println()
	}
}
