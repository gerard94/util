package resources

import (
	
	F	"path/filepath"
	"fmt"
	"testing"
	/*
	"os"
	*/

)

func Test_inSrc (t *testing.T) {
	is := inSrc()
	fmt.Println("inSrc =", is)
	if F.Base(is) != "src" {t.Fail()}
}

func TestFindDir (t *testing.T) {
	fmt.Println("TestFindDir")
	fmt.Println(FindDir("util/json"))
	fmt.Println(FindDir("util/babel/api"))
	fmt.Println(FindDir("duniter/sandbox"))
	fmt.Println(FindDir("duniterClient"))
//	fmt.Println(FindDir("in"))
	fmt.Println(FindDir("util/babel"))
}
