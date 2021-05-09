package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestHourGlassSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-2d-arrays/problem")

	err := utils.PrintTest(func() {
		arr := [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		hourGlassSum(arr)
	}, []string{
		"19",
	})
	assert.NoError(t, err)
}
