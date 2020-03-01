package tutorial30daysofcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestDivisorSum(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-interfaces/problem", func(t *testing.T) {
		c := &calculator{}

		for i, data := range []struct {
			input    int
			expected int
		}{
			{1, 1},
			{6, 12},
			{20, 42},
		} {
			assert.Equal(t, divisorSum(c, data.input), data.expected, "failed %d", i)
		}
	})
}
