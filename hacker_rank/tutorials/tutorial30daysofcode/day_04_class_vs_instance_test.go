package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestClassAndInstance(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-class-vs-instance/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			classAndInstance(-1)
			classAndInstance(10)
			classAndInstance(16)
			classAndInstance(18)
		}, []string{
			"Age is not valid, setting age to 0.",
			"You are young.",
			"You are young.",
			"",
			"You are young.",
			"You are a teenager.",
			"",
			"You are a teenager.",
			"You are old.",
			"",
			"You are old.",
			"You are old.",
			"",
		})
		assert.NoError(t, err)
	})
}
