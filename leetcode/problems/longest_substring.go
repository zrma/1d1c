package problems

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	curLength := 1
	maxLength := 1
	prevIdx := 0

	var visited [math.MaxUint8]int
	for i := range visited {
		visited[i] = -1
	}
	visited[s[0]] = 0

	for i := 1; i < len(s); i++ {
		prevIdx = visited[s[i]]
		if prevIdx == -1 || (i-curLength > prevIdx) {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			curLength = i - prevIdx
		}
		visited[s[i]] = i
	}

	if curLength > maxLength {
		maxLength = curLength
	}
	return maxLength
}
