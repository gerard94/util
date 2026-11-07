/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package misc

import (
	
	F	"path/filepath"
	SC	"strconv"
		"log"
		"os"
		"strings"
		"testing"

)

const (
	
	version = "1.0.9"

)

type (
	
	Nerror interface {
		error
		ErrorNum () int
	}
	
	numError struct {
		n int
		s string
	}

)

var (
	
	stdLg = log.New(os.Stderr, "", log.LstdFlags)
	lg = stdLg

)

func GetStdLog () *log.Logger {
	return stdLg
}

func GetLog () *log.Logger {
	return lg
}

func SetLog (log *log.Logger) {
	lg = log
}

func Odd (n int) bool {
	return n & 1 == 1
}

func Abs (a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func Abs32 (a int32) int32 {
	if a < 0 {
		a = -a
	}
	return a
}

func Abs64 (a int64) int64 {
	if a < 0 {
		a = -a
	}
	return a
}

func AbsF32 (a float32) float32 {
	if a < 0 {
		a = -a
	}
	return a
}

func AbsF64 (a float64) float64 {
	if a < 0 {
		a = -a
	}
	return a
}

func haltCommon (flag ... interface{}) {
	lg.Println("*** ERROR ***");
	if len(flag) == 0 {
		lg.Println("Assert error")
		panic("Assert error")
	}
	for _, f := range flag {
		switch e := f.(type) {
		case error:
			lg.Println(e.Error())
		case string:
			lg.Println(e)
		case int:
			lg.Println(SC.Itoa(e))
		default:
			lg.Println(f)
		}
	}
	panic(flag[len(flag) - 1])
}


func Halt (flag ... interface{}) {
	haltCommon(flag ...)
}

func Assert (cond bool, flag ... interface{}) {
	if !cond {
		haltCommon(flag ...)
	}
}

func Want (cond bool, t *testing.T) {
	if !cond {
		t.Fail()
	}
}

func hidePath (shown string) string {
	dir, name := F.Split(shown)
	return dir + "." + name
}

func showPath (hidden string) string {
	dir, name := F.Split(hidden)
	Assert(len(name) > 1 && name[0] == '.', 20)
	return dir + strings.Replace(name, ".", "", 1)
}

func InstantCreate (name string) (*os.File, error) {
	return os.Create(hidePath(name))
}

func InstantClose (f *os.File) error {
	hidden := f.Name()
	err := f.Close()
	if err != nil {
		return err
	}
	return os.Rename(hidden, showPath(hidden))
}

func Version () string {
	return version
}

func (e *numError) Error () string {
	return e.s
}

func (e *numError) ErrorNum () int {
	return e.n
}
 
func NewError (text string, number int) Nerror {
	return &numError{s: text, n: number}
}

const (
	
	year0 = 1970
	month0 = 1
	day0 = 1

)

var (
	
	d0 int

)

func Date2Day (year, month, day int) int {
	y := year
	m := month - 3
	if m < 0 {
		m += 12
		y--
	}
	d := y * 1461 / 4 + (m * 153 + 2) / 5 + day - 306
	if d > 577737 {
		d -= (y / 100 * 3 - 5) / 4
	}
	return d
}

func Timestamp2Time (year, month, day, hour, minute, second int) int64 {
	ts := int64(Date2Day(year, month, day))
	ts = (ts - int64(d0)) * 24 + int64(hour)
	ts = ts * 60 + int64(minute)
	ts = ts * 60 + int64(second)
	return ts
}

func Day2Date (d int) (year, month, day int) {
	var c int
	if d > 577737 {
		d = d * 4 +1215
		c = d / 146097
		d = d % 146097 / 4
	} else {
		d += 305
		c = 0
	}
	d = d * 4 + 3
	y := d / 1461
	d = d % 1461 / 4
	d = d * 5 + 2
	m := d / 153
	d = d % 153 / 5
	if m > 9 {
		m -= 12
		y++
	}
	year = 100 * c + y
	month = m + 3
	day = d + 1
	return
}

func Time2Timestamp (ts int64) (year, month, day, hour, minute, second int) {
	second = int(ts % 60); ts /= 60
	minute = int(ts % 60); ts /= 60
	hour = int(ts % 24); ts /= 24
	year, month, day = Day2Date(int(ts) + d0)
	return
}

func init () {
	d0 = Date2Day(year0, month0, day0)
}
