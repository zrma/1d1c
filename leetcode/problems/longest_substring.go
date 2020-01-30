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

	// 한 문자가 다시 출현할 때까지의 최대 거리를 구한다.
	// abcabc
	// |--|
	// 0123
	//
	// prevIdx = 0
	// curLength -> 2까지 증가
	//
	// i < 3 일 경우
	//
	// b, c의 경우 모두 prevIdx == -1
	//
	// i == 1 : 1 - 1 == 0 (prevIdx == 0)
	// i == 2 : 2 - 1 == 0 -> curLength++ -> 2
	// i == 3 : curLength(2) > maxLength(1) -> maxLength(2), curLength = 3 - 0 = 3
	//
	// ... 중략
	//
	// curLength(3) > maxLength(2) -> maxLength(3)
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
