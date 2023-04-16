package json

import (
	
	F	"path/filepath"
		"testing"
		"os"
		"fmt"

)

func TestJ1 (t *testing.T) {
	var (
		
		v0, v1, v2 Value
		j Json
	
	)
	fmt.Fprintln(os.Stdout, "TestJ1")
	v0 = new(Null)
	v1 = &Integer{42}
	v2 = &String{"OK"}
	j = &Array{Values{v0, v1, v2}}
	j.Write(os.Stdout)
	fmt.Fprint(os.Stdout, "\n")
}

func TestJ2 (t *testing.T) {
	fmt.Fprintln(os.Stdout, "TestJ2")
	mk := NewMaker()
	mk.StartObject()
	mk.PushBoolean(true)
	mk.BuildField("truth")
	mk.PushBoolean(false)
	mk.BuildField("falseness")
	mk.BuildObject()
	j := mk.GetJson()
	j.Write(os.Stdout)
	fmt.Fprint(os.Stdout, "\n")
}

func TestJ3 (t *testing.T) {
	fmt.Fprintln(os.Stdout, "TestJ3")
	mk := NewMaker()
	mk.StartArray()
	mk.PushNull()
	mk.PushInteger(42)
	mk.PushString("OK\"OK\"")
	mk.StartObject()
	mk.PushBoolean(true)
	mk.BuildField("truth")
	mk.PushBoolean(false)
	mk.BuildField("falseness")
	mk.BuildObject()
	mk.PushFloat(3.14159265)
	mk.BuildArray()
	j := mk.GetJson()
	j.Write(os.Stdout)
	fmt.Fprint(os.Stdout, "\n")
}

func TestJ4 (t *testing.T) {
	fmt.Fprintln(os.Stdout, "TestJ4")
	mk := NewMaker()
	mk.StartObject()
	mk.PushBoolean(true)
	mk.BuildField("truth")
	mk.StartArray()
	mk.PushNull()
	mk.PushInteger(42)
	mk.PushString("OK\"OK\"")
	mk.PushFloat(3.14159265)
	mk.BuildArray()
	mk.BuildField("list")
	mk.PushBoolean(false)
	mk.BuildField("falseness")
	mk.BuildObject()
	j := mk.GetJson()
	j.Write(os.Stdout)
	fmt.Fprint(os.Stdout, "\n")
}

func TestJ5 (t *testing.T) {
	fmt.Fprintln(os.Stdout, "TestJ5")
	path := F.Join(wd, "examples/result.json")
	j := ReadFile(path)
	if j == nil {
		t.FailNow()
	}
	j.Write(os.Stdout)
	fmt.Fprintln(os.Stdout)
	j.WriteFlat(os.Stdout)
}
