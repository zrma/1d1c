package problems

import "math"

func rev(x int) int {
	var res int
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

func reverse(x int) int {
	negative := x < 0
	if negative {
		x = -x
		return -rev(x)
	}
	return rev(x)
}
