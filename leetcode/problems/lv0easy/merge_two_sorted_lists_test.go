package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/zrma/1d1go/leetcode/problems/common"
)

func TestMergeTwoLists(t *testing.T) {
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
			l1:   []int{1, 2, 4},
			l2:   []int{1, 3, 4},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:   []int(nil),
			l2:   []int(nil),
			want: []int(nil),
		},
		{
			l1:   []int(nil),
			l2:   []int{1, 3, 4},
			want: []int{1, 3, 4},
		},
		{
			l1:   []int{1, 2, 4},
			l2:   []int(nil),
			want: []int{1, 2, 4},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l1 := build(tt.l1)
			l2 := build(tt.l2)
			want := build(tt.want)
			got := mergeTwoLists(l1, l2)
			assert.Equal(t, want.Traversal(), got.Traversal())
		})
	}
}
