package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Log("https://leetcode.com/problems/palindrome-number/")

	for _, tt := range []struct {
		given int
		want  bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{1234, false},
		{1234321, true},
		{0, true},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			got := isPalindrome(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
