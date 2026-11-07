/* 
Util: Utility tools.

Copyright (C) 2017…2019 GérardMeunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License  for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
*/

package json

import (
	
	C	"git.duniter.org/gerard94/util/babel/compil"
	F	"path/filepath"
	J	"encoding/json"
	M	"git.duniter.org/gerard94/util/misc"
	R	"reflect"
	RS	"git.duniter.org/gerard94/util/resources"
	SC	"strconv"
		"bufio"
		"bytes"
		"fmt"
		"io"
		"os"
		"strings"
		"unicode"
		"strconv"
		"errors"

)

const (
	
	double_quote = '\x22'
	reverse_solidus = '\x5C'
	solidus = '\x2F'
	backspace = '\x08'
	form_feed = '\x0C'
	line_feed = '\x0A'
	carriage_return = '\x0D'
	horizontal_tab = '\x09'
	
	jsonDir = "util/json" // Directory of JSON resources files
	compName = "json.tbl" // Json compiler

)

// Json Kind
const (
	
	ObjectKind = iota
	ArrayKind

)

// Value Kind
const (
	
	JsonKind = iota
	StringKind
	IntKind
	FloatKind
	BooleanKind
	NullKind

)

type (
	
	// For compiler internalizing
	directory struct { // C.Directory
		r *bufio.Reader
	}
	
	compilationer struct {
		r []rune
		size,
		pos int
		json Json
	}

)

type (
	
	// Outputs of ReadString, ReadText && ReadFile
	
	// A Json may be an 'Object' or an 'Array'
	Json interface {
		
		// 'ObjectKind' or 'ArrayKind'
		JKind () int
		
		// Panic if not 'Object'
		GetObject () *Object
		
		// Panic if not 'Array'
		GetArray () *Array
		
		// 'GetString' returns a formatted string for reading
		GetString () string
		
		// 'GetFlatString' returns a compact string
		GetFlatString () string
		
		// 'Write' writes 'GetString()' to 'w'
		Write (w io.Writer)
		
		// 'WriteFlat' writes 'GetFlatString' to 'w'
		WriteFlat (w io.Writer)
	
	}
	
	// Json Object
	Object struct {
		Fields Fields
	}
	
	// Fields of Object
	Fields []*Field
	
	// Field of Object
	Field struct {
		Name string
		Value Value
	}
	
	// Json Array
	Array struct {
		Elements Values
	}
	
	Values []Value
	
	// Any 'Value' in 'Object' or 'Array'
	Value interface {
		
		// Value Kind
		VKind () int
	
	}
	
	// Json Value
	JsonVal struct {
		Json Json
	}
	
	// String Value
	String struct {
		S string
	}
	
	// Int Value
	Int struct {
		N int64
	}
	
	// Float Value
	Float struct {
		F float64
	}
	
	// Boolean Value
	Boolean struct {
		Boolean bool
	}
	
	// Null Value
	Null struct {
	}
	
	// Value returned by 'Find'
	ValueStruct struct {
		// Kind of the value (VKind)
		Kind int
		// Only one of the following fields is filled, according to 'Kind'
		Json Json
		String string
		Int int64
		Float float64
		Boolean bool
	}

)

// Stack used for building Json objects
type (
	
	stacker interface {
		next () stacker
		setNext (s stacker)
	}
	
	stack struct {
		nxt stacker
	}
	
	// Beginning of an 'Object' in a 'stack'
	startObj struct {
		stack
	}
	
	// Beginning of an 'Array' in a 'stack'
	startArr struct {
		stack
	}
	
	// 'Value' in a 'stack'
	stValue struct {
		stack
		val Value
	}
	
	// 'Field' in a 'Stack'
	stField struct {
		stack
		name string
		value Value
	}
	
	// A 'Maker' builds 'Json' by stack manipulation
	Maker struct {
		stk stacker
	}

)

var (
	
	// Directory of Json resources
	wd = RS.FindDir(jsonDir)
	// File of Json Compiler
	compPath = F.Join(wd, compName)
	// Compiler of json texts
	comp *C.Compiler

)

// 'Json' kind of an 'Object'
func (*Object) JKind () int {
	return ObjectKind
}

// 'Json' kind of an 'Array'
func (*Array) JKind () int {
	return ArrayKind
}

// 'Value' kind of a 'Json'
func (*JsonVal) VKind () int {
	return JsonKind
}

// 'Value' kind of a 'String'
func (*String) VKind () int {
	return StringKind
}

