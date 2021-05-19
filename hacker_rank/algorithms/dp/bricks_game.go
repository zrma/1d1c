package dp

import (
	"1d1go/utils/integer"
)

func bricksGame(arr []int32) int64 {
	length := len(arr)
	if length <= 3 {
		var sum int64
		for _, n := range arr {
			sum += int64(n)
		}
		return sum
	}

	result := make([]int64, length)
	var total int64

	lastIndex := length - 1

	// case 0-2
	// arr    :  1,   2,   3,   4,   5
	// total  : 15 ← 14 ← 12  ← 9  ← 5
	// result :  6,   9,  12,   9,   5

	// case 1
	// arr    :  999,   1,   1,   1,   0
	// total  : 1002 ←  3 ←  2  ← 1  ← 0
	// result : 1001,   2,   2,   1,   0

	// case 3
	// arr    :    0,     1,     1,     1,   999,  999
	// total  : 2001 ← 2001 ← 2000 ← 1999 ← 1998 ← 999
	// result : 1001,  1000,  1001,  1999,  1998,  999

	for i := 0; i < 3; i++ {
		total += int64(arr[lastIndex-i])
		result[lastIndex-i] = total
	}

	for i := 3; i <= lastIndex; i++ {
		total += int64(arr[lastIndex-i])
		result[lastIndex-i] = total - integer.MinInt64(result[lastIndex-i+1], result[lastIndex-i+2], result[lastIndex-i+3])
	}

	return result[0]
}
