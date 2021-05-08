package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
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
