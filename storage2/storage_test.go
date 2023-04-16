package storage

import (
	"testing"
	M	"util/misc"
	"fmt"
)

func TestRecHeap (t *testing.T) {
	heap := NewRecHeap(10, 5)
	a := heap.NewRec()
	M.Want(a % 5 == 0, t)
	fmt.Println(a)
	b := heap.NewRec()
	M.Want(b % 5 == 0, t)
	fmt.Println(b)
	heap.DisposeRec(&a)
	fmt.Println(a)
	heap.DisposeRec(&b)
	a = heap.NewRec()
	fmt.Println(a)
	b = heap.NewRec()
	fmt.Println(b)
}

func TestArrHeap (t *testing.T) {
	heap := NewArrHeap(10, 3)
	a := heap.NewArr(9)
	M.Want(a % 3 == 0, t)
	fmt.Println(a)
	heap.DisposeArr(&a, 9)
	fmt.Println(a)
}
