package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeGetHeight(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-binary-search-trees/problem")

	var root *treeNode
	for _, num := range []int{3, 5, 2, 1, 4, 6, 7} {
		root = insert(root, num)
	}
	const want = 3

	got := getHeight(root)
	assert.EqualValues(t, want, got)
}