// 'Value' kind of an 'Int'
func (*Int) VKind () int {
	return IntKind
}

// 'Value' kind of a 'Float'
func (*Float) VKind () int {
	return FloatKind
}

// 'Value' kind of a 'Boolean'
func (*Boolean) VKind () int {
	return BooleanKind
}

// 'Value' kind of a 'Null'
func (*Null) VKind () int {
	return NullKind
}

// Get an 'Object' out of a 'Json'
func (o *Object) GetObject () *Object {
	return o
}

// Panic if the 'Json' isn't an 'Array'
func (o *Object) GetArray () *Array {
	M.Halt(100)
	return nil
}

// Panic if the 'Json' isn't an 'Object'
func (a *Array) GetObject () *Object {
	M.Halt(100)
	return nil
}

// Get an 'Array' out of a 'Json'
func (a *Array) GetArray () *Array {
	return a
}

// 'getField' returns the 'Field' in 'o' whose 'Name' is 'name', or nil if the 'Field' doesn't exist
func getField (o *Object, name string) *Field {
	if o != nil && o.Fields != nil {
		for _, f := range o.Fields {
			if f.Name == name {
				return f
			}
		}
	}
	return nil
}

// 'GetValue' returns the 'Value' in 'o' whose 'Name' is 'name' and true if the search was successful
func GetValue (o *Object, name string) (Value, bool) {
	f := getField(o, name)
	if f == nil {
		return nil, false
	}
	return f.Value, true
}

// SetValue puts 'value' into the 'Field' of 'o' whose 'Name' is 'name'; if the 'Field' doesn't already exist, it is created
func SetValue (o *Object, name string, value Value) {
	f := getField(o, name)
	if f == nil {
		o.Fields = append(o.Fields, &Field{Name: name, Value: value})
		return
	}
	f.Value = value
}

// 'Len' returns the number of elements in 'a', or 0 if 'a' doesn't exist
func (a *Array) Len () int {
	if a.Elements != nil {
		return len(a.Elements)
	}
	return 0
}

// 'GetElem' returns the element of rank 'num' in 'a', and true if this element exists, otherwise the 'Value' returned is nil
func GetElem (a *Array, num int) (Value, bool) {
	if a != nil && a.Elements != nil && num >= 0 && num < len(a.Elements) {
		return a.Elements[num], true
	}
	return nil, false
}

// 'SetElem' puts 'value' into the element of rank 'num' in 'a', and returns true if such an element exists, otherwise the operation is unsuccessful
func SetElem (a *Array, num int, value Value) bool {
	if a != nil && a.Elements != nil && num >= 0 && num < len(a.Elements) {
		a.Elements[num] = value
		return true
	}
	return false
}

// 'GetJson' returns the 'Json' contained in 'v', and true if 'v' contains actually a 'Json', otherwise the 'Json' returned is nil
func GetJson (v Value) (Json, bool) {
	if v == nil {
		return nil, false
	}
	j, ok := v.(*JsonVal)
	if !ok || j == nil {
		return nil, false
	}
	return j.Json, true
}

// 'SetJson' returns a 'Value' containing 'j'
func SetJson (j Json) Value {
	return &JsonVal{Json: j}
}

// 'GetObject' returns the 'Object' in 'v', or nil if 'v' doesn't contain an 'Object'
func GetObject (v Value) *Object {
	if v == nil {
		return nil;
	}
	j, ok := v.(*JsonVal)
	if !ok || j == nil || j.Json == nil {
		return nil
	}
	o, ok := j.Json.(*Object)
	if !ok {
		return nil
	}
	return o
}

// 'GetArray' returns the 'Array' in 'v', or nil if 'v' doesn't contain an 'Array'
func GetArray (v Value) *Array {
	if v == nil {
		return nil;
	}
	j, ok := v.(*JsonVal)
	if !ok || j == nil || j.Json == nil {
		return nil
	}
	a, ok := j.Json.(*Array)
	if !ok {
		return nil
	}
	return a
}

// 'GetString' returns the string in 'v' and true; if 'v' doesn't contain a string, "" and false are returned
func GetString (v Value) (string, bool) {
	if v == nil {
		return "", false
	}
	s, ok := v.(*String)
	if !ok || s == nil {
		return "", false
	}
	return s.S, true
}

// 'SetString' returns a 'Value' containing 's'
func SetString (s string) Value {
	return &String{S: s}
}

