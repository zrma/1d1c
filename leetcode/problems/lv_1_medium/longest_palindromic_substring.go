package lv_1_medium

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 || (length <= 3 && s[0] == s[length-1]) {
		return s
	}

	isPalindrome := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	var min, max int
	maxLength := 1
	for i := 1; i < length-1; i++ {
		for offset := 0; offset <= length/2; offset++ {
			start := i - offset
			if start < 0 {
				break
			}

			// 짝수 길이 0
			{
				end := i + offset - 1
				if end >= length {
					break
				}

				if isPalindrome(start, end) && (end-start+1) > maxLength {
					maxLength = end - start + 1
					min = start
					max = end
				}
			}

			// 홀수 길이
			{
				end := i + offset
				if end >= length {
					break
				}

				if isPalindrome(start, end) && (end-start+1) > maxLength {
					maxLength = end - start + 1
					min = start
					max = end
				}
			}

			// 짝수 길이 1
			{
				end := i + offset + 1
				if end >= length {
					break
				}

				if isPalindrome(start, end) && (end-start+1) > maxLength {
					maxLength = end - start + 1
					min = start
					max = end
				}
			}
		}
	}
	return s[min : max+1]
}
