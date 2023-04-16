package sets3

// The module "util/sets3" implements the type Set, a set of strings, represented by A.Tree(s) of Elem(s).

import (
	
	A	"git.duniter.org/gerard94/util/avl"
	M	"git.duniter.org/gerard94/util/misc"

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

// *Elem is a A.Copier
func (e *Elem) Copy () A.Copier {
	ee := *e
	return &ee
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

// Creates and returns a copy of s.
func (s *Set) Copy () *Set {
	return (*Set)((*A.Tree)(s).Copy())
}

// Puts in s1 the union of s1 and s2
func (s1 *Set) Add (s2 *Set) {
	t1 := (*A.Tree)(s1)
	t2 := (*A.Tree)(s2)
	e := t2.Next(nil)
	for e != nil {
		t1.SearchIns(e.Val().(*Elem))
		e = t2.Next(e)
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
	e := t1.Next(nil)
	for e != nil {
		ee := e.Val().(*Elem)
		if _, found, _ := t2.Search(ee); found {
			t.SearchIns(ee)
		}
		e = t1.Next(e)
	}
	return s
}

// Puts in t the difference between t1 and t2
func diff (t1, t2, t *A.Tree) {
	e := t1.Next(nil)
	for e != nil {
		ee := e.Val().(*Elem)
		if _, found, _ := t2.Search(ee); !found {
			t.SearchIns(ee)
		}
		e = t1.Next(e)
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

// Attach the set s to a SetIterator and return the later.
func (s *Set) Attach () *SetIterator {
	return &SetIterator{t: (*A.Tree)(s)}
}

// On return, e contains the first element of the set attached to i. Returns true in ok if such an element exists. Usage: e, ok := i.FirstE(); for ok { ... e, ok = i.NextE()}
func (i *SetIterator) FirstE () (e string, ok bool) {
	i.e = i.t.Next(nil)
	ok = i.e != nil
	if ok {
		e = string(*i.e.Val().(*Elem))
	}
	return
}

// On return, e contains the next element of the set attached to i. Returns true in ok if such an element exists. i.FirstE must have been called once before i.NextE. Usage: e, ok := i.FirstE(); for ok { ... e, ok = i.NextE()}
func (i *SetIterator) NextE () (e string, ok bool) {
	i.e = i.t.Next(i.e)
	ok = i.e != nil
	if ok {
		e = string(*i.e.Val().(*Elem))
	}
	return
}