// 'GetInt' returns the int64 in 'v' and true; if 'v' doesn't contain an int64, 0 and false are returned
func GetInt (v Value) (int64, bool) {
	if v == nil {
		return 0, false
	}
	i, ok := v.(*Int)
	if !ok || i == nil {
		return 0, false
	}
	return i.N, true
}

// 'SetInt' returns a 'Value' containing 'n'
func SetInt (n int64) Value {
	return &Int{N: n}
}

// 'GetFloat' returns the float64 in 'v' and true; if 'v' doesn't contain a float64, 0 and false are returned
func GetFloat (v Value) (float64, bool) {
	if v == nil {
		return 0, false
	}
	ff, ok := v.(*Float)
	if !ok || ff == nil {
		return 0, false
	}
	return ff.F, true
}

// 'SetFloat' returns a 'Value' containing 'f'
func SetFloat (f float64) Value {
	return &Float{F: f}
}

// 'GetBool' returns the bool in 'v' and true; if 'v' doesn't contain a GetBool, false and false are returned
func GetBool (v Value) (bool, bool) {
	if v == nil {
		return false, false
	}
	b, ok := v.(*Boolean)
	if !ok || b == nil {
		return false, false
	}
	return b.Boolean, true
}

// 'SetBool' returns a 'Value' containing 'b'
func SetBool (b bool) Value {
	return &Boolean{Boolean: b}
}

// 'GetNull' returns true if 'v' is a 'Null' 'Value'
func GetNull (v Value) bool {
	if v == nil {
		return false
	}
	_, ok := v.(*Null)
	return ok
}

// 'SetNull' returns a 'Null' 'Value'
func SetNull () Value {
	return &Null{}
}

//*********** Json -> Go ***********

// Implementation of the standard procedures of Babel

func (d *directory) ReadInt () int32 {
	m := uint32(0)
	p := uint(0)
	for i := 0; i < 4; i++ {
		n, err := d.r.ReadByte()
		M.Assert(err == nil, err, 100)
		m += uint32(n) << p
		p += 8
	}
	return int32(m)
} //ReadInt

func (c *compilationer) Read () (ch rune, cLen int) {
	if c.pos >= c.size {
		ch = C.EOF1
	} else {
		ch = c.r[c.pos]
		c.pos++
	}
	cLen = 1
	return
} //Read

func (c *compilationer) Pos () int {
	return c.pos
} //Pos

func (c *compilationer) SetPos (pos int) {
	c.pos = min(pos, c.size)
} //SetPos

func (c *compilationer) Error (pos, line, col int, msg string) {
}

func (c *compilationer) Map (index string, p ...string) string {
	return index
}

const (
	
	object = 1
	array = 2
	nilC = 3
	cons = 4
	
	json = 1
	stringC = 2
	number = 3
	trueC = 4
	falseC = 5
	null = 6

)

func eval (o *C.Object) Json {
	
	EvalVal := func (o *C.Object) Value {
		var v Value
		switch o.ObjType() {
		case C.UserObj:
			v = o.ObjUser().(Value)
		case C.TermObj:
			v = &JsonVal{eval(o)}
		}
		return v
	} //EvalVal

	EvalObj := func (o *C.Object) Fields {
		oo := o; n := 0
		for oo.ObjFunc() == cons {
			n++
			oo = oo.ObjTermSon(3)
		}
		M.Assert(oo.ObjFunc() == nilC, 100)
		f := make(Fields, n)
		for i := 0; i < n; i++ {
			oo := o.ObjTermSon(1)
			s := oo.ObjString()
			f[i] = &Field{Name: s[1:len(s) - 1], Value: EvalVal(o.ObjTermSon(2))}
			o = o.ObjTermSon(3)
		}
		return f
	} //EvalObj

	EvalArr := func (o *C.Object) Values {
		oo := o; n := 0
		for oo.ObjFunc() == cons {
			n++
			oo = oo.ObjTermSon(2)
		}
		M.Assert(oo.ObjFunc() == nilC, 100)
		v := make(Values, n)
		for i := 0; i < n; i++ {
			v[i] = EvalVal(o.ObjTermSon(1))
			o = o.ObjTermSon(2)
		}
		return v
	} //EvalArr

	//eval
	var j Json
	M.Assert(o.ObjType() == C.TermObj, 20)
	switch o.ObjFunc() {
	case object:
		j = &Object{EvalObj(o.ObjTermSon(1))}
	case array:
		j = &Array{EvalArr(o.ObjTermSon(1))}
	}
	return j
} //eval

