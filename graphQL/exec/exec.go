/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and / or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/ 

package exec

import (
	
	A	"github.com/gerard94/util/avl"
	F	"path/filepath"
	G	"github.com/gerard94/util/graphQL"
	J	"github.com/gerard94/util/json"
	M	"github.com/gerard94/util/misc"
	R	"github.com/gerard94/util/resources"
		"fmt"
		"io"
		"os"
	_	"github.com/gerard94/util/graphQL/exec/static"

)

const (

	execDir = "util/graphQL/exec"

)

type (
	
	responseEventManager struct { //G.ResponseEventManager
		w io.Writer
	}
	
	ResolversFixer  func (ts G.TypeSystem)
	
	Executor struct {
		typeSystemName string
		ts G.TypeSystem
		es G.ExecSystem
		doc *G.Document
		variableValues *A.Tree
		operationNames G.StrArray
		documentRead,
		variableValuesRead bool
	}
	
	vTypeSystem struct {
		G.TypeSystem
	}

)

var (
	
	execPath = R.FindDir(execDir)

)

func printErrors (w io.Writer, errors *A.Tree) {
	e := errors.Next(nil)
	for e != nil {
		el := e.Val().(*G.ErrorElem)
		fmt.Fprint(w, el.Message)
		if el.Locations != nil {
			fmt.Fprint(w, " at")
			for i, loc := range el.Locations {
				if i > 0 {
					fmt.Fprint(w, ",")
				}
				fmt.Fprint(w, " ", loc.L, ":", loc.C)
			}
		}
		p := el.Path
		if p != nil {
			fmt.Fprint(w, " in ")
			for {
				switch p := p.(type) {
				case *G.PathString:
					fmt.Fprint(w, p.S)
				case *G.PathNb:
					fmt.Fprint(w, p.N)
				}
				p = p.Next()
				if p != nil {
					fmt.Fprint(w, ".")
				}
				if p == nil {break}
			}
		}
		fmt.Fprintln(w)
		e = errors.Next(e)
	}
} //printErrors

func (exe *Executor) ReadDocument (path string) bool {
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, path)
	exe.documentRead = false
	var r G.Response
	exe.doc, r = G.ReadGraphQL(path)
	if exe.doc != nil {
		if G.ExecutableDefinitions(exe.doc) {
			if exe.ts == nil {
				fmt.Fprintln(os.Stderr, "Error: No type system")
			} else {
				exe.es = exe.ts.ExecValidate(exe.doc)
				if exe.es.GetErrors().IsEmpty() {
					exe.documentRead = true
					exe.operationNames = exe.es.ListOperations()
					M.Assert(len(exe.operationNames) > 0)
					fmt.Fprintln(os.Stdout, "OK")
					return true
				} else {
					printErrors(os.Stderr, exe.es.GetErrors())
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, "Error: Not executable definitions")
		}
	} else {
		err := r.Errors()
		M.Assert(err != nil && !err.IsEmpty())
		printErrors(os.Stderr, err)
	}
	return false
} //ReadDocument

func (exe *Executor) ReadVariableValues (path string) { 

	Read := func () bool {
		j := J.ReadFile(path)
		if j != nil {
			exe.variableValues = A.New()
			switch j := j.(type) {
			case *J.Object:
				fs := j.Fields
				for _, f := range fs {
					if f.Name == "@" { // enum of 0 length
						return false
					}
					G.InsertJsonValue(exe.variableValues, f.Name, f.Value)
				}
				return true
			default:
			}
		}
		return false
	} //Read

	// ReadVariableValues
	exe.variableValuesRead = Read()
} //ReadVariableValues

func printValue (m *J.Maker, value G.Value) {
	
	PrintList := func (lv *G.ListValue) {
		m.StartArray()
		for l := lv.First(); l != nil; l = lv.Next(l) {
			printValue(m, l.Value)
		}
		m.BuildArray()
	} //PrintList
	
	//printValue
	switch value := value.(type) {
	case *G.IntValue:
		m.PushInteger(value.Int)
	case *G.FloatValue:
		m.PushFloat(value.Float)
	case *G.StringValue:
		if value == nil {
			m.PushNull()
		} else {
			m.PushString(value.String.S)
		}
	case *G.BooleanValue:
		m.PushBoolean(value.Boolean)
	case *G.NullValue:
		m.PushNull()
	case *G.EnumValue:
		m.PushString(value.Enum.S)
	case *G.ListValue:
		PrintList(value)
	case *G.OutputObjectValue:
		printObject(m, value)
	}
} //printValue

