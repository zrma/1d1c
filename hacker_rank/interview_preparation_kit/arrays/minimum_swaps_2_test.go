package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/minimum-swaps-2/problem")

	for i, tt := range []struct {
		given []int32
		want  int32
	}{
		{[]int32{4, 3, 1, 2}, 3},
		{[]int32{2, 3, 4, 1, 5}, 3},
		{[]int32{1, 3, 5, 2, 4, 6, 7}, 3},
		{[]int32{7, 1, 3, 2, 4, 5, 6}, 5},
		{[]int32{1, 3, 4, 5, 6}, -1},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := minimumSwaps(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFind(t *testing.T) {
	for i, tt := range []struct {
		description string
		given       int32
		wantOk      bool
		wantVal     int
	}{
		{
			description: "배열에서 주어진 오프셋 위치부터 타겟을 잘 찾는다.",
			given:       3,
			wantOk:      true,
			wantVal:     2,
		},
		{
			description: "오프셋 앞쪽의 타겟은 찾지 않는다.",
			given:       2,
			wantOk:      false,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			arr := []int32{1, 2, 3, 4, 5}
			length := len(arr)

			const offset = 2
			gotVal, gotOk := find(arr, tt.given, offset, length)
			assert.Equal(t, tt.wantOk, gotOk)

			if gotOk {
				assert.Equal(t, tt.wantVal, gotVal)
			}
		})
	}
}
