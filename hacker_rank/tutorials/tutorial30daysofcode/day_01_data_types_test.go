package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestDayaType(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-data-types/problem")

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
}
