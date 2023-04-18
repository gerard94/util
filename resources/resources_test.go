package resources

import (
	
	"fmt"
	"testing"

)

func Test_inSrc (t *testing.T) {
	root, src := inSrc()
	fmt.Println("root =", root, "src =", src)
}

func TestFindDir (t *testing.T) {
	fmt.Println("TestFindDir")
	fmt.Println(FindDir("util/json"))
	fmt.Println(FindDir("util/babel/api"))
	fmt.Println(FindDir("duniter/sandbox"))
	fmt.Println(FindDir("duniterClient"))
	fmt.Println(FindDir("util/babel"))
}
