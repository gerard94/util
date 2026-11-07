package misc

import (
	
	"testing"
	"fmt"
	"time"

)

func TestVersion (t *testing.T) {
	fmt.Println("Version =", Version())
}

func Test2 (t *testing.T) {
	const name = "Test InstantCreate.txt"
	f, err := InstantCreate(name); Assert(err == nil, 100)
	fmt.Fprint(f, "Hello")
	err = InstantClose(f); Assert(err == nil, 101)
}

func TestOdd (t *testing.T) {
	Want(!Odd(0), t)
	Want(Odd(1), t)
	Want(!Odd(1000), t)
	Want(Odd(1001), t)
}

func TestTimeStamp (t *testing.T) {
	ts1 := int64(time.Now().Unix())
	fmt.Println("Time1 =", ts1)
	year, month, day, hour, minute, second := Time2Timestamp(ts1)
	fmt.Printf("%v/%v/%v-%v:%v:%v\n", day, month, year, hour, minute, second)
	ts2 := Timestamp2Time(year, month, day, hour, minute, second)
	fmt.Println("Time2 =", ts2)
	Assert(ts1 == ts2)
}
