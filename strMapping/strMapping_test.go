package strMapping

import (

	"testing"
	"fmt"

)

func TestDico (t *testing.T) {
	fmt.Println(NewStdLanguage().Language())
	l := NewLanguage("en")
	fmt.Println(l.Language())
	fmt.Println(l.Map("#util/babel:Ehard"))
	fmt.Println(l.Map("#duniter:NotANumber"))
	fmt.Println(l.Map("#duniterClient:FirstEntriesFluxPM"))
	l = NewLanguage("fr")
	fmt.Println(l.Language())
	fmt.Println(l.Map("#util/babel:Ehard"))
	fmt.Println(l.Map("#duniter:NotANumber"))
	fmt.Println(l.Map("#duniterClient:FirstEntriesFluxPM"))
	l = NewLanguage("es")
	fmt.Println(l.Language())
	fmt.Println(l.Map("#util/babel:Ehard"))
	fmt.Println(l.Map("#duniter:NotANumber"))
	fmt.Println(l.Map("#duniterClient:FirstEntriesFluxPM"))
}
