package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestClassAndInstance(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-class-vs-instance/problem")

	for _, tt := range []struct {
		given int
		want  []string
	}{
		{
			given: -1,
			want: []string{
				"Age is not valid, setting age to 0.",
				"You are young.",
				"You are young.",
				"",
			},
		},
		{
			given: 10,
			want: []string{
				"You are young.",
				"You are a teenager.",
				"",
			},
		},
		{
			given: 16,
			want: []string{
				"You are a teenager.",
				"You are old.",
				"",
			},
		},
		{
			given: 18,
			want: []string{
				"You are old.",
				"You are old.",
				"",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			err := utils.PrintTest(func() {
				classAndInstance(tt.given)
			}, tt.want)
			assert.NoError(t, err)
		})
	}
}
