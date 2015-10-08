package imath

func Sqrt2(n int) int {
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

func Sqrt(x int) int {
	var op, res, one int

	op = x
	res = 0

	// "one" starts at the highest power of four <= than the argument.
	one = 1 << 30 // second-to-top bit set
	for one > op {
		one >>= 2
	}

	for one != 0 {
		if op >= res+one {
			op -= res + one
			res += one << 1 // 2 * one
		}
		res >>= 1
		one >>= 2
	}
	return res
}

func Sqrt3(n int) int {
	var root, remainder, place int

	root = 0
	remainder = n
	place = 0x40000000

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
	return root
}
