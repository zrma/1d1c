package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseAnd(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-bitwise-and/problem")

	for _, tt := range []struct {
		n, k int32
		want int32
	}{
		{5, 2, 1},
		{8, 5, 4},
		{2, 2, 0},
		{955, 236, 235},
	} {
		t.Run(fmt.Sprintf("%d %d", tt.n, tt.k), func(t *testing.T) {
			got := bitwiseAND(tt.n, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
