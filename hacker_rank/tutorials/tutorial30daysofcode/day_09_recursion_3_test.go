package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-recursion/problem", func(t *testing.T) {
		{
			actual := factorial(3)
			assert.Equal(t, int32(6), actual)
		}

		{
			actual := factorial(10)
			assert.Equal(t, int32(3628800), actual)
		}
	})
}
