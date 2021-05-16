package lv1medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log("https://leetcode.com/problems/longest-substring-without-repeating-characters/")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		given string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"abba", 2},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
