/*
util: Set of tools.

Copyright (C) 2001-2020 Gérard Meunier

This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA 02111-1307, USA.
*/

package resources

// resources.FindDir looks for a directory containing resources for the executable that started the current process.

import (
	
	F	"path/filepath"
	M	"git.duniter.org/gerard94/util/misc"
		"os"

)

func inSrc () string {
	d, err := os.Getwd(); M.Assert(err == nil, err, 100)
	for ; d != "." && d != "/" && d != ".."; d = F.Dir(d) {
		f, err := os.Open(d); M.Assert(err == nil, err, d, 101)
		defer f.Close()
		list, err := f.Readdirnames(0); M.Assert(err == nil, err, 102)
		for _, fi := range list {
			if fi == "go.mod" {
				for d = F.Dir(d); F.Base(d) != "src" && d != "." && d != "/" && d != ".."; d = F.Dir(d) {
				}
				M.Assert(F.Base(d) == "src", d, 103)
				return d
			}
		}
	}
	return ""
}

func findModsRec (d *os.File, root string, l *[]string) {
	list, err := d.Readdirnames(0); M.Assert(err == nil, err, 104)
	for _, n := range list {
		r := F.Join(root, n)
		if n == "go.mod" {
			*l = append(*l, F.Dir(r))
		} else {
			dd, err := os.Open(r);  M.Assert(err == nil, err, r, 105)
			defer dd.Close()
			fi, err := dd.Stat(); M.Assert(err == nil, err, 106)
			if fi.IsDir() {
				findModsRec(dd, r, l)
			}
		}
	}
}

func findMods (root string) []string {
	l := make([]string, 0)
	d, err := os.Open(root);  M.Assert(err == nil, err, root, 108)
	defer d.Close()
	findModsRec(d, root, &l)
	return l
}

func splitFirstDir (path string) (string, string) {
	first := ""; rest := ""
	for ; path != "." && path != "/"; path = F.Dir(path) {
		rest = F.Join(first, rest)
		first = F.Base(path)
	}
	return first, rest
}

func FindDir (path string) string {
	var r string
	root := inSrc()
	if root != "" {
		mod, rest := splitFirstDir(path)
		list := findMods(root)
		for _, d := range list {
			if F.Base(d) == mod {
				r = F.Join(d, "rsrc", rest)
			}
		}
		M.Assert(r != "", "Unknown module", 109)
	} else {
		root, err := os.Getwd(); M.Assert(err == nil, err, 110)
		r = F.Join(root, "rsrc", path)
	}
	err := os.MkdirAll(r, 0777); M.Assert(err == nil, err, r, 111)
	return r
}
