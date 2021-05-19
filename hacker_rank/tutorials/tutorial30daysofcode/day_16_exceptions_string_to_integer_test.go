package tutorial30daysofcode

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestStringToInteger(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	err := utils.PrintTest(func() {
		stringToInteger("3")
		stringToInteger("za")
	}, []string{
		"3",
		"Bad String",
	})
	assert.NoError(t, err)
}
