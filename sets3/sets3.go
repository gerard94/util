package sets3

// The module "util/sets3" implements the type Set, a set of strings, represented by A.Tree(s) of Elem(s).

import (
	
	A	"github.com/gerard94/util/avl"
	M	"github.com/gerard94/util/misc"

)

type (
	
	Elem string
	
	Set A.Tree

	SetIterator struct {
		t *A.Tree
		e *A.Elem
	}

)

// *Elem is a A.Comparer
func (e1 *Elem) Compare (e2 A.Comparer) A.Comp {
	ee2, ok := e2.(*Elem); M.Assert(ok, 100)
	if *e1 < *ee2 {
		return A.Lt
	}
	if *e1 > *ee2 {
		return A.Gt
	}
	return A.Eq
}

// Creates and returns a new empty set.
func NewSet () *Set {
	return (*Set)(A.New())
}

// Tests whether s is empty.
func (s *Set) IsEmpty () bool {
	return (*A.Tree)(s).IsEmpty()
}

// Return the number of elements in s
func (s *Set) NbElems () int {
	return (*A.Tree)(s).NumberOfElems()
}

// Creates and returns a deep copy of s.
func (s *Set) Copy () *Set {
	return (*Set)((*A.Tree)(s).Copy())
}

// Puts in s1 the union of s1 and s2
func (s1 *Set) Add (s2 *Set) {
	t1 := (*A.Tree)(s1)
	t2 := (*A.Tree)(s2)
	for e := t2.Next(nil); e != nil; e = t2.Next(e) {
		t1.SearchIns(e.Val().(*Elem))
	}
}

// Returns the union of s1 and s2.
func (s1 *Set) Union (s2 *Set) *Set {
	if s1.NbElems() < s2.NbElems() {
		s1, s2 = s2, s1
	}
	s := s1.Copy()
	s.Add(s2)
	return s
}

// Returns the intersection of s1 and s2.
func (s1 *Set) Inter (s2 *Set) *Set {
	if s1.NbElems() > s2.NbElems() {
		s1, s2 = s2, s1
	}
	t1 := (*A.Tree)(s1)
	t2 := (*A.Tree)(s2)
	s := NewSet()
	t := (*A.Tree)(s)
	for e := t1.Next(nil); e != nil; e = t1.Next(e) {
		ee := e.Val().(*Elem)
		if _, found, _ := t2.Search(ee); found {
			t.SearchIns(ee)
		}
	}
	return s
}

// Puts in t the difference between t1 and t2
func diff (t1, t2, t *A.Tree) {
	for e := t1.Next(nil); e != nil; e = t1.Next(e) {
		ee := e.Val().(*Elem)
		if _, found, _ := t2.Search(ee); !found {
			t.SearchIns(ee)
		}
	}
}

// Returns the difference between s1 and s2.
func (s1 *Set) Diff (s2 *Set) *Set {
	t1 := (*A.Tree)(s1)
	t2 := (*A.Tree)(s2)
	s := NewSet()
	t := (*A.Tree)(s)
	diff(t1, t2, t)
	return s
}

// Returns the exclusive union of s1 and s2.
func (s1 *Set) XOR (s2 *Set) *Set {
	t1 := (*A.Tree)(s1)
	t2 := (*A.Tree)(s2)
	s := NewSet()
	t := (*A.Tree)(s)
	diff(t1, t2, t)
	diff(t2, t1, t)
	return s
}

// Includes the string st in the set s.
func (s *Set) Incl (st string) (didInclude bool) {
	t := (*A.Tree)(s)
	e := Elem(st)
	_, found, _ := t.SearchIns(&e)
	return !found
}

// Excludes the string st from the set s.
func (s *Set) Excl (st string) (didExclude bool) {
	t := (*A.Tree)(s)
	e := Elem(st)
	return t.Delete(&e)
}

// Tests whether the string st is included in the set s.
func (s *Set) In (st string) bool {
	t := (*A.Tree)(s)
	e := Elem(st)
	_, found, _ := t.Search(&e)
	return found
}

// Tests whether s1 = s2.
func (s1 *Set) Equal (s2 *Set) bool {
	return s1.XOR(s2).IsEmpty()
}

// Tests whether s1 is a subset of s2.
func (s1 *Set) Subset (s2 *Set) bool {
	return s1.Diff(s2).IsEmpty()
}

// Return the element of 's' of rank 'rank' (from 0) in lexicographic order
func (s *Set) ElemOfRank (rank int) string {
	M.Assert(0 <= rank && rank < s.NbElems(), rank, s.NbElems(), 20)
	t := (*A.Tree)(s)
	e, _ := t.Find(rank + 1)
	return string(*e.Val().(*Elem))
}

// Attach the set s to a SetIterator and return the later.
func (s *Set) Attach () *SetIterator {
	return &SetIterator{t: (*A.Tree)(s)}
}

// On return, e contains the first element of the set attached to i. Returns true in ok if such an element exists. Usage: for e, ok := i.FirstE(); ok; e, ok = i.NextE() { ... }
func (i *SetIterator) FirstE () (e string, ok bool) {
	i.e = i.t.Next(nil)
	ok = i.e != nil
	if ok {
		e = string(*i.e.Val().(*Elem))
	}
	return
}

// On return, e contains the next element of the set attached to i. Returns true in ok if such an element exists. i.FirstE must have been called once before i.NextE. Usage: for e, ok := i.FirstE(); ok; e, ok = i.NextE() { ... }
func (i *SetIterator) NextE () (e string, ok bool) {
	i.e = i.t.Next(i.e)
	ok = i.e != nil
	if ok {
		e = string(*i.e.Val().(*Elem))
	}
	return
}
