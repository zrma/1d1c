package lv0easy

var guess func(int) int

func solveGuessNumber(n, expected int) int {
	guess = func(i int) int {
		switch {
		case expected < i:
			return -1
		case expected == i:
			return 0
		case expected > i:
			return 1
		}
		return 0
	}
	return guessNumber(n)
}

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		v := guess(mid)
		switch {
		case v == 0:
			return mid
		case v < 0:
			r = mid - 1
		case v > 0:
			l = mid + 1
		}
	}
	return -1
}
