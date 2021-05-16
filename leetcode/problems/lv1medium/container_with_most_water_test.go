package lv1medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	t.Log("https://leetcode.com/problems/container-with-most-water/")

	var (
		given = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		want  = 49
	)
	got := maxArea(given)
	assert.Equal(t, want, got)
}
