package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem", func(t *testing.T) {
		assert.Equal(t, "NO", isValid("aabbcd"))
		assert.Equal(t, "NO", isValid("aabbccddeefghi"))
		assert.Equal(t, "NO", isValid("abbccc"))
		assert.Equal(t, "NO", isValid("aaaabbcc"))
		assert.Equal(t, "NO", isValid("aaaaabc"))
		assert.Equal(t, "YES", isValid("aabbc"))
		assert.Equal(t, "YES", isValid("aabbcc"))
		assert.Equal(t, "YES", isValid("abcdefghhgfedecba"))
	})
}
