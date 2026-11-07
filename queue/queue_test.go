package queue

import (
	
	M	"git.duniter.org/gerard94/util/misc"
		"testing"
		"fmt"

)

func (q Queue[int]) printIntQueue () {
	for e := q.Next(nil); e != nil; e = q.Next(e) {
		fmt.Print(e.Val(), " ")
	}
}

func ExampleQueue_Push () {
	fmt.Println("TestPush")
	var q Queue[int]
	q.Init();
	for i := 0; i <= 9; i++ {
		q.Push(i)
	}
	fmt.Println("length =", q.Length())
	q.printIntQueue()
	fmt.Println(); fmt.Println()
	//Output:
	//TestPush
	//length = 10
	//0 1 2 3 4 5 6 7 8 9 
	//
}

func TestPushPull (t *testing.T) {
	fmt.Println("TestPushPull")
	var q Queue[int]
	q.Init();
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
	var q Queue[int]
	q.Init();
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
	var q Queue[int]
	q.Init();
	for i := 0; i <= 9; i++ {
		ok := q.PushN(i, i); M.Assert(ok, i, 100)
	}
	q.printIntQueue()
	fmt.Println(); fmt.Println()
}

func TestPushN0 (t *testing.T) {
	fmt.Println("TestPushN0")
	var q Queue[int]
	q.Init();
	for i := 0; i <= 9; i++ {
		ok := q.PushN(0, i); M.Assert(ok, i, 100)
	}
	q.printIntQueue()
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
		q.printIntQueue()
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
		q.printIntQueue()
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
		q.printIntQueue()
		fmt.Println()
		var i int
		q.PullN(n, &i)
		fmt.Println(i)
		q.printIntQueue()
		fmt.Println()
		q.PushN(n, i)
		q.printIntQueue()
		fmt.Println(); fmt.Println()
	}
	
	for n := 0; n <= 4; n++ {
		test(n, 0)
		test(n, n / 2)
		test(n, n)
	}
}

func TestConcat (t *testing.T) {
	fmt.Println("TestConcat")
	var q1, q2 Queue[int]
	q1.Init(); q2.Init()
	var (
		v1 = [...]int{2,6,7,4,5}
		v2 = [...]int{7,3,0,8,6}
	)
	for _, v := range v1 {
		q1.Push(v)
	}
	for _, v := range v2 {
		q2.Push(v)
	}
	fmt.Print("q1 = ")
	q1.printIntQueue()
	fmt.Println()
	fmt.Print("q2 = ")
	q2.printIntQueue()
	q1.Concat(q2)
	fmt.Println()
	fmt.Print("q1.Concat(q2) = ")
	q1.printIntQueue()
	fmt.Println(); fmt.Println()
}
