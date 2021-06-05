package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestHelloWorld(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-hello-world/problem")

	want := []string{
		"Hello, World.",
		"Welcome to 30 Days of Code!",
	}
	got, err := utils.GetPrinted(func() {
		helloWorld("Welcome to 30 Days of Code!")
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
