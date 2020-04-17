package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCond(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-conditional-statements/problem", func(t *testing.T) {
		{
			actual := cond(3)
			assert.Equal(t, weird, actual)
		}

		{
			actual := cond(24)
			assert.Equal(t, notWeird, actual)
		}
	})
}
