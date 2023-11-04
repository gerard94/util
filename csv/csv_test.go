package csv

import (
	
	R	"git.duniter.org/gerard94/util/resources"
		"testing"
		"fmt"

)

const (
	
	test1 = "util/csv/examples/test1.csv"
	test2 = "util/csv/examples/test2.csv"

)

var (
	
	rsrcTest1 = R.FindDir(test1)
	rsrcTest2 = R.FindDir(test2)

)

func Test1 (t *testing.T) {
	c1 := CompileFile(rsrcTest1)
	if c1 == nil {
		fmt.Println("c1 void")
	}
	fmt.Println("c1 =")
	for _, l := range c1 {
		for _, c := range l {
			fmt.Print(c, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func Test2 (t *testing.T) {
	c2 := CompileFile(rsrcTest2)
	if c2 == nil {
		fmt.Println("c2 void")
	}
	fmt.Println("c2 =")
	for _, l := range c2 {
		for _, c := range l {
			fmt.Print(c, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
