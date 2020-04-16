package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDifference(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-scope/problem", func(t *testing.T) {
		difference := newDifference([]int{1, 2, 5})
		assert.Equal(t, 4, difference.computeDifference())
	})
}
