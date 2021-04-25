package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuperReducedString(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/reduced-string/problem", func(t *testing.T) {
		assert.Equal(t, "abd", superReducedString("aaabccddd"))
		assert.Equal(t, "Empty String", superReducedString("aa"))
		assert.Equal(t, "Empty String", superReducedString("baab"))
		assert.Equal(t, "abab", superReducedString("abab"))
		assert.Equal(t, "Empty String", superReducedString("abccba"))
		assert.Equal(t, "Empty String", superReducedString(""))
	})
}
