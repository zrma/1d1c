package lv1medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Log("https://leetcode.com/problems/longest-palindromic-substring/")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		given string
		want  string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"cbbc", "cbbc"},
		{"abcd", "a"},
		{"", ""},
		{"a", "a"},
		{"abcdefedcba", "abcdefedcba"},
		{"aabcdefedcba", "abcdefedcba"},
		{"abcdefedcbaa", "abcdefedcba"},
		{"aaabcdefedcba", "abcdefedcba"},
		{"abcdefedcbaaa", "abcdefedcba"},
		{"bb", "bb"},
		{"aba", "aba"},
		{"abb", "bb"},
		{"aab", "aa"},
		{"abbb", "bbb"},
		{"aabbb", "bbb"},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := longestPalindrome(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
