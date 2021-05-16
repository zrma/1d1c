package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNode(t *testing.T) {
	node := ListNode{
		Val: 123,
		Next: &ListNode{
			Val: 456,
		},
	}

	{
		const want = 123
		got := node.GetVal()
		assert.Equal(t, want, got)
	}

	{
		next := node.GetNext()
		assert.NotNil(t, next)

		const want = 456
		got := next.GetVal()
		assert.Equal(t, want, got)
	}

	t.Run("Traversal 함수 검증", func(t *testing.T) {
		const want = "456123"
		got := node.Traversal()
		assert.Equal(t, want, got)
	})
}

//noinspection GoNilness
func TestListNodeNil(t *testing.T) {
	assert.NotPanics(t, func() {
		var node *ListNode
		assert.Zero(t, node.GetVal())
		assert.Zero(t, node.GetNext().GetVal())
	})
}
