package lv1medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/zrma/1d1go/leetcode/problems/common"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Log("https://leetcode.com/problems/add-two-numbers/")

	build := func(arr []int) *ListNode {
		l := &ListNode{}
		cur := l
		for _, n := range arr {
			cur.Next = &ListNode{Val: n}
			cur = cur.Next
		}
		return l.GetNext()
	}

	for i, tt := range []struct {
		l1, l2 []int
		want   []int
	}{
		{
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			l1:   []int{5},
			l2:   []int{5},
			want: []int{0, 1},
		},
		{
			l1:   nil,
			l2:   []int{5, 6, 4},
			want: []int{5, 6, 4},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l1 := build(tt.l1)
			l2 := build(tt.l2)
			want := build(tt.want)
			got := addTwoNumbers(l1, l2)
			assert.Equal(t, want.Traversal(), got.Traversal())
		})
	}
}
