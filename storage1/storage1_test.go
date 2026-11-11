package storage1

import (
	M	"github.com/gerard94/util/misc"
	U	"unsafe"
		"testing"
		"fmt"
)

func TestRecHeap (t *testing.T) {
	fmt.Println("TestRecHeap")
	var n int
	heap := NewRecHeap(U.Sizeof(n), U.Alignof(n))
	s := [...]int{4, 6, 9, 0, 1, 3, 5}
	ss := new([7]*int)
	for i, n := range s {
		ss[i] = (*int)(U.Pointer(heap.NewRec()))
		*ss[i] = n
	}
	for _, n := range s {
		fmt.Println(n)
	}
	fmt.Println()
	for i, a := range ss {
		fmt.Println(i, *a)
	}
	fmt.Println()
	for _, p := range ss {
		heap.DisposeRec((*Address)(U.Pointer(&p)))
	}
	for i := 0; i < len(ss); i++ {
		ss[i] = (*int)(U.Pointer(heap.NewRec()))
	}
	for _, p := range ss {
		fmt.Println(*p)
	}
	fmt.Println()
}

func TestArrHeap (t *testing.T) {
	fmt.Println("TestArrHeap")
	var n int
	heap := NewArrHeap(U.Sizeof(n), U.Alignof(n))
	a := (*[9]int)(U.Pointer(heap.NewArr(9)))
	b := *a
	s := [...]int{4, 6, 2, 7, 5, 9, 1, 8, 3}
	fmt.Println(s)
	fmt.Println()
	for i, n := range s{
		b[i] = n
	}
	fmt.Println(b)
	fmt.Println()
	heap.DisposeArr((*Address)(U.Pointer(&a)), 9)
	M.Assert(a == nil)
	a = (*[9]int)(U.Pointer(heap.NewArr(9)))
	b = *a
	fmt.Println(b)
	fmt.Println()
}
