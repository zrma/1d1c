package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestStaircase(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/staircase/problem")

	want := []string{
		"     #",
		"    ##",
		"   ###",
		"  ####",
		" #####",
		"######",
	}
	got, err := utils.GetPrinted(func() {
		staircase(6)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
