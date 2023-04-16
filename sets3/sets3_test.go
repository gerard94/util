package sets3

import (
	
	"fmt"
	"git.duniter.org/gerard94/util/alea"
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
	e, ok := i.FirstE()
	for ok {
		fmt.Println(e)
		e, ok = i.NextE()
	}
	fmt.Println()
}

func TestCreate (t *testing.T) {
	const (
		
		nb  = 10
	
	)
	
	printSet(createSet(nb, true))
}

func TestXorUnionInterDiffWithAdd (t *testing.T) {
	const (
		
		repeat = 100
		nb  = 8
	
	)
	for k := 1; k <= repeat; k++ {
		e1 := createSet(nb, false)
		printSet(e1)
		e2 := createSet(nb, false)
		printSet(e2)
		g := e1.XOR(e2)
		printSet(g)
		f := e1.Inter(e2)
		e1.Add(e2)
		f = e1.Diff(f)
		if !f.Equal(g) {
			t.Fail()
		}
		fmt.Println()
	}
}
