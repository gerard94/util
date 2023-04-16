/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and / or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/ 

package validation

import (
	
	G	"git.duniter.org/gerard94/util/graphQL"
	E	"git.duniter.org/gerard94/util/graphQL/exec"
	F	"path/filepath"
	R	"git.duniter.org/gerard94/util/resources"

)

const (
	
	typeSystemName = "Validation"
	
	tsDir = "util/graphQL/examples/validation"
	tsName = "01 - Type System.txt"

)

type (
	
	typeSystem struct {
		G.TypeSystem
	}

)

var (
	
	tsPath = F.Join(R.FindDir(tsDir), tsName)

)

func (ts *typeSystem) FixFieldResolvers () {
} //FixFieldResolvers

func (ts *typeSystem) FixStreamResolvers () {
} //FixStreamResolvers

func abstractTypeResolver (ts G.TypeSystem, td G.TypeDefinition, ov *G.OutputObjectValue) *G.ObjectTypeDefinition {
	return nil
} //abstractTypeResolver

func fixResolvers (ts G.TypeSystem) {
	newTS := &typeSystem{TypeSystem: ts}
	newTS.FixFieldResolvers()
	newTS.FixStreamResolvers()
	newTS.FixAbstractTypeResolver(abstractTypeResolver)
} //fixResolvers

func (ts *typeSystem) FixScalarCoercer (scalarName string, sc *G.ScalarCoercer) {
}

func NewTypeSystem () *E.Executor {
	
	const (
		
		rootName = "root"
	
	)
	
	ts := new(typeSystem)
	ts.TypeSystem = G.Dir.NewTypeSystem(ts)
	ov := G.NewOutputObjectValue()
	ov.InsertOutputField(rootName, G.MakeAnyValue(nil))
	exe := E.NewExecutor()
	exe.SetTypeSystem(ts.TypeSystem, typeSystemName, tsPath, fixResolvers, ov)
	return exe
} //NewTypeSystem

func Initialize () {
} //Initialize

func init () {
	Initialize()
} //init
