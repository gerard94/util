package csv

import (
	
	C	"git.duniter.org/gerard94/util/babel/compil"
	M	"git.duniter.org/gerard94/util/misc"
	R	"git.duniter.org/gerard94/util/resources"
	SM	"git.duniter.org/gerard94/util/strMapping"
		"bufio"
		"fmt"
		"io"
		"os"
		"strings"

)

const (
	
	compPath = "util/csv/csv.tbl"

)

type (
	
	Line []string
	
	Csv []Line
	
	directory struct { //C.Directory
		r *bufio.Reader
	}
	
	compilationer struct {
		r []rune
		size,
		pos int
		csv Csv
	}

)

var (
	
	lang = SM.NewStdLanguage()
	rsrcCompPath = R.FindDir(compPath)
	comp *C.Compiler

)

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
	if pos <= c.size {
		c.pos = pos
	}
} //SetPos

func (c *compilationer) Error (pos, line, col int, msg string) {
	fmt.Printf("\n***CSV ERROR***: %v:%v: %v\n", line, col, msg)
}

func (c *compilationer) Map (index string, p ...string) string {
	return lang.Map("#babel:" + index, p ...)
}

func (c *compilationer) Execution (fNum, parsNb int, pars C.ObjectsList) (o *C.Object, res C.Anyptr, ok bool) {
	
	const (
		
		appendC = 101
		null = 102
		record = 103
	
	)
	
	switch fNum {
	case appendC:
		o = C.Parameter(pars, 1)
		line := o.ObjUser().(Line)
		o1 := C.Parameter(pars, 2)
		typ := o1.ObjType()
		switch typ {
		case C.TermObj:
			line = append(line, "")
			res = line
		case C.StringObj:
			val := o1.ObjString()
			if val[0] == '"' {
				b := []rune(val)
				b = b[1:len(b) - 1]
				bb := make([]rune, 0)
				for i := 0; i < len(b); {
					bb = append(bb, b[i])
					if b[i] == '"' {
						M.Assert(i < len(b) - 1 && b[i + 1] == '"', string(b), i, 100)
						i += 2
					} else {
						i++
					}
				}
				val = string(bb)
			}
			line = append(line, val)
		default:
			M.Halt(typ, 100)
		}
		res = line
	case null:
		line := make(Line, 0)
		res = line
	case record:
		o = C.Parameter(pars, 1)
		line := o.ObjUser().(Line)
		c.csv = append(c.csv, line)
	}
	ok = true
	return
} //Execution

func Compile (rs io.ReadSeeker) Csv {
	n, err := rs.Seek(0, io.SeekEnd); M.Assert(err == nil, err, 100)
	_, err = rs.Seek(0, io.SeekStart); M.Assert(err == nil, err, 101)
	b := make([]byte, n)
	_, err = io.ReadFull(rs, b); M.Assert(err == nil, err, 102)
	r := []rune(string(b))
	co := &compilationer{r: r, size: len(r), pos: 0, csv: make(Csv, 0)}
	c := C.NewCompilation(co)
	if c.Compile(comp, true) {
		n := len(co.csv)
		csv := make(Csv, n)
		for i, cs := range co.csv {
			csv[n - 1 - i] = cs
		}
		return csv
	}
	return nil
} //Compile

func CompileString (s string)  Csv {
	return Compile(strings.NewReader(s))
} //CompileString

func CompileFile (path string) Csv {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	return Compile(f)
}

func SetRComp (rComp io.Reader) {
	comp = C.NewDirectory(&directory{r: bufio.NewReader(rComp)}).ReadCompiler()
}

func init () {
	f, err := os.Open(rsrcCompPath)
	if err == nil {
		defer f.Close()
		SetRComp(f)
	}
}
