package problems

import (
	"testing"

	"gotest.tools/assert"
)

func TestMyAtoI(t *testing.T) {
	t.Run("https://leetcode.com/problems/string-to-integer-atoi/", func(t *testing.T) {
		for i, data := range []struct {
			input    string
			expected int
		}{
			{"42", 42},
			{"   -42", -42},
			{"4193 with words", 4193},
			{"words and 987", 0},
			{"-91283472332", -2147483648},
			{"", 0},
			{"3.14159", 3},
			{"  -0012a42", -12},
			{"          ", 0},
			{"          -", 0},
			{"+1", 1},
			{"9223372036854775808", 2147483647},
			{"18446744073709551617", 2147483647},
			{"  0000000000012345678", 12345678},
			{"-000000000000001", -1},
		} {
			actual := myAtoI(data.input)
			assert.Equal(t, actual, data.expected, "test case: %d", i)
		}
	})
}
