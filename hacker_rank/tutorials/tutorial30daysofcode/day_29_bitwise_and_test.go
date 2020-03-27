package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseAnd(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-bitwise-and/problem", func(t *testing.T) {
		for i, data := range []struct {
			n, k     int32
			expected int32
		}{
			{5, 2, 1},
			{8, 5, 4},
			{2, 2, 0},
			{955, 236, 235},
		} {
			actual := bitwiseAND(data.n, data.k)
			assert.Equal(t, data.expected, actual, i)
		}
	})
}
