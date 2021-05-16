package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	t.Log("https://leetcode.com/problems/implement-strstr/")

	//noinspection SpellCheckingInspection
	for _, tt := range []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaaaa", "", 0},
		{"", "", 0},
		{"", "abc", -1},
	} {
		t.Run(fmt.Sprintf("%s %s", tt.haystack, tt.needle), func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			assert.Equal(t, tt.want, got)
		})
	}
}
