package imath

import (
	"testing"
)

func TestPow(t *testing.T) {
	input := []int{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}

	want5 := []int{
		1,
		5,
		25,
		125,
		625,
		3125,
		15625,
		78125,
		390625,
		1953125,
		9765625,
	}

	want8 := []int{
		1,
		8,
		64,
		512,
		4096,
		32768,
		262144,
		2097152,
		16777216,
		134217728,
		1073741824,
	}

	want10 := []int{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,
	}

	inputSpecial := []int{
		-1,
		-2,
		-100,
	}

	wantSpecial := []int{
		0,
		0,
		0,
	}
	for i, n := range input {
		if x := Pow(5, n); want5[i] != x {
			t.Errorf("Pow(%g) = %g, want %g", n, x, want5[i])
		}

		if x := Pow(8, n); want8[i] != x {
			t.Errorf("Pow(%g) = %g, want %g", n, x, want8[i])
		}

		if x := Pow(10, n); want10[i] != x {
			t.Errorf("Pow(%g) = %g, want %g", n, x, want10[i])
		}
	}

	for i, n := range inputSpecial {
		if x := Pow(5, n); wantSpecial[i] != x {
			t.Errorf("Pow(%g) = %g, want %g", n, x, wantSpecial[i])
		}
	}

}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow(8, 9)
	}
}
