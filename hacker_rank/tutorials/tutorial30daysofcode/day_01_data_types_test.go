package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestDataType(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-data-types/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			dataType(
				4, 12,
				4.0, 4.0,
				"HackerRank ", "is the best place to learn and practice coding!",
			)
		}, []string{
			"16",
			"8.0",
			"HackerRank is the best place to learn and practice coding!",
		})
		assert.NoError(t, err)
	})
}
