package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestStringToInteger(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	want := []string{
		"3",
		"Bad String",
	}
	got, err := utils.GetPrinted(func() {
		stringToInteger("3")
		stringToInteger("za")
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
