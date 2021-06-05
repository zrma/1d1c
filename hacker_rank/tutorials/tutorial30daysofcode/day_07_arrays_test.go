package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestPrintReverse(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-arrays/problem")

	want := []string{
		"2 3 4 1",
	}
	got, err := utils.GetPrinted(func() {
		printReverse([]int32{1, 4, 3, 2})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
