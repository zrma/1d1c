package tutorial30daysofcode

import (
	"1d1go/utils/integer"
)

type treeNode struct {
	data        int
	left, right *treeNode
}

func insert(root *treeNode, data int) *treeNode {
	if root == nil {
		return &treeNode{data: data}
	}

	var cur *treeNode
	if data <= root.data {
		cur = insert(root.left, data)
		root.left = cur
	} else {
		cur = insert(root.right, data)
		root.right = cur
	}
	return root
}

func getHeight(root *treeNode) int32 {
	if root == nil {
		return -1
	}

	return integer.MaxInt32(getHeight(root.left), getHeight(root.right)) + 1
}
