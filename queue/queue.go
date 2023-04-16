package queue

import (
	
	M	"git.duniter.org/gerard94/util/misc"

)

type (
	
	Elem [T any] struct {
		next *Elem[T]
		val T
	}
	
	Queue [T any] struct {
		end *Elem[T]
	}

)

func (q *Queue[T]) Init () {
	(*q).end = nil
} //Init

func NewQueue [T any] () *Queue[T] {
	q := new(Queue[T])
	q.Init()
	return q
} //NewQueue

func (q *Queue[T]) IsEmpty () bool {
	return q.end == nil
} //IsEmpty

func (q *Queue[T]) Push (val T) {
	e := &Elem[T]{val: val}
	if q.end == nil {
		e.next = e
		q.end = e
	} else {
		e.next = q.end.next
		q.end.next = e
		q.end = e
	}
} //Push

func (q *Queue[T]) PushN (n int, val T) (ok bool) {
	e := &Elem[T]{val: val}
	if q.end == nil {
		ok = n == 0
		if ok {
			e.next = e
			q.end = e
		}
	} else {
		ee := q.end
		end := n > 0
		for n > 0 {
			ee = ee.next
			n--
			if ee == q.end {break}
		}
		ok = n == 0
		if ok {
			e.next = ee.next
			ee.next = e
			if ee == q.end && end {
				q.end = e
			}
		}
	}
	return
} //PushN

func (q *Queue[T]) Pull (val *T) (ok bool) {
	ok = q.end != nil
	if ok {
		e := q.end.next
		q.end.next = e.next
		if q.end == e {
			q.end = nil
		}
		*val = e.val
	}
	return
} //Pull

func (q *Queue[T]) Pop () (val T) {
	ok := q.Pull(&val)
	M.Assert(ok, 100)
	return
} //Pop

func (q *Queue[T]) PullN (n int, val *T) (ok bool) {
	M.Assert(n >= 0, 20)
	ok = q.end != nil
	if ok {
		e := q.end
		for n > 0 {
			e = e.next
			n--
			if e.next == q.end {break}
		}
		ok = n == 0
		if ok {
			ee := e.next
			*val = ee.val
			if ee == e {
				q.end = nil
			} else if ee == q.end {
				q.end = e
			}
			e.next = ee.next
		}
	}
	return
} //PullN

func (q *Queue[T]) Next (e *Elem[T]) *Elem[T] {
	if q.end == nil {
		return nil
	}
	if e == nil {
		return q.end.next
	} else if e == q.end {
		return nil
	} else {
		return e.next
	}
} //Next

func (e *Elem[T]) Val () T {
	return e.val
} //Val
