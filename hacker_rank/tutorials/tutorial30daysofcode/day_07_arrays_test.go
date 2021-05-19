package tutorial30daysofcode

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPrintReverse(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-arrays/problem")

	err := utils.PrintTest(func() {
		printReverse([]int32{1, 4, 3, 2})
	}, []string{
		"2 3 4 1",
	})
	assert.NoError(t, err)
}
