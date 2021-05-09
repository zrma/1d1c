package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicArray(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/2d-array/problem")

	const givenN = 2
	givenArr := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}
	want := []int32{7, 3}

	got := dynamicArray(givenN, givenArr)
	assert.Equal(t, want, got)
}
