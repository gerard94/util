package stack

import (
	
		"testing"
		"fmt"

)

func TestPush (t *testing.T) {
	fmt.Println("TestPush")
	s := NewStack[int]()
	for i := 0; i <= 100; i++ {
		s.Push(i)
	}
	it := s.NewIterator()
	var i int
	for b := it.Next(true, &i); b; b = it.Next(false, &i) {
		fmt.Print(i, " ")
	}
	fmt.Println(); fmt.Println()
}

func TestPushPull (t *testing.T) {
	fmt.Println("TestPushPull")
	s := NewStack[int]()
	for i := 0; i <= 100; i++ {
		s.Push(i)
	}
	var i int
	for ok := s.Pull(&i); ok; ok = s.Pull(&i) {
		fmt.Print(i, " ")
	}
	fmt.Println(); fmt.Println()
}
