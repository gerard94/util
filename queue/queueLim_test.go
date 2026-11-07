package queue

import (
	
	M	"git.duniter.org/gerard94/util/misc"
		"fmt"
		"testing"

)

func TestQueueLim (t *testing.T) {
	fmt.Println("TestQueueLim")
	var q QueueLim[int]
	(&q).Init(5)
	q.Push(9)
	q.Push(8)
	q.Push(7)
	q.Push(6)
	q.Push(5)
	b := q.Push(4); M.Assert(!b)
	var i int
	for ok := q.Pull(&i); ok; ok = q.Pull(&i) {
		fmt.Println(i)
	}
	fmt.Println()
}
