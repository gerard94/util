/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

//#include <stdlib.h>
package storage

// unsafe.Pointer(s) may be moved by the garbage collector, but not uintptr(s).

/* This module implements a fast heap manager. Objects of different sizes have their instances allocated in different heaps. Free space is held in simply linked chains, one for each object size, so that allocation and deallocation are very fast. Heaps can increase, but never decrease. They can be erased when no more needed. */

import (
	M	"git.duniter.org/gerard94/util/misc"
		"C"
	U	"unsafe"
		/*
		"fmt"
		*/
)

const (
	heapSize = 0x1000
	maxLength = 0x1000
	addressSize = 8 // Available cells are linked in a single linked chain every link is an Address, of size addressSize
	nilA Address = 0 // Last link of the linked chain of available cells
)

type (
	
	Address = uintptr // Must be converted to the right pointer type
	
	bytes = [maxLength]byte
	
	// A heap block, allocated on Blackbox heap when needed
	heapBlock struct {
		next *heapBlock // Blocks are linked together, so that they can't be collected
		nbEl int
		h *bytes // Allocations happen here
	}
	
	// Heap for cells of the same size
	RecHeap struct {
		recSize, // Size of blocks, greatest multiple of recSize less than heapSizeMax
		align uintptr
		head Address // First available cell in this heap
		h *heapBlock // Linked chain of blocks of this heap
	}
	
	// Head of the chained blocks for one size of array cells
	blockElem struct {
		arrSize uintptr
		head Address
		h *heapBlock
	}
	
	blockArray []blockElem // Array of block chains for every number of elements in allocated arrray chains
	
	ArrHeap struct { // Heap for array cells, whose sizes are multiple of the same element size elemSize
		elemSize,
		align uintptr
		h blockArray // Array of block chains for every number of elements in allocated arrrays the number of elements is i + 1 for h[i]
	}
	
)

// Writes the Address value v at the address a
func setVal (a, v Address) {
	*(*Address)(U.Pointer(a)) = v
}

// Returns the Address value at Address a
func val (a Address) Address {
	return *(*Address)(U.Pointer(a))
}

// Creates a new block and stores into it as many free cells as possible in a linked chain
func newBlock (h **heapBlock, elemSize, align uintptr) Address {
	b := new(heapBlock); M.Assert(b != nil, 100)
	b.next = *h // Protection against garbage collector
	*h = b
	bl := U.Pointer(C.malloc(heapSize))
	b.h = (*bytes)(bl); M.Assert(b.h != nil, 101)
	p := uintptr(bl)
	start := p % align
	if start != 0 {
		start = align - start
	}
	p += start
	M.Assert(heapSize - start >= elemSize, 102)
	n := (heapSize - start) / elemSize
	b.nbEl = int(n)
	pMax := n * elemSize + p
	q := nilA
	for ; p < pMax; p += elemSize {
		setVal(p, q)
		q = p
	}
	return q
}

// Allocates into a a new cell of fixed size
func (heap *RecHeap) NewRec () (a Address) {
	if heap.head == nilA { // No more free cell: allocate a new block
		heap.head = newBlock(&heap.h, heap.recSize, heap.align)
	}
	a = heap.head
	heap.head = val(a) //Follow the linked chain
	return
}

// Deallocates the cell at Address a by putting it into the link chained
func (heap *RecHeap) DisposeRec (a *Address) {
	setVal(*a, heap.head) // Insert new free cell at the beginning of the linked chain
	heap.head = *a
	*a = nilA
}

// Initialize a blockElem
func initBlockElem (el *blockElem, elemSize uintptr, elemNb int) {
	el.arrSize = M.MaxUP(uintptr(elemNb) * elemSize, addressSize)
	el.head = nilA // No free cell yet
}

// Allocates into a a new cell for an array of elemNb elements of fixed size
func (heap *ArrHeap) NewArr (elemNb int) (a Address) {

	newA := func (el *blockElem) {
		if el.h == nil { // Initialize uninitialized blockElem
			initBlockElem(el, heap.elemSize, elemNb)
		}
		if el.head == nilA { // No more free cell: allocate a new block
			el.head = newBlock(&el.h, el.arrSize, heap.align)
		}
		a = el.head
		el.head = val(a) //Follow the linked chain
	}

	M.Assert(elemNb > 0, 20)
	i := len(heap.h)
	if elemNb > i { // blockArray too short: replace it
		M.Assert(uintptr(elemNb) * heap.elemSize <= heapSize, 21)
		arr := make(blockArray, elemNb); M.Assert(arr != nil, 100)
		j := elemNb
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
	newA(&heap.h[elemNb - 1]) // Allocate the new cell in the correct heap
	return
}

// Deallocates the cell at Address a and containing an array cell of elemNb elements by putting it in the link chained
func (heap ArrHeap) DisposeArr (a *Address, elemNb int) {
	M.Assert(*a != nilA, 20)
	M.Assert(elemNb > 0, 21)
	setVal(*a, heap.h[elemNb - 1].head)
	heap.h[elemNb - 1].head = *a
	*a = nilA
}

// Creates a new heap for allocation of cells with the unique size recSize
func NewRecHeap (recSize, align uintptr) (heap *RecHeap) {
	M.Assert(recSize > 0, 20)
	M.Assert(recSize <= heapSize, 21)
	M.Assert(align > 0, 22)
	heap = new(RecHeap); M.Assert(heap != nil, 100)
	recSize = M.MaxUP(recSize, addressSize) // Must keep room for free cells and their integer links
	heap.recSize = recSize
	heap.align = align
	heap.h = nil // No block yet...
	heap.head = nilA // ... and no free cell
	return
}

// Creates a new heap for allocation of array cells, whose sizes are multiple of the same element size elemSize
func NewArrHeap (elemSize, align uintptr) (heap *ArrHeap) {
	M.Assert(elemSize > 0, 20)
	M.Assert(elemSize <= heapSize, 21)
	M.Assert(align > 0, 22)
	heap = new(ArrHeap); M.Assert(heap != nil, 100)
	heap.elemSize = elemSize
	heap.align = align
	heap.h = make(blockArray, 0) // No element in allocated arrays, for the moment
	M.Assert(heap.h != nil, 101)
	return
}

func (heap *RecHeap) Stats () (nb, size int) {
	nb = 0
	for b := heap.h; b != nil; b = b.next {
		nb += b.nbEl
	}
	for a := heap.head; a != nilA; a = val(a) {
		nb--
	}
	size = nb * int(heap.recSize)
	return
}

func (heap *ArrHeap) Stats () (nb, nbElems, size int) {
	for s, el := range heap.h {
		if el.h != nil {
			n := 0
			for b := el.h; b != nil; b = b.next{
				n += b.nbEl
			}
			for a := el.head; a != nilA; a = val(a) {
				n--
			}
			nb += n
			nbElems += n * (s + 1)
		}
	}
	size = nbElems * int(heap.elemSize)
	return
}

func init () {
	var x byte
	M.Assert(U.Sizeof(x) == 1, 20)
	var y Address
	M.Assert(U.Sizeof(y) == addressSize, 21)
	M.Assert(Address(U.Pointer(nil)) == nilA, 22)
}
