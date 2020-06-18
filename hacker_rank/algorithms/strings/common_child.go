package strings

import "github.com/zrma/going/utils/integer"

// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func commonChild(s1 string, s2 string) int32 {
	l1 := len(s1) + 1
	l2 := len(s2) + 1
	arr := make([][]int32, l2)
	for i := range arr {
		arr[i] = make([]int32, l2)
	}

	for i := 0; i < l1; i++ {
		arr[0][i] = 0
		arr[i][0] = 0
	}

	// n * m 문제
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				arr[i+1][j+1] = arr[i][j] + 1
			} else {
				arr[i+1][j+1] = integer.MaxInt32(arr[i+1][j], arr[i][j+1])
			}
		}
	}

	return arr[l1-1][l2-1]
}
