package problems

import (
	"testing"

	"gotest.tools/assert"
)

func TestConvert(t *testing.T) {
	t.Run("https://leetcode.com/problems/zigzag-conversion/", func(t *testing.T) {
		t.Parallel()

		//noinspection SpellCheckingInspection
		const inputString = "PAYPALISHIRING"

		t.Run("case 1", func(t *testing.T) {
			//noinspection SpellCheckingInspection
			const expected = "PAHNAPLSIIGYIR"

			actual := convert(inputString, 3)
			assert.Equal(t, actual, expected)
		})

		t.Run("case 2", func(t *testing.T) {
			//noinspection SpellCheckingInspection
			const expected = "PINALSIGYAHRPI"

			actual := convert(inputString, 4)
			assert.Equal(t, actual, expected)
		})

		t.Run("it should not be panic", func(t *testing.T) {
			//noinspection SpellCheckingInspection
			const expected = "AB"

			actual := convert("AB", 1)
			assert.Equal(t, actual, expected)
		})
	})
}
