package goutils

// GCD implements greatest common divisor euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM implements least common multiple algorithm for an arbitrary number of
// integers.
func LCM(s ...int) int {
	switch len(s) {
	case 0:
		return 0
	case 1:
		return s[0]
	}
	res := s[0] * s[1] / GCD(s[0], s[1])
	for _, i := range s[2:] {
		res = LCM(res, i)
	}
	return res
}
