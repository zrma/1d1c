package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestPlusMinus(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/plus-minus/problem")

	want := []string{
		"0.500000",
		"0.333333",
		"0.166667",
	}
	got, err := utils.GetPrinted(func() {
		plusMinus([]int32{-4, 3, -9, 0, 4, 1})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
