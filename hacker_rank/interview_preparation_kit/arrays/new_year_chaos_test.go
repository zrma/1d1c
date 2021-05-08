package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestMinimumBribes(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/new-year-chaos/problem")

	for i, tt := range []struct {
		given []int32
		want  string
	}{
		{[]int32{2, 1, 5, 3, 4}, "3"},
		{[]int32{2, 5, 1, 3, 4}, "Too chaotic"},
		{[]int32{5, 1, 2, 3, 7, 8, 6, 4}, "Too chaotic"},
		{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, "7"},
		{[]int32{1, 2, 5, 3, 4, 7, 8, 6}, "4"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := utils.PrintTest(func() {
				minimumBribes(tt.given)
			}, []string{
				tt.want,
			})
			assert.NoError(t, err)
		})
	}
}
