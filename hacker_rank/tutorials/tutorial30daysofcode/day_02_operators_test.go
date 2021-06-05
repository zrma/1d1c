package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestOperators(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-operators/problem")

	want := []string{
		"15",
	}
	got, err := utils.GetPrinted(func() {
		operators(12.00, 20, 8)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
