package lv0easy

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	var rev int
	y := x
	for y > 0 {
		z := y % 10
		rev = rev*10 - z
		y = y / 10
	}

	return x+rev == 0
}
