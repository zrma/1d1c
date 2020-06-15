package greedy

import (
	"github.com/zrma/going/utils/integer"
)

func candies(n int32, arr []int32) int64 {
	candies := make([]int32, n)
	var candy int32 = 1

	candies[0] = candy
	for i := 1; i < int(n); i++ {
		if arr[i] > arr[i-1] {
			candy++
		} else {
			candy = 1
		}
		candies[i] = candy
	}

	candy = 1
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			candy++
		} else {
			candy = 1
		}
		candies[i] = integer.MaxInt32(candies[i], candy)
	}

	var sum int64
	for _, val := range candies {
		sum += int64(val)
	}
	return sum
}
