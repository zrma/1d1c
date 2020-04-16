package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestStringToInteger(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			stringToInteger("3")
			stringToInteger("za")
		}, []string{
			"3",
			"Bad String",
		})
		assert.NoError(t, err)
	})
}
