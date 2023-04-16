package storage

import (
	M	"util/misc"
		"testing"
		"fmt"
)

func TestRecHeap (t *testing.T) {
	heap := NewRecHeap[int]()
	s := [...]int{4, 6, 9, 0, 1, 3, 5}
	ss := make([]*int, len(s))
	for i, n := range s {
		ss[i] = heap.NewRec()
		*ss[i] = n
	}
	for _, n := range s {
		fmt.Println(n)
	}
	fmt.Println()
	for _, a := range ss {
		fmt.Println(*a)
	}
	fmt.Println()
	for _, p := range ss {
		heap.DisposeRec(&p)
	}
	for i := 0; i < len(ss); i++ {
		ss[i] = heap.NewRec()
	}
	for _, p := range ss {
		fmt.Println(*p)
	}
}

func TestArrHeap (t *testing.T) {
	heap := NewArrHeap[int]()
	a := heap.NewArr(9)
	b := *a
	s := [...]int{4, 6, 2, 7, 5, 9, 1, 8, 3}
	for i, n := range s{
		b[i] = n
	}
	fmt.Println(a)
	heap.DisposeArr(&a)
	M.Assert(a == nil)
	a = heap.NewArr(9)
	fmt.Println(a)
}
