package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestLoop(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-loops/problem")

	want := []string{
		"2 x 1 = 2",
		"2 x 2 = 4",
		"2 x 3 = 6",
		"2 x 4 = 8",
		"2 x 5 = 10",
		"2 x 6 = 12",
		"2 x 7 = 14",
		"2 x 8 = 16",
		"2 x 9 = 18",
		"2 x 10 = 20",
	}
	got, err := utils.GetPrinted(func() {
		loop(2)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
