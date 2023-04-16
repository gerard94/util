package queue

import (
	
	M	"git.duniter.org/gerard94/util/misc"
		"testing"
		"fmt"

)

func TestPush (t *testing.T) {
	fmt.Println("TestPush")
	q := NewQueue[int]()
	for i := 0; i <= 9; i++ {
		q.Push(i)
	}
	for e := q.Next(nil); e != nil; e = q.Next(e) {
		fmt.Print(e.Val(), " ")
	}
	fmt.Println(); fmt.Println()
}

func TestPushPull (t *testing.T) {
	fmt.Println("TestPushPull")
	q := NewQueue[int]()
	for i := 0; i <= 9; i++ {
		q.Push(i)
	}
	var i int
	for ok := q.Pull(&i); ok; ok = q.Pull(&i) {
		fmt.Print(i, " ")
	}
	fmt.Println(); fmt.Println()
}

func TestPushGet (t *testing.T) {
	fmt.Println("TestPushGet")
	q := NewQueue[int]()
	for i := 0; i <= 9; i++ {
		q.Push(i)
	}
	for !q.IsEmpty() {
		fmt.Print(q.Pop(), " ")
	}
	fmt.Println(); fmt.Println()
}

func TestPushN (t *testing.T) {
	fmt.Println("TestPushN")
	q := NewQueue[int]()
	for i := 0; i <= 9; i++ {
		ok := q.PushN(i, i); M.Assert(ok, i, 100)
	}
	for e := q.Next(nil); e != nil; e = q.Next(e) {
		fmt.Print(e.Val(), " ")
	}
	fmt.Println(); fmt.Println()
}

func TestPushN0 (t *testing.T) {
	fmt.Println("TestPushN0")
	q := NewQueue[int]()
	for i := 0; i <= 9; i++ {
		ok := q.PushN(0, i); M.Assert(ok, i, 100)
	}
	for e := q.Next(nil); e != nil; e = q.Next(e) {
		fmt.Print(e.Val(), " ")
	}
	fmt.Println(); fmt.Println()
}

func TestPushN2 (t *testing.T) {
	fmt.Println("TestPushN2")
	a := [...]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	for n := 0; n <= 9; n++ {
		q := NewQueue[int]()
		for i := 0; i <= n; i++ {
			ok := q.PushN(a[i], i); M.Assert(ok, i, a[i], 100)
		}
		for e := q.Next(nil); e != nil; e = q.Next(e) {
			fmt.Print(e.Val(), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestPushNN (t *testing.T) {
	fmt.Println("TestPushNN")
	a := [...]int{0, 0, 2, 2, 4, 4, 6, 6, 8, 8}
	for n := 0; n <= 9; n++ {
		q := NewQueue[int]()
		for i := 0; i <= n; i++ {
			ok := q.PushN(a[i], i); M.Assert(ok, i, a[i], 100)
		}
		for e := q.Next(nil); e != nil; e = q.Next(e) {
			fmt.Print(e.Val(), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestPullPushN (t *testing.T) {
	fmt.Println("TestPullPushN")
	
	test := func (m, n int) {
		fmt.Printf("(%v - %v)\n", m, n)
		q := NewQueue[int]()
		for i := 0; i <= m; i++ {
			q.Push(i)
		}
		for e := q.Next(nil); e != nil; e = q.Next(e) {
			fmt.Print(e.Val(), " ")
		}
		fmt.Println()
		var i int
		q.PullN(n, &i)
		fmt.Println(i)
		for e := q.Next(nil); e != nil; e = q.Next(e) {
			fmt.Print(e.Val(), " ")
		}
		fmt.Println()
		q.PushN(n, i)
		for e := q.Next(nil); e != nil; e = q.Next(e) {
			fmt.Print(e.Val(), " ")
		}
		fmt.Println(); fmt.Println()
	}
	
	for n := 0; n <= 4; n++ {
		test(n, 0)
		test(n, n / 2)
		test(n, n)
	}
}
