package graphQL
	
import (
	
	M	"git.duniter.org/gerard94/util/misc"
	R	"git.duniter.org/gerard94/util/resources"
	T	"testing"
		"fmt"
		"os"

)
	
const (

	path1 = "/home/gerard/goPerso/rsrc/duniter/TypeSystem.txt"
	path2 = "util/graphQL/examples/books/TypeSystem.txt"

)

var srcPath2 = R.FindDir(path2)

func TestWrite1 (t *T.T) {
	doc, _ := ReadGraphQL(path1); M.Assert(doc != nil, 100);
	doc.WriteFlat(os.Stdout)
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout)
	doc.Write(os.Stdout)
	fmt.Fprintln(os.Stdout)
}

func TestWrite2 (t *T.T) {
	fmt.Fprintln(os.Stdout)
	doc, _ := ReadGraphQL(srcPath2); M.Assert(doc != nil, 100);
	doc.WriteFlat(os.Stdout)
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout)
	doc.Write(os.Stdout)
	fmt.Fprintln(os.Stdout)
}
