package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Log("https://leetcode.com/problems/remove-duplicates-from-sorted-array/")

	for i, tt := range []struct {
		given []int
		want  []int
	}{
		{
			given: []int{1, 1, 2},
			want:  []int{1, 2},
		},
		{
			given: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:  []int{0, 1, 2, 3, 4},
		},
		{
			given: []int{0, 0, 3, 3, 3, 3, 3, 4, 5, 5},
			want:  []int{0, 3, 4, 5},
		},
		{
			given: nil,
			want:  nil,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := removeDuplicates(tt.given)
			wantLength := len(tt.want)
			assert.Equal(t, wantLength, got)
			assert.Equal(t, tt.want, tt.given[:wantLength])
		})
	}
}
