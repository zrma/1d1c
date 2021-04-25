package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConstruction(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/string-construction/problem", func(t *testing.T) {
		assert.EqualValues(t, 4, stringConstruction("abcd"))
		assert.EqualValues(t, 2, stringConstruction("abab"))
		assert.EqualValues(t, 3, stringConstruction("abcabc"))
	})
}
