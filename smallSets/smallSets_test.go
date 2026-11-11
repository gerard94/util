package smallSets

import (
	
	M	"github.com/gerard94/util/misc"
		"github.com/gerard94/util/alea"
		"testing"
		"fmt"
		"math"

)

func TestSets (t *testing.T) {
	s1 := Set(int64(alea.UintRand(0, math.MaxUint64)))
	s2 := Set(int64(alea.UintRand(0, math.MaxUint64)))
	sU := s1.Union(s2)
	sI := s1.Inter(s2)
	sS := s1.XOR(s2)
	sD := sU.Diff(sI)
	fmt.Println("s1 =", s1)
	fmt.Println("s2 =", s2)
	fmt.Println("sU =", sU)
	fmt.Println("sI =", sI)
	fmt.Println("sD =", sD)
	fmt.Println("sS =", sS)
	M.Want(sS == sD, t)
	M.Want(s1.Union(s1) == s1, t)
	M.Want(s2.Inter(s2) == s2, t)
	M.Want(s1.Union(Full()) == Full(), t)
	M.Want(s2.Inter(NewSet()) == NewSet(), t)
	M.Want(s1.Union(s1.Inter(s2)) == s1, t)
	M.Want(s1.Inter(s1.Union(s2)) == s1, t)
	M.Want(s1.Inter(s2) == s1.Diff(s1.Diff(s2)), t)
}
