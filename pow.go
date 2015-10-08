package imath

func Pow(a, n int) int {
	if n < 0 {
		return 0
	}
	x := 1
	for n != 0 {
		if (n & 1) == 1 {
			x *= a
		}
		n >>= 1
		a *= a
	}

	return x
}
