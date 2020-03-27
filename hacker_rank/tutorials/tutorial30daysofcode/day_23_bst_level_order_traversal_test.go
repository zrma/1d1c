package tutorial30daysofcode

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-binary-trees/problem", func(t *testing.T) {
		var root *treeNode
		for _, num := range []int{3, 5, 4, 7, 2, 1} {
			root = insert(root, num)
		}
		diff := deep.Equal(levelOrder(root), []int{3, 2, 5, 1, 4, 7})
		assert.Nil(t, diff)
	})
}
