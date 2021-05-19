package tutorial30daysofcode

import (
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-hello-world/problem")

	err := utils.PrintTest(func() {
		helloWorld("Welcome to 30 Days of Code!")
	}, []string{
		"Hello, World.",
		"Welcome to 30 Days of Code!",
	})
	assert.NoError(t, err)
}
