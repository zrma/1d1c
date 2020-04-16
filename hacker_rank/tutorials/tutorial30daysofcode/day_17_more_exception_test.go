package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestMoreException(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-more-exceptions/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			moreException(3, 5)
			moreException(2, 4)
			moreException(-1, -2)
			moreException(-1, 3)
		}, []string{
			"243",
			"16",
			"n and p should be non-negative",
			"n and p should be non-negative",
		})
		assert.NoError(t, err)
	})
}
