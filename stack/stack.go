package stack

import (
	
	M	"github.com/gerard94/util/misc"

)

const (
	
	aSize = 200

)

type (
	
	Stack [T any] struct {
		end *elem[T]
	}
	
	elem [T any] struct {
		next *elem[T]
		val T
	}
	
	Iterator [T any] struct {
		st *Stack[T]
		e *elem[T]
	}

)

func (st *Stack[T]) Init () {
	st.end = nil
} //Init

func NewStack [T any] () *Stack[T] {
	st := new(Stack[T])
	st.Init()
	return st
} //NewStack

func (st Stack[T]) IsEmpty () bool {
	return st.end == nil
} //IsEmpty

func (st *Stack[T]) Push (e T) {
	st.end = &elem[T]{val: e, next: st.end}
} //Push

func (st *Stack[T]) Pull (e *T) (ok bool) {
	ok = st.end != nil
	if ok {
		*e = st.end.val
		st.end = st.end.next
	}
	return
} //Pull

func (st *Stack[T]) Pop () T {
	var e T
	ok := st.Pull(&e)
	M.Assert(ok, 100)
	return e
} //Pop

func (st *Stack[T]) NewIterator () *Iterator[T] {
	return &Iterator[T]{st: st, e: nil}
} //Iter

func (i *Iterator[T]) Next (first bool, e *T) (ok bool) {
	if first {
		i.e = i.st.end
	}
	ok = i.e != nil
	if ok {
		*e = i.e.val
		i.e = i.e.next
	}
	return
} //Next