func (c *compilationer) Execution (fNum, parsNb int, pars C.ObjectsList) (o *C.Object, res C.Anyptr, ok bool) {

	MakeString :=  func (rs []rune) string {
		sp := ""
		m := len(rs)
		pos := 0
		for pos < m {
			if rs[pos] == '\\' {
				pos++
				switch rs[pos] {
				case '"':
					sp += string(double_quote)
				case '\\':
					sp += string(reverse_solidus)
				case '/':
					sp += string(solidus)
				case 'b':
					sp += string(backspace)
				case 'f':
					sp += string(form_feed)
				case 'n':
					sp += string(line_feed)
				case 'r':
					sp += string(carriage_return)
				case 't':
					sp += string(horizontal_tab)
				case 'u':
					n := rune(0)
					for j := 1; j <= 4; j++ {
						var x rune
						pos++
						switch c := rs[pos]; {
						case '0' <= c && c <= '9':
							x = c - '0'
						case 'A' <= c && c <= 'F':
							x = c - 'A' + 0xA
						case 'a' <= c && c <= 'f':
							x = c - 'a' + 0xA
						}
						n = n * 0x10 + x
					}
					sp += string(n)
				default:
					M.Halt(100);
				}
			} else {
				sp += string(rs[pos])
			}
			pos++
		}
		return sp
	} //MakeString
	
	//Execution
	switch fNum {
	case json:
		o = pars[0]
		if o.ErrorIn() {
			ok = false
			return
		}
		c.json = eval(o)
	case stringC:
		o = pars[0]
		if o.ErrorIn() {
			ok = false
			return
		}
		r := []rune(o.ObjString())
		res = &String{MakeString(r[1:len(r) - 1])}
	case number:
		o = pars[0]
		if o.ErrorIn() {
			ok = false
			return
		}
		s := o.ObjString()
		i, err := SC.ParseInt(s, 0, 64)
		if err == nil {
			res = &Int{i}
		} else {
			f, err := SC.ParseFloat(s, 64); M.Assert(err == nil, 100)
			res = &Float{f}
		}
	case trueC, falseC:
		res = &Boolean{fNum == trueC}
	case null:
		res = &Null{}
	}
	ok = true
	return
} //Execution

// 'Compile' builds a Json from an 'io.ReadSeeker'
func Compile (rs io.ReadSeeker) Json {
	n, err := rs.Seek(0, io.SeekEnd); M.Assert(err == nil, err, 100)
	_, err = rs.Seek(0, io.SeekStart); M.Assert(err == nil, err, 101)
	b := make([]byte, n)
	_, err = io.ReadFull(rs, b); M.Assert(err == nil, err, 102)
	r := []rune(string(b))
	co := &compilationer{r: r, size: len(r), pos: 0, json: nil}
	c := C.NewCompilation(co)
	if c.Compile(comp, true) {
		return co.json
	}
	return nil
} //Compile

type
	
	ReadSeekCloser interface {
		io.ReadCloser
		io.Seeker
	}

// string -> Json
func ReadString (s string) Json {
	return Compile(strings.NewReader(s))
} //ReadString

// ReadSeekCloser -> Json
func ReadRSC (rsc ReadSeekCloser) Json {
	defer rsc.Close()
	return Compile(rsc)
} //ReadRSC

// os.File -> Json
func ReadFile (path string ) Json {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	return ReadRSC(f)
} //ReadFile

// ApplyTo modifies 'obj' according to 'j'; warning: 'obj' must be a pointer to a global, exported and modifiable record variable of its module, and only its exported and modifiable fields are taken into account.
func ApplyTo (j Json, obj interface{}) bool {
	buf := new(strings.Builder)
	j.Write(buf)
	err := J.Unmarshal([]byte(buf.String()), obj)
	return err == nil
} //ApplyTo

