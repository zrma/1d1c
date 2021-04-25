package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestCountApplesAndOranges(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/apple-and-orange/problem")

	err := utils.PrintTest(func() {
		countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
	}, []string{
		"1",
		"1",
	})
	assert.NoError(t, err)
}
