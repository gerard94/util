/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package storage

/* This module implements a fast heap manager. Objects of different sizes have their instances allocated in different heaps. Free space is held in simply linked chains, one for each object size, so that allocation and deallocation are very fast. Heaps can increase, but never decrease. They can be erased when no more needed. */

import (
	M	"git.duniter.org/gerard94/util/misc"
		"unsafe"
	/*
	"fmt"
	*/
)

const (
	
	blockSize = 100

)

type (
	
	ObjectElem [T any] struct {
		object T
		next *ObjectElem[T]
	}
	
	Objects [T any] []ObjectElem[T]
	
	// A heap block, allocated on Blackbox heap when needed
	heapBlock [T any] struct {
		next *heapBlock[T] // Blocks are linked together, so that they can't be collected
		h Objects[T] // Allocations happen here
	}
	
	// Heap for cells of the same size
	RecHeap [T any] struct {
		head *ObjectElem[T] // First available cell in this heap
		h *heapBlock[T] // Linked chain of blocks of this heap
	}
	
	// Head of the chained blocks for one size of array cells
	blockElem [T any] struct {
		head *ObjectElem[T] // First available cell in this heap
		h *heapBlock[T] // Linked chain of blocks of this heap
	}
	
	blockArray [T any] []blockElem[T] // Array of block chains for every number of elements in allocated arrray chains
	
	ArrHeap [T any] struct { // Heap for array cells
		h blockArray[[]T] // Array of block chains for every number of elements in allocated arrrays the number of elements is i + 1 for h[i]
	}
	
)

// Creates a new block and stores into it as many free cells as possible in a linked chain
func newBlock [T any] (h **heapBlock[T]) *ObjectElem[T] {
	b := new(heapBlock[T]); M.Assert(b != nil, 100)
	b.next = *h // Protection against garbage collector
	*h = b
	b.h = make(Objects[T], blockSize); M.Assert(b.h != nil, 101)
	q := (*ObjectElem[T])(nil)
	for i := 0; i < len(b.h); i++ { // range doesn't work!!!
		p := &b.h[i]
		p.next = q
		q = p
	}
	return q
}

// Allocates into a a new cell of fixed size
func (heap *RecHeap[T]) NewRec () *T {
	if heap.head == nil { // No more free cell: allocate a new block
		heap.head = newBlock[T](&heap.h)
	}
	a := heap.head
	heap.head = a.next //Follow the linked chain
	return &a.object
}

// Deallocates the cell at Address a by putting it into the link chained
func (heap *RecHeap[T]) DisposeRec (a **T) {
	p := (*ObjectElem[T])(unsafe.Pointer(*a))
	p.next = heap.head // Insert new free cell at the beginning of the linked chain
	heap.head = p
	*a = nil
}

// Allocates into a a new cell for an array of elemNb elements of fixed size
func (heap *ArrHeap[T]) NewArr (cellNb int) (a *[]T) {

	newA := func (el *blockElem[[]T], nb int) {
		if el.h == nil { // Initialize uninitialized blockElem
			el.head = nil
		}
		if el.head == nil { // No more free cell: allocate a new block
			el.head = newBlock[[]T](&el.h)
		}
		p := el.head
		el.head = p.next //Follow the linked chain
		if len(p.object) == 0 {
			p.object = make([]T, nb)
		}
		a = &p.object
	}

	M.Assert(cellNb > 0, 20)
	i := len(heap.h)
	if cellNb > i { // blockArray too short: replace it
		arr := make(blockArray[[]T], cellNb); M.Assert(arr != nil, 100)
		j := cellNb
		for j > i { // Initialize the new BlockElems...
			j--
			arr[j].h = nil // Shows that the blockElem is not yet initialised
		}
		for i > 0 { // ... and copy the old ones
			i--
			arr[i] = heap.h[i]
		}
		heap.h = arr
	}
	newA(&heap.h[cellNb - 1], cellNb) // Allocate the new cell in the correct heap
	return
}

// Deallocates the cell at Address a and containing an array cell of elemNb elements by putting it in the link chained
func (heap *ArrHeap[T]) DisposeArr (a **[]T) {
	M.Assert(a != nil && *a != nil && **a != nil, 20)
	cellNb := len(**a)
	M.Assert(cellNb > 0, 21)
	p := (*ObjectElem[[]T])(unsafe.Pointer(*a))
	p.next = heap.h[cellNb - 1].head
	heap.h[cellNb - 1].head = p
	*a = nil
}

// Creates a new heap for allocation of cells with the unique size recSize
func NewRecHeap [T any] () (heap *RecHeap[T]) {
	heap = new(RecHeap[T]); M.Assert(heap != nil, 100)
	heap.h = nil // No block yet...
	heap.head = nil // ... and no free cell
	return
}

// Creates a new heap for allocation of array cells, whose sizes are multiple of the same element size elemSize
func NewArrHeap [T any] () (heap *ArrHeap[T]) {
	heap = new(ArrHeap[T]); M.Assert(heap != nil, 100)
	heap.h = make(blockArray[[]T], 0) // No element in allocated arrays, for the moment
	M.Assert(heap.h != nil, 101)
	return
}

func (heap *RecHeap[T]) Stats () (nb, size int) {
	n := 0
	for b := heap.h; b != nil; b = b.next {
		n++
	}
	nb = blockSize * n
	for a := heap.head; a != nil; a = a.next {
		nb--
	}
	var t T
	size = nb * int(unsafe.Sizeof(t))
	return
}

func (heap *ArrHeap[T]) Stats () (nb, nbElems, size int) {
	nb = 0
	nbElems = 0
	for s, el := range heap.h {
		if el.h != nil {
			n := 0
			for b := el.h; b != nil; b = b.next{
				n++
			}
			n *= blockSize
			for a := el.head; a != nil; a = a.next {
				n--
			}
			nb += n
			nbElems += n * (s + 1)
		}
	}
	var t T
	size = nbElems * int(unsafe.Sizeof(t))
	return
}

func init () {
	a := new(ObjectElem[int])
	M.Assert(uintptr(unsafe.Pointer(a)) == uintptr(unsafe.Pointer(&a.object)), 100)
}
