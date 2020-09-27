package greedy

import (
	"github.com/zrma/going/utils/integer"
)

func candies(n int32, arr []int32) int64 {
	candies := make([]int32, n)

	candies = forward(candies, arr)
	candies = backward(candies, arr)

	var sum int64
	for _, val := range candies {
		sum += int64(val)
	}
	return sum
}

func forward(candies, arr []int32) []int32 {
	var candy int32 = 1
	candies[0] = candy
	for i := 1; i < len(candies); i++ {
		if arr[i] > arr[i-1] {
			candy++
		} else {
			candy = 1
		}
		candies[i] = candy
	}
	return candies
}

func backward(candies, arr []int32) []int32 {
	var candy int32 = 1
	for i := len(candies) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			candy++
		} else {
			candy = 1
		}
		candies[i] = integer.MaxInt32(candies[i], candy)
	}
	return candies
}
