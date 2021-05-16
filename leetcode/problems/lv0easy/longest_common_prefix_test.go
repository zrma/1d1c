package lv0easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Log("https://leetcode.com/problems/longest-common-prefix/")

	//noinspection SpellCheckingInspection
	for i, tt := range []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"flower", "flower", "flower"}, "flower"},
		{[]string{"flowers", "flower", "flower"}, "flower"},
		{[]string{}, ""},
		{nil, ""},
		{[]string{"", "flower", "flower"}, ""},
		{[]string{"c", "c"}, "c"},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := longestCommonPrefix(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
