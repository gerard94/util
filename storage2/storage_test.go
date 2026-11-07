package storage2

import (
	"testing"
	//M	"git.duniter.org/gerard94/util/misc"
	"fmt"
	"unsafe"
)

func TestRecHeap (t *testing.T) {
	heap := NewRecHeap(10, 5)
	a := heap.NewRec()
	au := uintptr(unsafe.Pointer(a))
	//M.Want(au % 5 == 4, t)
	fmt.Println(au)
	b := heap.NewRec()
	bu := uintptr(unsafe.Pointer(b))
	//M.Want(bu % 5 == 4, t)
	fmt.Println(bu)
	heap.DisposeRec(&a)
	au = uintptr(unsafe.Pointer(a))
	fmt.Println(au)
	heap.DisposeRec(&b)
	a = heap.NewRec()
	au = uintptr(unsafe.Pointer(a))
	fmt.Println(au)
	b = heap.NewRec()
	bu = uintptr(unsafe.Pointer(b))
	fmt.Println(bu)
}

func TestArrHeap (t *testing.T) {
	heap := NewArrHeap(10, 3)
	a := heap.NewArr(9)
	au := uintptr(unsafe.Pointer(a))
	//M.Want(au % 3 == 2, t)
	fmt.Println(au)
	heap.DisposeArr(&a, 9)
	au = uintptr(unsafe.Pointer(a))
	fmt.Println(au)
}
