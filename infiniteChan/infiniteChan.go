package infiniteChan

// Channels with infinite capacity
// Can't be used in select statements 

import (
	
	Q	"github.com/gerard94/util/queue"

)

type (
	
	InfChan [T any] struct {
		c chan T
		q *Q.Queue[T]
	}

)

func (c *InfChan[T]) Init () {
	c.c = make(chan T)
	c.q = Q.NewQueue[T]()
	go func (c *InfChan[T]) {
		for {
			val := <-c.c
			c.q.Push(val)
		}
	}(c)
}

func NewInfChan [T any] () *InfChan[T] {
	c := new(InfChan[T])
	c.Init()
	return c
}

func (c *InfChan[T]) Send (val T) {
	c.c <- val
}

func (c *InfChan[T]) Receive () (T, bool) {
	var t0 T
	if c.q.IsEmpty() {
		return t0, false
	}
	return c.q.Pop(), true
}
