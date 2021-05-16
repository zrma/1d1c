package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	t.Log("https://leetcode.com/problems/search-insert-position/")

	for i, tt := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
