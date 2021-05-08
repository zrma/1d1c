package codesprint2019

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphaBeta(t *testing.T) {
	t.Log("https://www.hackerrank.com/contests/hackerrank-all-womens-codesprint-2019/challenges/alpha-and-beta")

	for i, tt := range []struct {
		given []int32
		want  int32
	}{
		{[]int32{1, 2, 3, 3, 2, 1}, 2},
		{[]int32{1, 2, 3, 4, 5}, 4},
		{[]int32{1, 2, 3, 2, 1}, 2},
		{[]int32{1, 2, 4, 3, 2, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := alphaBeta(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
