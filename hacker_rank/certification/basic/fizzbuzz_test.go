package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestFizzBuzz(t *testing.T) {
	want := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	got, err := utils.GetPrinted(func() {
		fizzBuzz(15)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
