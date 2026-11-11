package main

import (
	
	B	"github.com/gerard94/util/graphQL/examples/books"
	E	"github.com/gerard94/util/graphQL/exec"
	F	"path/filepath"
	M	"github.com/gerard94/util/misc"
	R	"github.com/gerard94/util/resources"
		"fmt"
		"os"
		"sync"

)

const (
	
	booksDir = "util/graphQL/examples/books"
	answersDir = "answers"
	typeSystemName = "TypeSystem.txt"

)

type (
	
	inT struct {
	}

	outT struct {
		pos int
	}
	
	nullWriter struct {
	}

)

var (
	
	in *inT
	out *outT
	
	booksPath = R.FindDir(booksDir)
	booksAnswersPath = F.Join(booksPath, answersDir)
	typeSystemPath = F.Join(booksPath, typeSystemName)
	
	nullW nullWriter

)

func (o *outT) Write (p []byte) (n int, err error) {
	const lim = 80
	w := os.Stdout
	rs := []rune(string(p))
	n = 0
	for _, r := range rs {
		s := string(r)
		var m int
		m, err = fmt.Fprint(w, s)
		n += m
		if err != nil {
			return
		}
		o.pos++
		if o.pos >= lim {
			fmt.Fprintln(w)
			o.pos = 0
		}
		if r == '\n' {
			o.pos = 0
		}
	}
	return
}

func (o *outT) Str (s string) {
	fmt.Fprint(o, s)
}

func (o *outT) Ln () {
	fmt.Fprintln(o)
}

func newF (name string) *os.File {
	f, err := os.Create(F.Join(booksAnswersPath, name))
	M.Assert(err == nil, err, 100)
	return f
}

func readDocument (exe *E.Executor, docName string) bool {
	return exe.ReadDocument(F.Join(booksPath, docName + ".txt"))
}

func readVariableValues (exe *E.Executor, docName string) {
	exe.ReadVariableValues(F.Join(booksPath, docName + ".vars.txt"))
}

func executeDocument (exe *E.Executor, operation string) {
	f := newF(operation)
	defer f.Close()
	exe.ExecuteDocument(f, operation)
}

func execOneSeq (exe *E.Executor, v, d string) {
	exe = exe.CopyExecutorTS ()
	if v != "" {
		readVariableValues(exe, v)
	}
	readDocument(exe, d)
	for _, c := range exe.OperationNames() {
		fmt.Fprintln(os.Stdout, c)
		func () {
			executeDocument(exe, c)
		}()
	}
}

func execOnePar (wg *sync.WaitGroup, exe *E.Executor, v, d string) {
	exe = exe.CopyExecutorTS ()
	if v != "" {
		readVariableValues(exe, v)
	}
	readDocument(exe, d)
	for _, c := range exe.OperationNames() {
		fmt.Fprintln(os.Stdout, c)
		go func (wg *sync.WaitGroup, exe *E.Executor, c string) {
			defer wg.Done()
			executeDocument(exe, c)
		}(wg, exe, c)
	}
}

func countPars (exe *E.Executor, ds ...string) int {
	n := 0
	for _, d := range ds {
		if readDocument(exe, d) {
			n += len(exe.OperationNames())
		}
	}
	return n
}

func main () {
	exe := B.NewTypeSystem()
	execOneSeq(exe, "", "AuthorCreated")
	execOneSeq(exe, "", "BookCreated")
	execOneSeq(exe, "", "PublisherCreated")
	execOneSeq(exe, "CreateBooks", "CreateBooks")
	var wg sync.WaitGroup
	wg.Add(countPars(exe, "AllSimpleQueries", "AuthorsOfBooksOf", "AllSimpleFalseQueries", "AllSimpleFalseQueries2"))
	execOnePar(&wg, exe, "", "AllSimpleQueries")
	execOnePar(&wg, exe, "", "AllSimpleFalseQueries")
	execOnePar(&wg, exe, "AllSimpleFalseQueries2", "AllSimpleFalseQueries2")
	execOnePar(&wg, exe, "", "AuthorsOfBooksOf")
	wg.Wait()
	execOneSeq(exe, "", "DeleteBook")
	execOneSeq(exe, "DeleteBooks", "DeleteBooks")
	execOneSeq(exe, "", "BooksIds")
}

func init () {
	os.MkdirAll(booksAnswersPath, 0777)
}