func writeJson (w io.Writer, j Json, flat bool, indent int) {
	
	lnC := func () {
		if !flat {
			fmt.Fprintln(w)
		}
	} //lnC
	
	Indent := func (indent int) {
		if !flat {
			for indent > 0 {
				fmt.Fprint(w, "\t")
				indent--
			}
		}
	} //Indent
	
	WriteString := func (s string) {
		
		// Code special chars of 's' with backslashed representations
		JsonFilter := func (s string) string {
			ss := ""
			for _, r := range s {
				escape := true
				var c rune
				switch r {
				case double_quote, reverse_solidus:
					c = r
				case backspace:
					c = 'b'
				case form_feed:
					c = 'f'
				case line_feed:
					c = 'n'
				case carriage_return:
					c = 'r'
				case horizontal_tab:
					c = 't'
				default:
					escape = false
					c = r
				}
				if escape {
					ss += "\\"
				}
				ss += string(c)
			}
			return ss
		} //JsonFilter
		
		//WriteString
		fmt.Fprintf(w, "%s", "\"")
		fmt.Fprintf(w, "%s", JsonFilter(s))
		fmt.Fprintf(w, "%s", "\"")
	} //WriteString
	
	WriteVal := func (v Value, indent int) {
		if _, b := v.(*JsonVal); !b {
			Indent(indent)
		}
		switch v := v.(type) {
		case *String:
			WriteString(v.S)
		case *Int:
			fmt.Fprintf(w, "%s", SC.FormatInt(v.N, 10))
		case *Float:
			fmt.Fprintf(w, "%s", SC.FormatFloat(v.F, 'g', -1, 64))
		case *JsonVal:
			writeJson(w, v.Json, flat, indent)
		case *Boolean:
			if v.Boolean {
				fmt.Fprintf(w, "%s", "true")
			} else {
				fmt.Fprintf(w, "%s", "false")
			}
		case *Null:
			fmt.Fprintf(w, "%s", "null")
		default:
			M.Halt("Incorrect value type", 100)
		}
	} //WriteVal
	
	WriteArr := func (vs Values, indent int) {
		if len(vs) > 0 {
			for i, v := range vs {
				if i > 0 {
					fmt.Fprintf(w, "%s", ",")
					lnC()
				}
				WriteVal(v, indent)
			}
			lnC()
		}
	} //WriteArr
	
	WriteObj := func (fs Fields, indent int) {
		
		WriteField := func (f *Field, indent int) {
			Indent(indent)
			WriteString(f.Name)
			fmt.Fprintf(w, "%s", ":")
			lnC()
			WriteVal(f.Value, indent + 1)
		} //WriteField
		
		//WriteObj
		if len(fs) > 0 {
			for i, f := range fs {
				if i > 0 {
					fmt.Fprintf(w, "%s", ",")
					lnC()
				}
				WriteField(f, indent)
			}
			lnC()
		}
	} //WriteObj

	//writeJson
	switch j := j.(type) {
	case *Object:
		Indent(indent)
		fmt.Fprintf(w, "%s", "{")
		lnC()
		WriteObj(j.Fields, indent + 1)
		Indent(indent)
		fmt.Fprintf(w, "%s", "}")
	case *Array:
		Indent(indent)
		fmt.Fprintf(w, "%s", "[")
		lnC()
		WriteArr(j.Elements, indent + 1)
		Indent(indent)
		fmt.Fprintf(w, "%s", "]")
	}
} //writeJson

// 'GetString' returns a formatted string representing 'o'
func (o *Object) GetString () string {
	M.Assert(o != nil, 20)
	var buf = new(bytes.Buffer)
	writeJson(buf, o, false, 0)
	return buf.String()
} //GetString

// 'GetString' returns a formatted string representing 'a'
func (a *Array) GetString () string {
	M.Assert(a != nil, 20)
	var buf = new(bytes.Buffer)
	writeJson(buf, a, false, 0)
	return buf.String()
} //GetString

// 'GetFlatString' returns a compact string representing 'o'
func (o *Object) GetFlatString () string {
	M.Assert(o != nil, 20)
	var buf = new(bytes.Buffer)
	writeJson(buf, o, true, 0)
	return buf.String()
} //GetFlatString

// 'GetFlatString' returns a compact string representing 'a'
func (a *Array) GetFlatString () string {
	M.Assert(a != nil, 20)
	var buf = new(bytes.Buffer)
	writeJson(buf, a, true, 0)
	return buf.String()
} //GetFlatString

// 'Write' writes a formatted string representing 'o' with 'w'
func (o *Object) Write (w io.Writer) {
	writeJson(w, o, false, 0)
} //Write

// 'Write' writes a formatted string representing 'a' with 'w'
func (a *Array) Write (w io.Writer) {
	writeJson(w, a, false, 0)
} //Write

// 'WriteFlat' writes a compact string representing 'o' with 'w'
func (o *Object) WriteFlat (w io.Writer) {
	writeJson(w, o, true, 0)
} //WriteFlat

