// The package orderedSets implements sets of generic types, which have to be both avlT comparers and copiers (see util/avlT). These types of sets using avl trees are ordered by design.
package orderedSets

import (
	
	A	"github.com/gerard94/util/avlT"
	M	"github.com/gerard94/util/misc"

)

type (
	
	Set[T A.Comparer[T]] struct {
		*A.TreeS[T]
	}

	SetIterator[T  A.Comparer[T]] struct {
		s *Set[T]
		e *A.Elem[T]
	}

)

// Create and return a new empty set.
func NewSet [T A.Comparer[T]] () *Set[T] {
	s := &Set[T]{TreeS: A.NewS[T]()}
	return s
}

// Test whether 's' is empty.
func (s *Set [T]) IsEmpty () bool {
	return s.TreeS.IsEmpty()
}

// Return the number of elements in 's'.
func (s *Set [T]) NbElems () int {
	return s.NumberOfElems()
}

// Create and return a copy of 's'.
func (s *Set[T]) Copy () *Set[T] {
	return &Set[T]{TreeS: s.TreeS.Copy()}
}

// Put in 's1' the union of 's1' and 's2'.
func (s1 *Set [T]) Add (s2 *Set[T]) {
	for e := s2.Next(nil); e != nil; e = s2.Next(e) {
		s1.SearchIns(e.Val())
	}
}

// Return the union of 's1' and 's2'.
func (s1 *Set[T]) Union (s2 *Set[T]) *Set[T] {
	if s1.NbElems() < s2.NbElems() {
		s1, s2 = s2, s1
	}
	s := s1.Copy()
	s.Add(s2)
	return s
}

// Return the intersection of 's1' and 's2'.
func (s1 *Set[T]) Inter (s2 *Set[T]) *Set[T] {
	if s1.NbElems() > s2.NbElems() {
		s1, s2 = s2, s1
	}
	s := NewSet[T]()
	for e := s1.Next(nil); e != nil; e = s1.Next(e) {
		ee := e.Val()
		if _, found, _ := s2.Search(ee); found {
			s.SearchIns(ee)
		}
	}
	return s
}

// Put in 's' the difference between 's1' and 's2'.
func diff[T A.Comparer[T]] (s1, s2, s *Set[T]) {
	for e := s1.Next(nil); e != nil; e = s1.Next(e) {
		ee := e.Val()
		if _, found, _ := s2.Search(ee); !found {
			s.SearchIns(ee)
		}
	}
}

// Return the difference between 's1' and 's2'.
func (s1 *Set[T]) Diff (s2 *Set[T]) *Set[T] {
	s := NewSet[T]()
	diff[T](s1, s2, s)
	return s
}

// Return the exclusive union of 's1' and 's2'.
func (s1 *Set[T]) XOR (s2 *Set[T]) *Set[T] {
	s := NewSet[T]()
	diff[T](s1, s2, s)
	diff[T](s2, s1, s)
	return s
}

// Include the element 'e' into the set 's'.
func (s *Set[T]) Incl (e T) (didInclude bool) {
	_, found, _ := s.SearchIns(e)
	return !found
}

// Exclude the element 'e' from the set 's'.
func (s *Set[T]) Excl (e T) (didExclude bool) {
	return s.Delete(e)
}

// Test whether the element 'e' is included in the set 's'.
func (s *Set[T]) In (e T) bool {
	_, found, _ := s.Search(e)
	return found
}

// Test whether 's1' = 's2'.
func (s1 *Set[T]) Equal (s2 *Set[T]) bool {
	return s1.XOR(s2).IsEmpty()
}

// Test whether 's1' is a subset of 's2'.
func (s1 *Set[T]) Subset (s2 *Set[T]) bool {
	return s1.Diff(s2).IsEmpty()
}

// Return the element of 's' of rank 'rank' (from 0).
func (s *Set[T]) ElemOfRank (rank int) T {
	M.Assert(0 <= rank && rank < s.NbElems(), rank, s.NbElems(), 20)
	e, _ := s.Find(rank + 1)
	return e.Val()
}

// Attach the set 's' to a SetIterator and return the later.
func (s *Set[T]) Attach () *SetIterator[T] {
	return &SetIterator[T]{s: s}
}

// On return, 'e' contains the first element of the set attached to 'i'. Return true in 'ok' if such an element exists. Usage: for e, ok := i.FirstE(); ok; e, ok = i.NextE() { ... }
func (i *SetIterator[T]) FirstE () (e T, ok bool) {
	i.e = i.s.Next(nil)
	ok = i.e != nil
	if ok {
		e = i.e.Val()
	}
	return
}

// On return, 'e' contains the next element of the set attached to 'i'. Return true in 'ok' if such an element exists. i.FirstE must have been called once before i.NextE. Usage: for e, ok := i.FirstE(); ok; e, ok = i.NextE() { ... }
func (i *SetIterator[T]) NextE () (e T, ok bool) {
	i.e = i.s.Next(i.e)
	ok = i.e != nil
	if ok {
		e = i.e.Val()
	}
	return
}
