package graphQL
	
import (
	
	M	"git.duniter.org/gerard94/util/misc"
	T	"testing"
		"fmt"
		"os"

)
	
const (

	path1 = "/home/gerard/goPerso/rsrc/duniter/TypeSystem.txt"
	path2 = "/home/gerard/goPerso/rsrc/util/graphQL/examples/books/TypeSystem.txt"

)

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
	doc, _ := ReadGraphQL(path2); M.Assert(doc != nil, 100);
	doc.WriteFlat(os.Stdout)
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout)
	doc.Write(os.Stdout)
	fmt.Fprintln(os.Stdout)
}
