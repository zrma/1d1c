package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestLetsReview(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-loops/problem")

	want := []string{
		"Hce akr",
		"Rn ak",
	}
	got, err := utils.GetPrinted(func() {
		letsReview("Hacker")
		letsReview("Rank")
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
