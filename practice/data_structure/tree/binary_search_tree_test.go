package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("삽입, 조회가 잘 된다", func(t *testing.T) {
		bst := binarySearchTree{}

		bst.insert(9)
		bst.insert(5)
		bst.insert(3)
		bst.insert(7)
		bst.insert(1)

		var result []int
		bst.traverse(func(data int) {
			result = append(result, data)
		})

		assert.Equal(t, result, []int{1, 3, 5, 7, 9})

		t.Run("정확한 위치에 잘 들어간다", func(t *testing.T) {
			//           9
			//       5
			//    3    7
			// 1
			assert.Equal(t, bst.top.data, 9, "9")

			assert.NotNil(t, bst.top.left, "5")
			assert.Equal(t, bst.top.left.data, 5, "5")

			assert.Nil(t, bst.top.right)

			assert.NotNil(t, bst.top.left.left, "3")
			assert.Equal(t, bst.top.left.left.data, 3, "3")

			assert.NotNil(t, bst.top.left.right, "7")
			assert.Equal(t, bst.top.left.right.data, 7, "7")

			assert.NotNil(t, bst.top.left.left.left, "1")
			assert.Equal(t, bst.top.left.left.left.data, 1, "1")

			assert.Nil(t, bst.top.left.left.right, "nil")
		})
	})

	t.Run("중복 입력해도 한 번만 입력", func(t *testing.T) {
		bst := binarySearchTree{}

		bst.insert(3)
		bst.insert(1)
		bst.insert(2)
		bst.insert(1)
		bst.insert(2)
		bst.insert(3)
		bst.insert(3)

		var result []int
		bst.traverse(func(data int) {
			result = append(result, data)
		})

		assert.Equal(t, result, []int{1, 2, 3})
	})
}
