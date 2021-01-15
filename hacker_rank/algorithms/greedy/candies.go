package greedy

import (
	"github.com/zrma/going/utils/integer"
)

func candies(n int32, arr []int32) int64 {
	result := make([]int32, n)

	result = forward(result, arr)
	result = backward(result, arr)

	var sum int64
	for _, val := range result {
		sum += int64(val)
	}
	return sum
}

func forward(result, arr []int32) []int32 {
	var candy int32 = 1
	result[0] = candy
	for i := 1; i < len(result); i++ {
		if arr[i] > arr[i-1] {
			candy++
		} else {
			candy = 1
		}
		result[i] = candy
	}
	return result
}

func backward(result, arr []int32) []int32 {
	var candy int32 = 1
	for i := len(result) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			candy++
		} else {
			candy = 1
		}
		result[i] = integer.MaxInt32(result[i], candy)
	}
	return result
}
