package main

import (
	
//	F	"path/filepath"
	R	"github.com/gerard94/util/resources"
	V	"github.com/gerard94/util/graphQL/examples/validation"
		"fmt"
		"os"
	_	"github.com/gerard94/util/graphQL/static"

)

const (
	
	validationDir = "util/graphQL/examples/validation"

)

type (
	
	inT struct {
	}

	outT struct {
		pos int
	}

)

var (
	
	in *inT
	out *outT
	
	validationPath = R.FindDir(validationDir)

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

func main () {
	exe := V.NewTypeSystem()
	exe.ReadDocument(validationPath + "/02 - Invalid - Executable Definitions.txt")
	exe.ReadDocument(validationPath + "/03 - Valid - Operation Name Uniqueness.txt")
	exe.ReadDocument(validationPath + "/04 - Invalid - Operation Name Uniqueness.txt")
	exe.ReadDocument(validationPath + "/05 - Invalid - Operation Name Uniqueness.txt")
	exe.ReadDocument(validationPath + "/06 - Valid - Lone Anonymous Operation.txt")
	exe.ReadDocument(validationPath + "/07 - Invalid - Lone Anonymous Operation.txt")
	exe.ReadDocument(validationPath + "/08 - Valid - Subscription Single root field.txt")
	exe.ReadDocument(validationPath + "/09 - Valid - Subscription Single root field.txt")
	exe.ReadDocument(validationPath + "/10 - Invalid - Subscription Single root field.txt")
	exe.ReadDocument(validationPath + "/11 - Invalid - Subscription Single root field.txt")
	exe.ReadDocument(validationPath + "/12 - Invalid - Subscription Single root field.txt")
	exe.ReadDocument(validationPath + "/13 - Invalid - Scoped Target Fields.txt")
	exe.ReadDocument(validationPath + "/14 - Valid - Scoped Target Fields.txt")
	exe.ReadDocument(validationPath + "/15 - Invalid - Scoped Target Fields.txt")
	exe.ReadDocument(validationPath + "/16 - Valid - Scoped Target Fields.txt")
	exe.ReadDocument(validationPath + "/17 - Invalid - Scoped Target Fields.txt")
	exe.ReadDocument(validationPath + "/18 - Valid - Field Selection Merging.txt")
	exe.ReadDocument(validationPath + "/19 - Invalid - Field Selection Merging.txt")
	exe.ReadDocument(validationPath + "/20 - Valid - Field Selection Merging.txt")
	exe.ReadDocument(validationPath + "/21 - Invalid - Field Selection Merging.txt")
	exe.ReadDocument(validationPath + "/22 - Valid - Field Selection Merging.txt")
	exe.ReadDocument(validationPath + "/23 - Invalid - Field Selection Merging.txt")
	exe.ReadDocument(validationPath + "/24 - Valid - Leaf Field Selections.txt")
	exe.ReadDocument(validationPath + "/25 - Invalid - Leaf Field Selections.txt")
	exe.ReadDocument(validationPath + "/26 - Invalid - Leaf Field Selections.txt")
	exe.ReadDocument(validationPath + "/27 - Valid - Arguments.txt")
	exe.ReadDocument(validationPath + "/28 - Invalid - Argument Names.txt")
	exe.ReadDocument(validationPath + "/29 - Invalid - Argument Names.txt")
	exe.ReadDocument(validationPath + "/30 - Valid - Argument Names.txt")
	exe.ReadDocument(validationPath + "/31 - Invalid - Argument Uniqueness.txt")
	exe.ReadDocument(validationPath + "/32 - Valid - Required Arguments.txt")
	exe.ReadDocument(validationPath + "/33 - Valid - Required Arguments.txt")
	exe.ReadDocument(validationPath + "/34 - Invalid - Required Arguments.txt")
	exe.ReadDocument(validationPath + "/35 - Invalid - Required Arguments.txt")
	exe.ReadDocument(validationPath + "/36 - Valid - Fragment Name Uniqueness.txt")
	exe.ReadDocument(validationPath + "/37 - Invalid - Fragment Name Uniqueness.txt")
	exe.ReadDocument(validationPath + "/38 - Valid - Fragment Spread Type Existence.txt")
	exe.ReadDocument(validationPath + "/39 - Invalid - Fragment Spread Type Existence.txt")
	exe.ReadDocument(validationPath + "/40 - Valid - Fragments On Composite Types.txt")
	exe.ReadDocument(validationPath + "/41 - Invalid - Fragments On Composite Types.txt")
	exe.ReadDocument(validationPath + "/42 - Invalid - Fragments Must Be Used.txt")
	exe.ReadDocument(validationPath + "/43 - Invalid - Fragment spread target defined.txt")
	exe.ReadDocument(validationPath + "/44 - Invalid - Fragment spreads must not form cycles.txt")
	exe.ReadDocument(validationPath + "/45 - Invalid - Fragment spreads must not form cycles.txt")
	exe.ReadDocument(validationPath + "/46 - Valid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/47 - Invalid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/48 - Valid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/49 - Valid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/50 - Valid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/51 - Invalid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/52 - Valid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/53 - Invalid - Fragment spread is possible.txt")
	exe.ReadDocument(validationPath + "/54 - Valid - Values of Correct Type.txt")
	exe.ReadDocument(validationPath + "/55 - Invalid - Values of Correct Type.txt")
	exe.ReadDocument(validationPath + "/56 - Valid - Values of Correct Type.txt")
	exe.ReadDocument(validationPath + "/57 - Valid - Input Object Field Names.txt")
	exe.ReadDocument(validationPath + "/58 - Invalid - Input Object Field Names.txt")
	exe.ReadDocument(validationPath + "/59 - Invalid - Input Object Field Uniqueness.txt")
	exe.ReadDocument(validationPath + "/60 - Invalid - Input Object Required Fields.txt")
	exe.ReadDocument(validationPath + "/61 - Invalid - Input Object Required Fields.txt")
	exe.ReadDocument(validationPath + "/62 - Invalid - Directives Are In Valid Locations.txt")
	exe.ReadDocument(validationPath + "/63 - Valid - Directives Are In Valid Locations.txt")
	exe.ReadDocument(validationPath + "/64 - Invalid - Directives Are Unique Per Location.txt")
	exe.ReadDocument(validationPath + "/65 - Valid - Directives Are Unique Per Location.txt")
	exe.ReadDocument(validationPath + "/66 - Invalid - Variable Uniqueness.txt")
	exe.ReadDocument(validationPath + "/67 - Valid - Variable Uniqueness.txt")
	exe.ReadDocument(validationPath + "/68 - Valid - Variables Are Input Types.txt")
	exe.ReadDocument(validationPath + "/69 - Invalid - Variables Are Input Types.txt")
	exe.ReadDocument(validationPath + "/70 - Valid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/71 - Invalid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/72 - Valid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/73 - Invalid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/74 - Invalid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/75 - Valid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/76 - Invalid - All Variable Uses Defined.txt")
	exe.ReadDocument(validationPath + "/77 - Invalid - All Variables Used.txt")
	exe.ReadDocument(validationPath + "/78 - Valid - All Variables Used.txt")
	exe.ReadDocument(validationPath + "/79 - Invalid - All Variables Used.txt")
	exe.ReadDocument(validationPath + "/80 - Invalid - All Variables Used.txt")
	exe.ReadDocument(validationPath + "/81 - Invalid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/82 - Invalid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/83 - Invalid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/84 - Valid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/85 - Invalid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/86 - Valid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/87 - Valid - All Variable Usages are Allowed.txt")
	exe.ReadDocument(validationPath + "/88 - Valid - Introspection.txt")
	exe.ReadDocument(validationPath + "/89 - Valid - Introspection.txt")
	exe.ReadDocument(validationPath + "/90 - Valid - Introspection.txt")
	exe.ReadDocument(validationPath + "/91 - Invalid - Introspection.txt")
	exe.ReadDocument(validationPath + "/92 - Invalid - Introspection.txt")
}