// 'WriteFlat' writes a compact string representing 'a' with 'w'
func (a *Array) WriteFlat (w io.Writer) {
	writeJson(w, a, true, 0)
} //WriteFlat

// Replace backslashed representations of special characters in 's' by these characters themselves
func DecodeString (s string) string {
	
	const (
		
		backspace = string(rune(0x8))
		form_feed = string(rune(0xC))
		line_feed = string(rune(0xA))
		carriage_return = string(rune(0xD))
		horizontal_tab = string(rune(0x9))
	
	)
	
	rs := []rune(s)
	m := len(rs)
	sp := ""
	pos := 0
	for pos < m {
		if rs[pos] == '\\' {
			pos++
			switch rs[pos] {
			case '"':
				sp += "\""
			case '\\':
				sp += "\\"
			case '/':
				sp += "/"
			case 'b':
				sp += backspace
			case 'f':
				sp += form_feed
			case 'n':
				sp += line_feed
			case 'r':
				sp += carriage_return
			case 't':
				sp += horizontal_tab
			case 'u':
				var n rune = 0
				for j := 1; j <= 4; j++ {
					pos++
					var x rune
					switch c := rs[pos]; {
						case '0' <= c && c <= '9':
							x = c - '0'
						case 'A' <= c && c <= 'F':
							x = c - 'A' + 0xA
						case 'a' <= c && c <= 'f':
							x = c - 'a' + 0xA
					}
					n = n * 0x10 + x
				}
				sp += string(n)
			}
		} else {
			sp += string(rs[pos])
		}
		pos++
	}
	return sp
} //DecodeString

//*********** Go -> Json ***********

// stack procedures

func (s *stack) next () stacker {
	return s.nxt
}

func (s *stack) setNext (ss stacker) {
	s.nxt = ss
}

func (m *Maker) push (s stacker) {
	s.setNext(m.stk)
	m.stk = s
}

func (m *Maker) pull () stacker {
	M.Assert(m.stk != nil, 20);
	s := m.stk
	m.stk = s.next()
	return s
}

// 'StartObject' begins a json object
func (m *Maker) StartObject () {
	m.push(new(startObj))
}

// 'StartArray' begins a json array
func (m *Maker) StartArray () {
	m.push(new(startArr));
}

// 'PushString' pushes the string 's' on the stack
func (m *Maker) PushString (s string) {
	m.push(&stValue{val: &String{s}})
}

// 'PushInteger' pushes the integer 'n' on the stack
func (m *Maker) PushInteger (n int64) {
	m.push(&stValue{val: &Int{n}})
}

// 'PushReal' pushes the float 'f' on the stack
func (m *Maker) PushFloat (f float64) {
	m.push(&stValue{val: &Float{f}})
}

// 'PushBoolean' pushes the boolean 'b' on the stack
func (m *Maker) PushBoolean (b bool) {
	m.push(&stValue{val: &Boolean{b}})
}

// 'PushNull' pushes a json null value on the stack
func (m *Maker) PushNull () {
	m.push(&stValue{val: new(Null)})
}

// 'PushValue' pushes the value 'v' on the stack
func (m *Maker) PushValue (v Value) {
	m.push(&stValue{val: v})
}

// 'PushJson' pushes the Json 'j' on the stack
func (m *Maker) PushJson (j Json) {
	m.push(&stValue{val: &JsonVal{j}})
}

// 'PushField' pushes the field 'f' on the stack
func (m *Maker) PushField (f *Field) {
	m.push(&stField{name: f.Name, value: f.Value})
}

// 'BuildField' builds the json object field whose name is 'name' and whose value is the last element on the stack and replaces this element with the built field
func (m *Maker) BuildField (name string) {
	s, ok := m.pull().(*stValue)
	M.Assert(ok, 100)
	m.push(&stField{name: name, value: s.val})
}

// 'BuildObject' builds a json 'Object' from all the fields stacked after the last 'StartObject' call, remove these fields and stacks the 'Object'
func (m *Maker) BuildObject () {
	s := m.stk
	n := 0
	for {
		ok := s != nil
		if ok {
			_, ok = s.(*stField)
		}
		if !ok {break}
		n++
		s = s.next()
	}
	ok := s != nil
	if ok {
		_, ok = s.(*startObj)
	}
	M.Assert(ok, 100)
	o := new(Object)
	if n == 0 {
		o.Fields = nil
	} else {
		o.Fields = make(Fields, n)
		for i := n - 1; i >= 0; i-- {
			sf := m.pull().(*stField)
			o.Fields[i] = &Field{sf.name, sf.value}
		}
	}
	m.pull()
	m.push(&stValue{val: &JsonVal{o}})
}

