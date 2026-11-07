package queue

import (
	
	M	"git.duniter.org/gerard94/util/misc"

)

type (
	
	// Element of 'Queue'
	Elem [T any] struct {
		// Next element
		next *Elem[T]
		// Value of element
		val T
	}
	
	// Queue
	Queue [T any] struct {
		// Last element in queue
		end *Elem[T]
	}

)

// Initialisation
func (q *Queue[T]) Init () {
	q.end = nil
} //Init

// Creation of a queue
func NewQueue [T any] () *Queue[T] {
	q := new(Queue[T])
	q.Init()
	return q
} //NewQueue

// Test whether 'q' is empty
func (q Queue[T]) IsEmpty () bool {
	return q.end == nil
} //IsEmpty

// Return the number of elements in 'q'
func (q Queue[T]) Length () int {
	n := 0
	if q.end == nil {
		return n
	}
	e := q.end.next
	for {
		n++
		if e == q.end {break}
		e = e.next
	}
	return n
}

// Add a new element with value 'val' at the rear of 'q'
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

// Add a new element with value 'val' at position 'n' into 'q'; return if there is such a position in 'q'; positions are numbered from 0
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

// Extract the first element at front of 'q' and return its value in 'val'; if 'q' is empty, 'ok' is false and 'val' is not valid, otherwise, 'ok' is true
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

// Simplified version of Pull; panic if 'q' is empty
func (q *Queue[T]) Pop () (val T) {
	ok := q.Pull(&val)
	M.Assert(ok, 100)
	return
} //Pop

// Extract the element at position 'n' in 'q' and return its value in 'val'; if 'q' is empty, or if there is not such a position, 'ok' is false and 'val' is not valid, otherwise, 'ok' is true; positions are numbered from 0
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

// Concatenate 'q1' and 'q2'; result in 'q1', 'q2' emptied
func (q1 *Queue[T]) Concat (q2 Queue[T]) {
	if q2.end != nil {
		if q1.end == nil {
			q1.end = q2.end
		} else {
			q1.end.next, q2.end.next = q2.end.next, q1.end.next
			q1.end = q2.end
			q2.end = nil
		}
	}
} //Concat

// Return the next element in 'q' after 'e'; if 'e' is nil, the first element in 'q' is returned (nil if 'q' is empty); the element returned after the last one in 'q' is nil
// Example:
//	for e := q.Next(nil); e != nil; e = q.Next(e) {
//		v := e.Val()
//		...
//	}
func (q Queue[T]) Next (e *Elem[T]) *Elem[T] {
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

// Return the value of 'e'
func (e Elem[T]) Val () T {
	return e.val
} //Val
