package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperReducedString(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/reduced-string/problem")

	for _, tt := range []struct {
		given string
		want  string
	}{
		{"aaabccddd", "abd"},
		{"aa", "Empty String"},
		{"baab", "Empty String"},
		{"abab", "abab"},
		{"abccba", "Empty String"},
		{"", "Empty String"},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := superReducedString(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
