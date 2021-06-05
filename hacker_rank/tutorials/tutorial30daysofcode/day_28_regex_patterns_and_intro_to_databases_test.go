package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestFilter(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-regex-patterns/problem")

	//noinspection SpellCheckingInspection
	want := []string{
		"julia",
		"julia",
		"riya",
		"samantha",
		"tanya",
	}
	//noinspection SpellCheckingInspection
	got, err := utils.GetPrinted(func() {
		filter([][]string{
			{"riya", "riya@gmail.com"},
			{"julia", "julia@julia.me"},
			{"julia", "sjulia@gmail.com"},
			{"julia", "julia@gmail.com"},
			{"samantha", "samantha@gmail.com"},
			{"tanya", "tanya@gmail.com"},
		})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
