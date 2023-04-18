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

func inSrc () (root, src string) {
	d, err := os.Getwd(); M.Assert(err == nil, err, 100)
	for ; d != "." && d != "/" && d != ".."; d = F.Dir(d) {
		f, err := os.Open(d); M.Assert(err == nil, err, d, 101)
		defer f.Close()
		list, err := f.Readdirnames(0); M.Assert(err == nil, err, 102)
		for _, fi := range list {
			if fi == "go.mod" {
				dd := d
				for dd = F.Dir(dd); F.Base(dd) != "src" && dd != "." && dd != "/" && dd != ".."; dd = F.Dir(dd) {
				}
				if F.Base(dd) == "src" {
					return d, dd
				}
				return d, ""
			}
		}
	}
	return "", ""
}

func findModsRec (d *os.File, root string, l *[]string) {
	list, err := d.Readdirnames(0); M.Assert(err == nil, err, 103)
	for _, n := range list {
		r := F.Join(root, n)
		if n == "go.mod" {
			*l = append(*l, F.Dir(r))
		} else {
			dd, err := os.Open(r);  M.Assert(err == nil, err, r, 104)
			defer dd.Close()
			fi, err := dd.Stat(); M.Assert(err == nil, err, 105)
			if fi.IsDir() {
				findModsRec(dd, r, l)
			}
		}
	}
}

func findMods (root string) []string {
	l := make([]string, 0)
	d, err := os.Open(root);  M.Assert(err == nil, err, root, 106)
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

// Let 'path' = "origin/rest/...". If the current directory is in a module whose root directory is "<absolute path>/root", FindDir returns "<absolute path>/root/rsrc/rest/...". But, if this module is a descendant of a directory named 'src' and if several modules are contained in 'src' and if one of these modules has root == origin, then FindDir returns "<absolute path>/origin/rsrc/rest". If the current directory "<current absolute path>" is not in a module, FindDir returns "<current absolute path>/rsrc/origin/rest"
func FindDir (path string) string {
	var r string
	root, src := inSrc()
	if root != "" {
		mod, rest := splitFirstDir(path)
		if src != "" {
			list := findMods(src)
			for _, d := range list {
				if F.Base(d) == mod {
					r = F.Join(d, "rsrc", rest)
				}
			}
			if  r == "" {
				r = F.Join(root, "rsrc", rest)
			}
		} else {
			r = F.Join(root, "rsrc", rest)
		}
	} else {
		root, err := os.Getwd(); M.Assert(err == nil, err, 107)
		r = F.Join(root, "rsrc", path)
	}
	return r
}
