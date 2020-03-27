package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestBinaryNumbers(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-binary-numbers/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			binaryNumbers(5)
			binaryNumbers(13)
		}, []string{
			"1",
			"2",
		})
		assert.NoError(t, err)
	})
}
