package imath

// Alogorithm is from
// http://graphics.stanford.edu/~seander/bithacks.html#IntegerLogDeBruijn
func Log2(n int) (r int) {

	var multiplyDeBruijnBitPosition = []int{
		0, 9, 1, 10, 13, 21, 2, 29, 11, 14, 16, 18, 22, 25, 3, 30,
		8, 12, 20, 28, 15, 17, 24, 7, 19, 27, 23, 6, 26, 5, 4, 31,
	}

	// De Bruijn sequence
	const b = uint32(0x07C4ACDD)

	if n < 0 {
		panic("log_2 is only for prositive inputs defined")
	}

	// first round down to one less than a power of 2
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16

	i := (uint32(n) * b) >> 27
	r = multiplyDeBruijnBitPosition[i]
	return
}
