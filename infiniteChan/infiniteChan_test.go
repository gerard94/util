package infiniteChan

import (
	
	"testing"
	"fmt"

)

var (
	
	infC = NewInfChan[int]()

)

func TestSend (t *testing.T) {
	fmt.Println("TestSend")
	a := [...]int{8, 6, 3, 4, 7, 9, 1, 2, 3}
	for _, v := range a {
		fmt.Println(v)
		infC.Send(v)
	}
}

func TestReceive (t *testing.T) {
	fmt.Println("TestReceive")
	val, ok := infC.Receive()
	for ok {
		fmt.Println(val)
		val, ok = infC.Receive()
	}
}

