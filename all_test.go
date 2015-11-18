package imath

import (
	"testing"
)

func TestAbs(t *testing.T) {
	input := []int{
		0,
		-1,
		2,
		-4,
		24,
		48,
		-123,
	}

	want := []int{
		0,
		1,
		2,
		4,
		24,
		48,
		123,
	}
	for i, n := range input {
		if abs := Abs(n); want[i] != abs {
			t.Errorf("Abs(%d) = %d, want %d", n, abs, want[i])
		}
	}

}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-i)
		Abs(i)
	}
}

func TestLog(t *testing.T) {
	input := []int{
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
		3,
		9,
		27,
		81,
		243,
		729,
		2187,
		6561,
		19683,
		59049,
	}

	want2 := []int{
		0,
		3,
		6,
		9,
		12,
		15,
		18,
		21,
		24,
		27,
		30,
		1,
		3,
		4,
		6,
		7,
		9,
		11,
		12,
		14,
		15,
	}

	for i, n := range input {
		if x := Log2(n); want2[i] != x {
			t.Errorf("Log2(%d) = %d, want %d", n, x, want2[i])
		}

	}

}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Log2(i)
	}
}

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
			t.Errorf("Pow(%d) = %d, want %d", n, x, want5[i])
		}

		if x := Pow(8, n); want8[i] != x {
			t.Errorf("Pow(%d) = %d, want %d", n, x, want8[i])
		}

		if x := Pow(10, n); want10[i] != x {
			t.Errorf("Pow(%d) = %d, want %d", n, x, want10[i])
		}
	}

	for i, n := range inputSpecial {
		if x := Pow(5, n); wantSpecial[i] != x {
			t.Errorf("Pow(%d) = %d, want %d", n, x, wantSpecial[i])
		}
	}

}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow(8, 9)
	}
}

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
			t.Errorf("Sqrt(%d) = %d, want %d", a, f, sqrt[i])
		}
	}

	for i := 0; i < len(vi); i++ {
		a := vi[i]
		if f := Sqrt2(a); sqrt[i] != f {
			t.Errorf("Sqrt(%d) = %d, want %d", a, f, sqrt[i])
		}
	}
	for i := 0; i < len(vi); i++ {
		a := vi[i]
		if f := Sqrt3(a); sqrt[i] != f {
			t.Errorf("Sqrt(%d) = %d, want %d", a, f, sqrt[i])
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
