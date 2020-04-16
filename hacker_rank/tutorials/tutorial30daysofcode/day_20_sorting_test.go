package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestSorting(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-sorting/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			sorting([]int32{1, 2, 3})
			sorting([]int32{3, 2, 1})
		}, []string{
			"Array is sorted in 0 swaps.",
			"First Element: 1",
			"Last Element: 3",
			"Array is sorted in 3 swaps.",
			"First Element: 1",
			"Last Element: 3",
		})
		assert.NoError(t, err)
	})
}
