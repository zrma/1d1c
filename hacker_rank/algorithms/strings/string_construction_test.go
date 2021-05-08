package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConstruction(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/string-construction/problem")

	for _, tt := range []struct {
		given string
		want  int32
	}{
		{"abcd", 4},
		{"abab", 2},
		{"abcabc", 3},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := stringConstruction(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
