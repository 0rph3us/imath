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
