package lv0easy

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveFirstBadVersion(t *testing.T) {
	t.Log("https://leetcode.com/problems/first-bad-version/")

	for _, tt := range []struct {
		given int
		want  int
	}{
		{5, 4},
		{10, 3},
		{100, 37},
		{100, 38},
		{100, 49},
		{100, 50},
		{100, 51},
		{1024, 511},
		{1024, 512},
		{1024, 513},
		{1024, 1024},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			want := int(math.Ceil(math.Log2(float64(tt.given))))
			got, callCount := solveFirstBadVersion(tt.given, tt.want)
			assert.Equal(t, tt.want, got)
			assert.LessOrEqual(t, callCount, want, "성능 제한")
		})
	}
}
