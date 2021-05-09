package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestOperators(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-operators/problem")

	err := utils.PrintTest(func() {
		operators(12.00, 20, 8)
	}, []string{
		"15",
	})
	assert.NoError(t, err)
}
