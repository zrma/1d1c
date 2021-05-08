package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestPlusMinus(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/plus-minus/problem")

	err := utils.PrintTest(func() {
		plusMinus([]int32{-4, 3, -9, 0, 4, 1})
	}, []string{
		"0.500000",
		"0.333333",
		"0.166667",
	})
	assert.NoError(t, err)
}
