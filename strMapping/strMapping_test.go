package strMapping

import (

	"testing"
	"fmt"

)

func TestDico (t *testing.T) {
	fmt.Println(NewStdLanguage().Language())
	l := NewLanguage("en")
	fmt.Println(l.Language())
	//fmt.Println(l.Map("#in:WrongSymbAr"))
	//fmt.Println(l.Map("#in:Bytes"))
	fmt.Println(l.Map("#util/babel:Ehard"))
	l = NewLanguage("fr")
	fmt.Println(l.Language())
	//fmt.Println(l.Map("#in:WrongSymbAr"))
	//fmt.Println(l.Map("#in:Bytes"))
	fmt.Println(l.Map("#util/babel:Ehard"))
}
