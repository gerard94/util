package sort

import (
	
	A	"git.duniter.org/gerard94/util/avl"
	
		"testing"
	
		"git.duniter.org/gerard94/util/alea"
		"bytes"
		"fmt"
		"unicode"

)

type (
	
	list []int

)

func (l list) Less (p1, p2 int) bool {
	return l[p1] < l[p2]
}

func (l list) Swap (p1, p2 int) {
	l[p1], l[p2] = l[p2], l[p1]
}

func TestSort1 (t *testing.T) {
	
	const (
		
		n = 1000000
		m = 5000000
		repeat = 20
	
	)
	
	fmt.Println("TestSort1")
	l := make(list, n)
	var u TS
	u.Sorter = l
	for j := repeat; j > 0; j-- {
		fmt.Println(j)
		fmt.Println("Phase 1")
		for i := 0; i < n; i++ {
			l[i] = int(alea.Random() * float64(m))
		}
		fmt.Println("Phase 2")
		u.QuickSort(0, len(l) - 1)
		fmt.Println("Phase 3")
		for i := 1; i < len(l); i++ {
			if l[i - 1] > l[i] {
				fmt.Println("i =", i, "l[i] =", l[i])
				t.Fail()
			}
		}
	}
	fmt.Println()
}

type (
	
	Comp = A.Comp

)

const (
	
	Lt = A.Lt
	Eq = A.Eq
	Gt = A.Gt

)

// Extract the significant characters in s; only alphanumeric characters are significant, and their case of lowest rank is returned
func ToDown (s string) string {
	rs := bytes.Runes([]byte(s))
	buf := new(bytes.Buffer)
	for _, r := range rs {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(min(unicode.ToUpper(r), unicode.ToLower(r)))
		}
	}
	return string(buf.Bytes())
} //ToDown

// Extract the significant characters in s; only alphanumeric characters are significant, and are returned
func strip (s string) string {
	rs := bytes.Runes([]byte(s))
	buf := new(bytes.Buffer)
	for _, r := range rs {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}
	return string(buf.Bytes())
} //strip

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

func compS (s1, s2 string) Comp {
	rs1 := bytes.Runes([]byte(s1))
	rs2 := bytes.Runes([]byte(s2))
	l1 := len(rs1)
	l2 := len(rs2)
	l := min(l1, l2)
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

// Standard comparison procedure for identifiers; they are first compared with only significant characters and case ignored, and if still equal, with all characters and case taken into account
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

type (
	
	idVal struct {
		uid string
		val float64
	}
	
	idVals []*idVal

)

func (l idVals) Less (iv1, iv2 int) bool {
	return l[iv1].val > l[iv2].val || l[iv1].val == l[iv2].val && CompP(l[iv1].uid, l[iv2].uid) == Lt
} //Less

func (l idVals) Swap (iv1, iv2 int) {
	l[iv1], l[iv2] = l[iv2], l[iv1]
} //Swap

func TestSort2 (t *testing.T) {

	const (
		
		n = 100000
		m = 2000
		r = 5
		min = 0x20
		max = 0x10000
	
	)
	
	fmt.Println("TestSort2")
	l := make(idVals, n)
	var u TS
	u.Sorter = l
	fmt.Println("Phase 1")
	for i := 0; i < n; i++ {
		v := float64(alea.IntRand(0, m)) / float64(m)
		b := make([]rune, r)
		for j := 0; j < r; j++ {
			b[j] = rune(alea.IntRand(min, max))
		}
		s := string(b)
		l[i] = &idVal{s, v}
	}
	fmt.Println("Phase 2")
	u.QuickSort(0, len(l) - 1)
	fmt.Println("Phase 3")
	for i := 0; i < len(l) - 1; i++ {
		if l.Less(i + 1, i) {
			fmt.Println("i =", i, "l[i] =", l[i])
			t.Fail()
		}
	}
	fmt.Println()
}
