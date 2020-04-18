package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperReducedString(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/reduced-string/problem", func(t *testing.T) {
		for input, expected := range map[string]string{
			"aaabccddd": "abd",
			"aa":        "Empty String",
			"baab":      "Empty String",
			"abab":      "abab",
			"abccba":    "Empty String",
			"":          "Empty String",
		} {
			actual := superReducedString(input)
			assert.Equal(t, expected, actual)
		}
	})
}
