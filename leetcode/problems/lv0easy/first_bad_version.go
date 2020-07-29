package lv0easy

var isBadVersion func(int) bool

func solveFirstBadVersion(n, expected int) (int, int) {
	var callCount int
	isBadVersion = func(i int) bool {
		callCount++
		return i >= expected
	}
	return firstBadVersion(n), callCount
}

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
