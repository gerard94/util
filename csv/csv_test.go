package csv

import (
	
	R	"github.com/gerard94/util/resources"
		"testing"
		"fmt"

)

const (
	
	test0 = "util/csv/examples/test0.csv"
	test1 = "util/csv/examples/test1.csv"
	test2 = "util/csv/examples/test2.csv"
	test3 = "util/csv/examples/test3.csv"

)

var (
	
	rsrcTest0 = R.FindDir(test0)
	rsrcTest1 = R.FindDir(test1)
	rsrcTest2 = R.FindDir(test2)
	rsrcTest3 = R.FindDir(test3)

)

func test (t *testing.T, path string) {
	c1 := CompileFile(path)
	fmt.Println("Content of", path, "is")
	if c1 == nil {
		fmt.Println("Empty")
		fmt.Println()
		return
	}
	for _, l := range c1 {
		for _, c := range l {
			fmt.Print(c, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func Test0 (t *testing.T) {
	test(t, rsrcTest0)
}

func Test1 (t *testing.T) {
	test(t, rsrcTest1)
}

func Test2 (t *testing.T) {
	test(t, rsrcTest2)
}

func Test3 (t *testing.T) {
	test(t, rsrcTest3)
}
