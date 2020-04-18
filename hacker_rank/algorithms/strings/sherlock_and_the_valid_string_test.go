package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem", func(t *testing.T) {
		for input, expected := range map[string]string{
			"aabbcd":            "NO",
			"aabbccddeefghi":    "NO",
			"abbccc":            "NO",
			"aaaabbcc":          "NO",
			"aaaaabc":           "NO",
			"aabbc":             "YES",
			"aabbcc":            "YES",
			"abcdefghhgfedecba": "YES",
		} {
			actual := isValid(input)
			assert.Equal(t, expected, actual)
		}
	})
}
