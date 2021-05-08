package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotLeft(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/array-left-rotation/problem")

	for i, tt := range []struct {
		given int32
		want  []int32
	}{
		{0, []int32{1, 2, 3, 4, 5}},
		{1, []int32{2, 3, 4, 5, 1}},
		{2, []int32{3, 4, 5, 1, 2}},
		{3, []int32{4, 5, 1, 2, 3}},
		{4, []int32{5, 1, 2, 3, 4}},
		{5, []int32{1, 2, 3, 4, 5}},
		{6, []int32{2, 3, 4, 5, 1}},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			arr := []int32{1, 2, 3, 4, 5}

			got := rotLeft(arr, tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
