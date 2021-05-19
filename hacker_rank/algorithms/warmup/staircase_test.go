package warmup

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestStaircase(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/staircase/problem")

	err := utils.PrintTest(func() {
		staircase(6)
	}, []string{
		"     #",
		"    ##",
		"   ###",
		"  ####",
		" #####",
		"######",
	})
	assert.NoError(t, err)
}
