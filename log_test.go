package imath

import (
	"testing"
)

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
