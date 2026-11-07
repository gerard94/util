package alea

import (
    "testing"
    "math"
    "fmt"
)

var
	
	gT *Generator = New()

func TestGenerator_Random (_ *testing.T) {
	const n = 100
	for i := 0; i < n; i++ {
		fmt.Println(gT.Random())
	}
	fmt.Println()
}

func TestGenerator_IntRand (t *testing.T) {
	const (n = 100; min = 1000; max = 10000)
	for i := 0; i < n; i++ {
		fmt.Println(gT.IntRand(min, max))
	}
	fmt.Println()
}

func TestGenerator_GaussRand (t *testing.T) {
	
	const (
		n = 10000000
		m = 2.
		sig = 3.
	)
	
	var
		x float64
	
	s := 0.; s2 := 0.
	p := 0;
	for i := 1; i <= n; i++ {
		x = m + sig * gT.GaussRand()
		s = s + x
		s2 = s2 + x * x
		if math.Abs(x - m) < sig {
			p += 1
		}
	}
	s = s / n
	s2 = math.Sqrt(s2 / n - s * s)
	fmt.Printf("Mean: %v (%v)\n", s, m)
	fmt.Printf("Standard deviation: %v (%v)\n", s2, sig)
	fmt.Printf("prop. at sigma: %v%% (~68.27%%)\n", 100. * float64(p) / n)
	fmt.Println()
}
