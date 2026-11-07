/*
Babel: a compiler compiler.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package main
	
/* The module testgram is part of the Babel subsystem, a compiler compiler. testgram tests whether a given text satisfies a given grammar. */

import (
	A "git.duniter.org/gerard94/util/avl"
	B "git.duniter.org/gerard94/util/babel/compil"
	M	"git.duniter.org/gerard94/util/misc"
	R	"git.duniter.org/gerard94/util/resources"
		"bufio"
		"io"
		"os"
		"fmt"
)

type (
	
	directory struct {
		r *bufio.Reader
	}
	
	compilation struct {
		c B.Compilation
		r []rune
		size,
		pos int
		err *A.Tree
	}
	
	errElem struct {
		pos,
		line,
		col int
		mes string
	}
)

func (e1 *errElem) Compare (e2 A.Comparer) A.Comp {
	e := e2.(*errElem)
	switch {
		case e1.pos < e.pos:
			return A.Lt
		case e1.pos > e.pos:
			return A.Gt
		case e1.mes < e.mes:
			return A.Lt
		case e1.mes > e.mes:
			return A.Gt
		default:
			return A.Eq
	}
}

func (d *directory) ReadInt () (res int32) {
	m := uint32(0)
	p := uint(0)
	for i := 0; i < 4; i++ {
		n, err := d.r.ReadByte()
		M.Assert(err == nil, err, 100)
		m += uint32(n) << p
		p += 8
	}
	res = int32(m)
	return
}

func (c *compilation) Read () (ch rune, cLen int) {
	if c.pos >= c.size {
		ch = B.EOF1
	} else {
		ch = c.r[c.pos]
		c.pos++
	}
	cLen = 1
	return
}

func (c *compilation) Pos () int {
	return c.pos
}

func (c *compilation) SetPos (pos int) {
	c.pos = min(pos, c.size)
}

func (c *compilation) Execution (fNum, parsNb int, pars B.ObjectsList) (objPos *B.Object, res B.Anyptr, ok bool) {
	ok = true
	return
}

func (comp *compilation) Error (p int, li, co int, msg string) {
	if comp.err != nil {
		e := &errElem{pos: p, line: li, col: co, mes: msg}
		_, b, _ := comp.err.SearchIns(e)
		M.Assert(!b, 100)
	}
}

func (c *compilation) Map (index string, p ...string) (val string) {
	switch index {
		case "CForbidden":
			val = "Unexpected character."
		case "COpenedComment":
			val = "Opened comment at the end of text."
		case "CNoComment":
			val = "Unexpected end of comment."
		case "Cor":
			val = "or"
		case "Cexpected":
			val = "expected."
	}
	return
}

func setGram (name string) *B.Compiler {
	name = R.FindDir(name)
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	d := B.NewDirectory(&directory{r: bufio.NewReader(f)})
	return d.ReadCompiler()
}

func setReader (name string) ([]rune) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fi, err := f.Stat()
	M.Assert(err == nil, err, 101)
	b := make([]byte, fi.Size())
	_, err = io.ReadFull(f, b)
	M.Assert(err == nil, err, 102)
	return []rune(string(b))
}

func main () {
	if len(os.Args) != 3 {
		fmt.Println("$testgram <rsrc path to compiler.tbl> <path to the text to verify>")
		return
	}
	comp := setGram(os.Args[1])
	r := setReader(os.Args[2])
	if comp == nil || r == nil {
		return
	}
	c := &compilation{r: r, size: len(r), pos: 0, err: A.New()}
	co := B.NewCompilation(c)
	c.c = co
	if co.Compile(comp, false) {
		fmt.Println("OK")
	} else {
		e := c.err.Next(nil)
		for e != nil {
			ee := e.Val().(*errElem)
			fmt.Fprintf(os.Stderr, "%v:%v: %v\n", ee.line, ee.col, ee.mes)
			e = c.err.Next(e)
		}
	}
}
