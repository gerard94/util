package json

import (
	
	F	"path/filepath"
	M	"git.duniter.org/gerard94/util/misc"
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
	v1 = &Int{42}
	v2 = &String{"OK"}
	j = &Array{Values{v0, v1, v2}}
	j.Write(os.Stdout)
	fmt.Fprint(os.Stdout, "\n")
}

func TestJ2 (t *testing.T) {
	fmt.Fprintln(os.Stdout, "\nTestJ2")
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
	fmt.Fprintln(os.Stdout, "\nTestJ3")
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
	fmt.Fprintln(os.Stdout, "\nTestJ4")
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
	fmt.Fprintln(os.Stdout, "\nTestJ5")
	path := F.Join(wd, "examples/result.json")
	j := ReadFile(path)
	if j == nil {
		t.FailNow()
	}
	j.Write(os.Stdout)
	fmt.Fprintln(os.Stdout)
	j.WriteFlat(os.Stdout)
	fmt.Fprintln(os.Stdout)
}

func TestJ6 (t *testing.T) {
	fmt.Fprintln(os.Stdout, "\nTestJ6")
	
	s1 := `
		{
			"un":
				"un",
			"deux":
				2,
			"trois":
				3.1,
			"quatre":
				true,
			"cinq":
				null,
			"six":
				{
					"sept":
						7
				},
			"huit":
				[
					"neuf",
					10,
					11.11,
					false,
					null,
					{
						"dix":
							"dix"
					}
				]
		}
	`
	s2 := `
	[
		2,
		4,
		6,
		8,
		{
			"target":
				"result"
		}
	]
	`
	
	j := ReadString(s1); M.Assert(j != nil)
	vs, err := Find(j, "un", StringKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == StringKind && vs.String == "un")
	vs, err = Find(j, "deux", IntKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == IntKind && vs.Int == 2)
	vs, err = Find(j, "trois", FloatKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == FloatKind && vs.Float == 3.1)
	vs, err = Find(j, "quatre", BooleanKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == BooleanKind && vs.Boolean)
	vs, err = Find(j, "cinq", NullKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == NullKind)
	vs, err = Find(j, "six", JsonKind, ObjectKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == JsonKind && vs.Json.JKind() == ObjectKind)
	vs, err = Find(j, "six.sept", IntKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == IntKind && vs.Int == 7)
	vs, err = Find(j, "huit", JsonKind, ArrayKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == JsonKind && vs.Json.JKind() == ArrayKind)
	vs, err = Find(j, "huit[0]", StringKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == StringKind && vs.String == "neuf")
	vs, err = Find(j, "huit[3]", BooleanKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == BooleanKind && !vs.Boolean)
	vs, err = Find(j, "huit[5]", JsonKind, ObjectKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == JsonKind && vs.Json.JKind() == ObjectKind)
	vs, err = Find(j, "huit[5].dix", StringKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == StringKind && vs.String == "dix")
	vs, err = Find(j, "huit[6]")
	M.Assert(err != nil); fmt.Println(err)
	vs, err = Find(j, "sept")
	M.Assert(err != nil); fmt.Println(err)
	
	j = ReadString(s2); M.Assert(j != nil)
	vs, err = Find(j, "[2]", IntKind)
	M.Assert(vs.Kind == IntKind && vs.Int == 6)
	M.Assert(err == nil, err)
	vs, err = Find(j, "[4]", JsonKind, ObjectKind)
	M.Assert(vs.Kind == JsonKind && vs.Json.JKind() == ObjectKind)
	j = vs.Json
	vs, err = Find(j, "target", StringKind)
	M.Assert(err == nil, err)
	M.Assert(vs.Kind == StringKind && vs.String == "result")
	vs, err = Find(j, "result")
	M.Assert(err != nil); fmt.Println(err)
}
