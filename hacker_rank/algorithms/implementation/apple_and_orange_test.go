package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestCountApplesAndOranges(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/apple-and-orange/problem")

	want := []string{
		"1",
		"1",
	}
	got, err := utils.GetPrinted(func() {
		countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