// 'BuildArray' builds a json 'Array' from all the values stacked from the last 'StartArray' call, remove these values and stacks the 'Array'
func (m *Maker) BuildArray () {
	s := m.stk
	n := 0
	for {
		ok := s != nil
		if ok {
			_, ok = s.(*stValue)
		}
		if !ok {break}
		n++
		s = s.next()
	}
	ok := s != nil
	if ok {
		_, ok = s.(*startArr)
	}
	M.Assert(ok, 100)
	a := new(Array)
	if n == 0 {
		a.Elements = nil
	} else {
		a.Elements = make(Values, n)
		for i := n - 1; i >= 0; i-- {
			a.Elements[i] = m.pull().(*stValue).val
		}
	}
	m.pull()
	m.push(&stValue{val: &JsonVal{a}})
}

// stack[1] <- stack[n]
func (m *Maker) Roll (n int) {
	M.Assert(n >= 1, 20)
	s := m.stk
	for s != nil && n > 2 {
		s = s.next(); n--
	}
	M.Assert(s != nil && s.next() != nil, 21)
	if n >= 2 {
		ss := s.next()
		s.setNext(ss.next())
		ss.setNext(m.stk)
		m.stk = ss
	}
}

// stack[n] <- stack[1]
func (m *Maker) RollD (n int) {
	M.Assert(n >= 1, 20)
	s := m.stk
	for s != nil && n > 1 {
		s = s.next(); n--
	}
	M.Assert(s != nil, 21)
	if n >= 2 {
		ss := m.stk
		m.stk = m.stk.next()
		ss.setNext(s.next())
		s.setNext(ss)
	}
}

// stack[2] <-> stack[1]
func (m *Maker) Swap () {
	m.Roll(2)
}

// GetJson builds and returns a Json built from the last object on the stack
func (m *Maker) GetJson () Json {
	ok := m.stk != nil
	var (sv *stValue; jv *JsonVal)
	if ok {
		sv, ok = m.stk.(*stValue)
	}
	if ok {
		jv, ok = sv.val.(*JsonVal)
	}
	M.Assert(ok, 20)
	return jv.Json
}

// 'NewMaker' creates and returns a 'Maker' able to build 'Json' by stack manipulations
func NewMaker () *Maker {
	return &Maker{stk: (*stack)(nil)}
}

func process (v R.Value, m *Maker) {
	switch v.Kind() {
	case R.Bool:
		m.PushBoolean(v.Bool())
	case R.Int32:
		if v.Type().Name() == "rune" {
			m.PushString(string(rune(v.Int())))
		} else {
			m.PushInteger(v.Int())
		}
	case R.Int, R.Int8, R.Int16, R.Int64:
		m.PushInteger(v.Int())
	case R.Uint, R.Uint8, R.Uint16, R.Uint32, R.Uint64, R.Uintptr:
		m.PushInteger(int64(v.Int()))
	case R.Float32, R.Float64:
		m.PushFloat(v.Float())
	case R.Complex64, R.Complex128:
		c := v.Complex()
		m.StartObject()
		m.PushFloat(real(c))
		m.BuildField("real")
		m.PushFloat(imag(c))
		m.BuildField("imag")
		m.BuildObject()
	case R.Array, R.Slice:
		m.StartArray()
		for k := 0; k < v.Len(); k++ {
			process(v.Index(k), m)
		}
		m.BuildArray()
	case R.Interface:
		p := v.Elem()
		process(p, m)
	case R.Ptr:
		if v.IsNil() {
			m.PushNull()
		} else {
			p := v.Elem()
			process(p, m)
		}
	case R.String:
		m.PushString(v.String())
	case R.Struct:
		t := v.Type()
		n := t.NumField()
		m.StartObject()
		for k := 0; k < n; k++ {
			process(v.Field(k), m)
			m.BuildField(t.Field(k).Name)
		}
		m.BuildObject()
	default:
		M.Halt("Incorrect value type", 100)
	}
} //process

// BuildJsonFrom builds and returns a 'Json' describing 'obj'; warning: 'obj' must be nil or a pointer to a global exported struct variable of its package, and only its exported fields are taken into account
func BuildJsonFrom (obj interface{}) Json {
	if obj == nil {
		return nil
	}
	v := R.ValueOf(obj)
	m := NewMaker()
	process(v, m)
	return m.GetJson()
} //BuildJsonFrom

