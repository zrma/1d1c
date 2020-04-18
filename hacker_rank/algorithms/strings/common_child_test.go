package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChild(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("https://www.hackerrank.com/challenges/common-child/problem", func(t *testing.T) {
		for _, data := range []struct {
			s1, s2   string
			expected int32
		}{
			{"ABCD", "ABDC", 3},
			{"HARRY", "SALLY", 2},
			{"AA", "BB", 0},
			{"SHINCHAN", "NOHARAAA", 3},
			{"ABCDEF", "FBDAMN", 2},
		} {
			actual := commonChild(data.s1, data.s2)
			assert.Equal(t, data.expected, actual)
		}
	})
}
