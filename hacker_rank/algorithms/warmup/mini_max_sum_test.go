package warmup

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestMiniMaxSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/mini-max-sum/problem")

	err := utils.PrintTest(func() {
		miniMaxSum([]int64{1, 2, 3, 4, 5})
	}, []string{
		"10 14",
	})
	assert.NoError(t, err)
}
