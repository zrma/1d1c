package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestLetsReview(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-loops/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			letsReview("Hacker")
			letsReview("Rank")
		}, []string{
			"Hce akr",
			"Rn ak",
		})
		assert.NoError(t, err)
	})
}
