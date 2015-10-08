package imath

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	vi := []int{
		1,
		2,
		4,
		24,
		48,
		123,
	}

	sqrt := []int{
		1,
		1,
		2,
		4,
		6,
		11,
	}
	for i := 0; i < len(vi); i++ {
		a := vi[i]
		if f := Sqrt(a); sqrt[i] != f {
			t.Errorf("Sqrt(%g) = %g, want %g", a, f, sqrt[i])
		}
	}

	for i := 0; i < len(vi); i++ {
		a := vi[i]
		if f := Sqrt2(a); sqrt[i] != f {
			t.Errorf("Sqrt(%g) = %g, want %g", a, f, sqrt[i])
		}
	}
	for i := 0; i < len(vi); i++ {
		a := vi[i]
		if f := Sqrt3(a); sqrt[i] != f {
			t.Errorf("Sqrt(%g) = %g, want %g", a, f, sqrt[i])
		}
	}
}

func BenchmarkSqrt(b *testing.B) {
	x, y := 0, 100
	for i := 0; i < b.N; i++ {
		x += Sqrt(y)
	}
}

func BenchmarkSqrt2(b *testing.B) {
	x, y := 0, 100
	for i := 0; i < b.N; i++ {
		x += Sqrt2(y)
	}
}

func BenchmarkSqrt3(b *testing.B) {
	x, y := 0, 100
	for i := 0; i < b.N; i++ {
		x += Sqrt3(y)
	}
}
