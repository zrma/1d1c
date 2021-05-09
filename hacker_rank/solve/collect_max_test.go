package solve

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCollectMax(t *testing.T) {
	t.Log("collect max diamond problem")

	for i, tt := range []struct {
		given [][]int
		want  int
	}{
		{
			given: [][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1},
			},
			want: 5,
		},
		{
			given: [][]int{
				{0, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: 7,
		},
		{
			given: [][]int{
				{0, 1, 1},
				{1, 0, -1},
				{1, 1, -1},
			},
			want: 0,
		},
		{
			given: [][]int{{}, {}},
			want:  0,
		},
		{
			given: [][]int{},
			want:  0,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := collectMax(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCombinations(t *testing.T) {
	for _, tt := range []struct {
		tot, num int
		want     int
	}{
		{3, 2, 3},
		{4, 2, 6},
		{4, 3, 4},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.tot, tt.num), func(t *testing.T) {
			got := combinations(tt.tot, tt.num)
			assert.Len(t, got, tt.want)
		})
	}
}

func TestArrayFind(t *testing.T) {
	for i, tt := range []struct {
		arr  []int
		n    int
		want bool
	}{
		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{}, 1, false},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := array(tt.arr).find(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCollectMaxPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		data := [][]int{
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 2, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
		}
		const want = 23

		got := collectMax(data)
		return assert.Equal(t, want, got)
	}, time.Second*10, time.Millisecond*100, "시간 초과")
}
