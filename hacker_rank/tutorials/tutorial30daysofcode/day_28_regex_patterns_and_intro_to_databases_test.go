package tutorial30daysofcode

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-regex-patterns/problem")

	//noinspection SpellCheckingInspection
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
}
