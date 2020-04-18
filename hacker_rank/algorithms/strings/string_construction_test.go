package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConstruction(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/string-construction/problem", func(t *testing.T) {
		for input, expected := range map[string]int32{
			"abcd":   4,
			"abab":   2,
			"abcabc": 3,
		} {
			actual := stringConstruction(input)
			assert.Equal(t, expected, actual)
		}
	})
}