func printObject (m *J.Maker, ov *G.OutputObjectValue) {
	
	PrintField := func (f *G.ObjectField) {
		printValue(m, f.Value)
		m.BuildField(f.Name.S)
	} //PrintField
	
	//printObject
	m.StartObject()
	for f := ov.First(); f != nil; f = ov.Next(f) {
		PrintField(f)
	}
	m.BuildObject()
} //printObject

func printResponse (w io.Writer, r G.Response) {
	
	PrintData := func (ov *G.OutputObjectValue) {
		m := J.NewMaker()
		printObject(m, ov)
		m.GetJson().Write(w)
		fmt.Fprintln(w)
	} //PrintData
	
	//printResponse
	if r.Errors() != nil {
		printErrors(w, r.Errors())
	}
	switch r := r.(type) {
	case *G.InstantResponse:
		if r.Data != nil {
			PrintData(r.Data)
		}
	case *G.SubscribeResponse:
		if r.Data != nil {
			m := new(responseEventManager)
			f, err := os.Create(F.Join(execPath, r.Name)); M.Assert(err == nil, err, 100)
			m.w = f
			/*
			r.Data.RecordResponseEventManager(m)
			*/
			/**/
			r.Data.FixResponseStream(m)
			/**/
		}
	}
} //printResponse

func (m *responseEventManager) ManageResponseEvent (r G.Response) {
	printResponse(m.w, r)
} //ManageResponseEvent

func (exe *Executor) ExecuteDocument (w io.Writer, opName string) {
	if exe.documentRead {
		vars := exe.variableValues
		if vars == nil {
			vars = A.New()
		}
		r := exe.es.Execute(exe.doc, opName, vars)
		printResponse(w, r)
	}
} //ExecuteDocument

func (exe *Executor) SetTypeSystem (newTS G.TypeSystem, systemName, systemPath string, fixResolvers ResolversFixer, initialValue *G.OutputObjectValue) {
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, systemPath)
	w := os.Stderr
	exe.typeSystemName = ""
	exe.variableValuesRead = false
	exe.documentRead = false
	exe.operationNames = make(G.StrArray, 0)
	var r G.Response
	exe.doc, r = G.ReadGraphQL(systemPath)
	err := r.Errors()
	if err != nil && !err.IsEmpty() {
		printErrors(w, err)
		return
	}
	fmt.Fprintln(os.Stdout, "OK")
	M.Assert(exe.doc != nil, 100)
	tsRead := false
	newTS.InitTypeSystem(exe.doc)
	if G.ExecutableDefinitions(exe.doc) {
		newTS.MappedError("#util/graphQL/exec:NoTypeSystemInDoc", "", "", nil, nil)
	} else {
		fixResolvers(newTS)
		newTS.FixInitialValue(initialValue)
		tsRead = newTS.GetErrors().IsEmpty()
	}
	if !tsRead {
		exe.ts = nil
		printErrors(w, newTS.GetErrors())
		fmt.Fprintln(w)
		return
	}
	exe.ts = newTS
	exe.typeSystemName = systemName
} //SetTypeSystem

func (exe *Executor) TypeSystemName () string {
	return exe.typeSystemName
}

func (exe *Executor) OperationNames () G.StrArray {
	return exe.operationNames
}

func NewExecutor () *Executor {
	return &Executor{}
}

func (exe *Executor) CopyExecutorTS () *Executor {
	return &Executor{ts: exe.ts, typeSystemName: exe.typeSystemName}
}

func (ts *vTypeSystem) FixScalarCoercer (scalarName string, sc *G.ScalarCoercer) {
}

func VerifyTypeSystem (systemPath string) {
	fmt.Fprintln(os.Stdout)
	fmt.Fprintln(os.Stdout, systemPath)
	doc, r := G.ReadGraphQL(systemPath)
	err := r.Errors()
	if err != nil && !err.IsEmpty() {
		printErrors(os.Stderr, err)
		return
	}
	M.Assert(doc != nil, 100)
	tsRead := false
	ts := new(vTypeSystem)
	ts.TypeSystem = G.Dir.NewTypeSystem(ts)
	ts.InitTypeSystem(doc)
	if G.ExecutableDefinitions(doc) {
		ts.MappedError("NoTypeSystemInDoc", "", "", nil, nil)
	} else {
		tsRead = ts.GetErrors().IsEmpty()
	}
	if tsRead {
		fmt.Fprintln(os.Stdout, "OK")
	} else {
		printErrors(os.Stderr, ts.GetErrors())
	}
} //VerifyTypeSystem

func init () {
	os.MkdirAll(execPath, 0777)
}
