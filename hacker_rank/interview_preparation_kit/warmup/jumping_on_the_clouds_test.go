package warmup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpingOnClouds(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem")

	for i, tt := range []struct {
		given []int32
		want  int32
	}{
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
		{[]int32{0, 0, 0, 0, 0}, 2},
		{[]int32{0, 1, 0, 0, 1, 0}, 3},
		{[]int32{0, 1, 1, 0, 1, 0}, -1},
		{[]int32{0, 0, 1, 0, 1, 0}, 3},
		{[]int32{0, 0, 0, 1, 0, 0}, 3},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := jumpingOnClouds(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
