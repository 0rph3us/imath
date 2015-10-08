package imath

func Sqrt2(n int) int {
	// set the highest bit
	c := 1 << 31
	g := 1 << 31

	for {
		if g*g > n {
			g ^= c
		}
		c >>= 1
		if c == 0 {
			return g
		}
		g |= c
	}
}

func Sqrt(x int) (root int) {
	op := x
	root = 0

	// "one" starts at the highest power of four <= than the argument.
	one := 1 << 30 // second-to-top bit set
	for one > op {
		one >>= 2
	}

	for one != 0 {
		if op >= root+one {
			op -= root + one
			root += one << 1 // 2 * one
		}
		root >>= 1
		one >>= 2
	}
	return
}

func Sqrt3(n int) (root int) {
	remainder := n
	place := 0x40000000
	root = 0

	for place > remainder {
		place = place >> 2
	}
	for place != 0 {
		if remainder >= root+place {
			remainder = remainder - root - place
			root = root + (place << 1)
		}
		root = root >> 1
		place = place >> 2
	}
	return
}
