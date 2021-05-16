package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	t.Log("https://leetcode.com/problems/remove-element/")

	for i, tt := range []struct {
		nums []int
		val  int
		want []int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: []int{2, 2},
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: []int{0, 1, 3, 0, 4},
		},
		{
			nums: []int{0, 0, 3, 3, 3, 3, 3, 4, 5, 5},
			val:  3,
			want: []int{0, 0, 4, 5, 5},
		},
		{
			nums: nil,
			val:  3,
			want: nil,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			wantLength := len(tt.want)
			assert.Equal(t, wantLength, got)
			assert.Equal(t, tt.want, tt.nums[:wantLength])
		})
	}
}
