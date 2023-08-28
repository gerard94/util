package strings

import (
	
	M	"git.duniter.org/gerard94/util/misc"
	
	"bytes"
	"strings"
	"unicode"

)

type (
	
	// Result of a comparaison.
	Comp = int8

)

const (
	
	// Comp; result of a comparaison.
	Lt	= Comp(-1) // Less than.
	Eq	= Comp(0) // Equal.
	Gt	= Comp(+1) // Greater than.

)

// 'ToDown' extracts the significant characters in s; only alphanumeric characters are significant, and their case of lowest rank is returned
func ToDown (s string) string {
	rs := bytes.Runes([]byte(s))
	buf := new(bytes.Buffer)
	for _, r := range rs {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(M.Min32(unicode.ToUpper(r), unicode.ToLower(r)))
		}
	}
	return string(buf.Bytes())
} //ToDown

// Extract the significant characters in s; only alphanumeric characters are significant, and are returned
func strip (s string) string {
	rs := bytes.Runes([]byte(s))
	var buf bytes.Buffer
	for _, r := range rs {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}
	return string(buf.Bytes())
} //strip

// 'compR' compares two runes: an alphanumeric rune is less than a non-alphanumeric one; otherwise the unicode order is kept
func compR (r1, r2 rune) Comp {
	ld1 := unicode.IsLetter(r1) || unicode.IsDigit(r1)
	ld2 := unicode.IsLetter(r2) || unicode.IsDigit(r2)
	if ld1 && !ld2 {
		return Lt
	}
	if !ld1 && ld2 {
		return Gt
	}
	if r1 < r2 {
		return Lt
	}
	if r1 > r2 {
		return Gt
	}
	return Eq
} //compR

// 'compS' compares two strings, extending lexicographically the rune order of 'compR'
func compS (s1, s2 string) Comp {
	rs1 := bytes.Runes([]byte(s1))
	rs2 := bytes.Runes([]byte(s2))
	l1 := len(rs1)
	l2 := len(rs2)
	l := M.Min(l1, l2)
	for i := 0; i < l; i++ {
		switch compR(rs1[i], rs2[i]) {
		case Lt:
			return Lt
		case Gt:
			return Gt
		case Eq:
		}
	}
	if l1 < l2 {
		return Lt
	}
	if l1 > l2 {
		return Gt
	}
	return Eq
} //compS

// Standard comparison procedure for identifiers; they are first compared with only significant characters and case ignored, and if still equal, with only significant characters and case taken into account, and if still equal, with all characters and case taken into account
func CompP (s1, s2 string) Comp {
	ss1 := ToDown(s1); ss2 := ToDown(s2)
	if ss1 < ss2 {
		return Lt
	}
	if ss1 > ss2 {
		return Gt
	}
	ss1 = strip(s1); ss2 = strip(s2)
	if ss1 < ss2 {
		return Lt
	}
	if ss1 > ss2 {
		return Gt
	}
	return compS(s1, s2)
} //CompP

// 'compR' compares the characters of s1 and s2 and returns whether the first one is a prefix of the second one or not
func Prefix (s1, s2 string) bool {
	return strings.HasPrefix(s2, s1)
} //Prefix
