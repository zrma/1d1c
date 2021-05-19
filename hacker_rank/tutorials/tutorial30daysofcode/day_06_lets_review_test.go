package tutorial30daysofcode

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestLetsReview(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-loops/problem")

	err := utils.PrintTest(func() {
		letsReview("Hacker")
		letsReview("Rank")
	}, []string{
		"Hce akr",
		"Rn ak",
	})
	assert.NoError(t, err)
}
