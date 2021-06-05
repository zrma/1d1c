package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestDataType(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-data-types/problem")

	want := []string{
		"16",
		"8.0",
		"HackerRank is the best place to learn and practice coding!",
	}
	got, err := utils.GetPrinted(func() {
		dataType(
			4, 12,
			4.0, 4.0,
			"HackerRank ", "is the best place to learn and practice coding!",
		)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
