package books

import (
	
	E	"git.duniter.org/gerard94/util/graphQL/exec"
	F	"path/filepath"
	M	"git.duniter.org/gerard94/util/misc"
	R	"git.duniter.org/gerard94/util/resources"
		"fmt"
		"os"
		"testing"
	
)

const (
	
	booksDir = "util/graphQL/examples/books"
	answersDir = "answers"

)

var (
	
	booksPath = R.FindDir(booksDir)
	booksAnswersPath = F.Join(booksPath, answersDir)

)

func newF (name string) *os.File {
	f, err := os.Create(F.Join(booksAnswersPath, name))
	M.Assert(err == nil, err, 100)
	return f
}

func readDocument (exe *E.Executor, docName string) {
	exe.ReadDocument(F.Join(booksPath, docName + ".txt"))
}

func readVariableValues (exe *E.Executor, docName string) {
	exe.ReadVariableValues(F.Join(booksPath, docName + ".vars.txt"))
}

func executeDocument (exe *E.Executor, operation string) {
	f := newF(operation)
	defer f.Close()
	exe.ExecuteDocument(f, operation)
}

func fillDataBase () {
	exe := NewTypeSystem()
	readDocument(exe, "CreateBooks")
	readVariableValues(exe, "CreateBooks")
	executeDocument(exe, "CreateBooks")
}

func TestAuthors (t *testing.T) {
	fmt.Println(TestAuthors)
	fillDataBase()
	e := db.authors.Next(nil)
	for e != nil {
		fmt.Println(e.Val().(*nameSt).name)
		e = db.authors.Next(e)
	}
}

func TestAuthor (t *testing.T) {
	fillDataBase()
	name := "Card"
	n := 0
	e, _, _ := db.authors.SearchNext(&nameSt{name: name})
	for e != nil && prefix(name, e.Val().(*nameSt).name) {
		n++
		e = db.authors.Next(e)
	}
	fmt.Println(n)
}
