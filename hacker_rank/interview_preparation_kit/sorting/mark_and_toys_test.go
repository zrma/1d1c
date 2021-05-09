package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumToys(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/mark-and-toys/problem")

	for i, tt := range []struct {
		k      int32
		prices []int32
		want   int32
	}{
		{
			k:      7,
			prices: []int32{1, 2, 3, 4},
			want:   3,
		},
		{
			k:      50,
			prices: []int32{1, 12, 5, 111, 200, 1000, 10},
			want:   4,
		},
		{
			k:      15,
			prices: []int32{3, 7, 2, 9, 4},
			want:   3,
		},
		{
			k:      1,
			prices: []int32{1},
			want:   1,
		},
		{
			k:      0,
			prices: []int32{1, 2, 3},
			want:   0,
		},
		{
			k:      3,
			prices: []int32{1, 2, 1, 2},
			want:   2,
		},
		{
			k:      4,
			prices: []int32{1, 2, 1, 2},
			want:   3,
		},
		{
			k:      100,
			prices: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   10,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := maximumToys(tt.prices, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