// 'path' is a concatenation of '.' denoting an 'Object', identifiers (letters and digits) denoting 'Fields' of 'Object'(s) and of '[<number>] denoting an element of an 'Array'; 'Find(j, path)' returns a '*ValueStruct' containing the value pointed by 'path' in 'j' and a possible error (see examples in json_test.go); the first value of 'expectedKinds', if present, tells which kind of value is expected, the second one tells, in the case of a json value, if an object or an array is expected, if these conditions are present and not respected an error is triggered
func Find (j Json, path string, expectedKinds ... int) (*ValueStruct, error) {
	
	ld := func (r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}
	
	valueOf := func (v Value) *ValueStruct {
		vs := new(ValueStruct)
		vs.Kind = v.VKind()
		switch vv := v.(type) {
		case *JsonVal:
			vs.Json = vv.Json
		case *String:
			vs.String = vv.S
		case *Int:
			vs.Int = vv.N
		case *Float:
			vs.Float = vv.F
		case *Boolean:
			vs.Boolean = vv.Boolean
		case *Null:
		}
		return vs
	}
	
	pos := 0
	if path[0] != '.' && path[0] != '[' {
		path = "." + path
		pos = -1
	}
	p := []rune(path)
	l := len(p)
	vs := new(ValueStruct)
	for l != 0 {
		var (v Value; ok bool; n int)
		switch p[0] {
		case '.':
			if j.JKind() != ObjectKind {
				return nil, errors.New("No object at " + strconv.Itoa(pos))
			}
			o := j.GetObject()
			n = 1
			pos++
			for n < l && ld(p[n]) {
				n++
			}
			if n == 1 {
				return nil, errors.New("No field at " + strconv.Itoa(pos))
			}
			id := string(p[1:n])
			v, ok = GetValue(o, id)
			if !ok {
				return nil, errors.New("No field at " + strconv.Itoa(pos))
			}
		case '[':
			if j.JKind() != ArrayKind {
				return nil, errors.New("No array at " + strconv.Itoa(pos))
			}
			a := j.GetArray()
			n = 1
			pos++
			for n < l && unicode.IsDigit(p[n]) {
				n++
			}
			if n == l || p[n] != ']' {
				return nil, errors.New("No ']' detected at " + strconv.Itoa(pos + n - 1))
			}
			if n == 1 {
				return nil, errors.New("No element at " + strconv.Itoa(pos))
			}
			nb, err := strconv.Atoi(string(p[1:n]))
			if err != nil {
				return nil, errors.New("Not a number at " + strconv.Itoa(pos))
			}
			if nb >= len(a.Elements) {
				return nil, errors.New("Out of range at " + strconv.Itoa(pos))
			}
			n++
			v = a.Elements[nb]
		default:
			return nil, errors.New("Unexpected character at " + strconv.Itoa(pos))
		}
		p = p[n:]
		l = len(p)
		vs = valueOf(v)
		if vs.Kind == JsonKind {
			j = vs.Json
		} else {
			if l != 0 {
				return nil, errors.New("Not a JSON at " + strconv.Itoa(pos))
			}
		}
		pos += n - 1
	}
	if expectedKinds != nil && vs.Kind != expectedKinds[0] {
		var kind string 
		switch vs.Kind {
		case JsonKind:
			kind = "Json"
		case StringKind:
			kind = "String"
		case IntKind:
			kind = "Int"
		case FloatKind:
			kind = "Float"
		case BooleanKind:
			kind = "Boolean"
		case NullKind:
			kind = "Null"
		}
		return nil, errors.New(kind + " kind was not expected")
	}
	if expectedKinds != nil && len(expectedKinds) > 1 && vs.Kind == JsonKind && vs.Json.JKind() != expectedKinds[1] {
		var jKind string
		switch vs.Json.JKind() {
		case ObjectKind:
			jKind = "Object"
		case ArrayKind:
			jKind = "Array"
		}
		return nil, errors.New(jKind + " kind was not expected")
	}
	return vs, nil
}

// Fix the current Json compiler from its Babel representation
func SetRComp (rComp io.Reader) {
	comp = C.NewDirectory(&directory{r: bufio.NewReader(rComp)}).ReadCompiler()
}

func init () {
	// Compilers reading
	f, err := os.Open(compPath)
	if err == nil {
		defer f.Close()
		SetRComp(f)
	}
} //init
