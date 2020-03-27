package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestFilter(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-regex-patterns/problem", func(t *testing.T) { //noinspection SpellCheckingInspection
		err := utils.PrintTest(func() {
			filter([][]string{
				{"riya", "riya@gmail.com"},
				{"julia", "julia@julia.me"},
				{"julia", "sjulia@gmail.com"},
				{"julia", "julia@gmail.com"},
				{"samantha", "samantha@gmail.com"},
				{"tanya", "tanya@gmail.com"},
			})
		}, []string{
			"julia",
			"julia",
			"riya",
			"samantha",
			"tanya",
		})
		assert.NoError(t, err)
	})
}
