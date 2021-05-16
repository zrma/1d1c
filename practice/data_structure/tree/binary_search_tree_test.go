package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	bst := buildBinarySearchTree(t)

	//           9
	//       5
	//    3    7
	// 1
	assert.Equal(t, 9, bst.top.data)

	assert.NotNil(t, bst.top.left)
	assert.Equal(t, 5, bst.top.left.data)

	assert.Nil(t, bst.top.right)

	assert.NotNil(t, bst.top.left.left)
	assert.Equal(t, 3, bst.top.left.left.data)

	assert.NotNil(t, bst.top.left.right)
	assert.Equal(t, 7, bst.top.left.right.data)

	assert.NotNil(t, bst.top.left.left.left)
	assert.Equal(t, 1, bst.top.left.left.left.data)

	assert.Nil(t, bst.top.left.left.right)
}

func TestBinarySearchTreeRemoveDuplicated(t *testing.T) {
	bst := buildBinarySearchTree(t)
	bst.insert(3)
	bst.insert(1)
	bst.insert(2)
	bst.insert(1)
	bst.insert(2)
	bst.insert(3)
	bst.insert(3)

	want := []int{1, 2, 3, 5, 7, 9}
	var got []int
	bst.traverse(func(data int) {
		got = append(got, data)
	})
	assert.Equal(t, want, got)
}

func buildBinarySearchTree(t *testing.T) binarySearchTree {
	bst := binarySearchTree{}

	bst.insert(9)
	bst.insert(5)
	bst.insert(3)
	bst.insert(7)
	bst.insert(1)

	want := []int{1, 3, 5, 7, 9}
	var got []int
	bst.traverse(func(data int) {
		got = append(got, data)
	})
	assert.Equal(t, want, got)
	return bst
}

func TestBinarySearchTreeExist(t *testing.T) {
	for _, tt := range []struct {
		given int
		want  bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			bst := buildBinarySearchTree(t)
			got := bst.exist(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}

	assert.NotPanics(t, func() {
		b := binarySearchTree{}
		assert.False(t, b.exist(1004))
	})
}
