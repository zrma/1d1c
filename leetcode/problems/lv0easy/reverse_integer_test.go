package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Log("https://leetcode.com/problems/reverse-integer/")

	for _, tt := range []struct {
		given int
		want  int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			got := reverse(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
