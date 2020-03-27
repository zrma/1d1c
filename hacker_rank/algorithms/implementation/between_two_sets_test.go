package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoTotalX(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/between-two-sets/problem", func(t *testing.T) {
		{
			const expected int32 = 3
			actual := getTotalX([]int32{2, 4}, []int32{16, 32, 96})
			assert.Equal(t, expected, actual)
		}

		{
			const expected int32 = 2
			actual := getTotalX([]int32{3, 4}, []int32{24, 48})
			assert.Equal(t, expected, actual)
		}
	})
}
