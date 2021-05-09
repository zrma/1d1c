package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeLevelOrder(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-binary-trees/problem")

	var root *treeNode
	for _, num := range []int{3, 5, 4, 7, 2, 1} {
		root = insert(root, num)
	}
	want := []int{3, 2, 5, 1, 4, 7}

	got := levelOrder(root)
	assert.Equal(t, want, got)
}
