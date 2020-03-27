package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHeight(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-binary-search-trees/problem", func(t *testing.T) {
		var root *treeNode
		for _, num := range []int{3, 5, 2, 1, 4, 6, 7} {
			root = insert(root, num)
		}

		const expected int32 = 3
		actual := getHeight(root)
		assert.Equal(t, expected, actual)
	})
}
