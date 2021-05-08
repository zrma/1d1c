package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHourGlassSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/2d-array/problem")

	for i, tt := range []struct {
		given [][]int32
		want  int32
	}{
		{
			given: [][]int32{
				{-9, -9, -9, 1, 1, 1},
				{0, -9, 0, 4, 3, 2},
				{-9, -9, -9, 1, 2, 3},
				{0, 0, 8, 6, 6, 0},
				{0, 0, 0, -2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: 28,
		},
		{
			given: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			want: 19,
		},
		{},
		{
			given: [][]int32{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:  0,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.NotPanics(t, func() {
				got := HourGlassSum(tt.given)
				assert.Equal(t, tt.want, got)
			})
		})
	}
}
