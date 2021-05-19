package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestAbstractClasses(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-abstract-classes/problem")

	//noinspection SpellCheckingInspection
	const (
		title  = "The Alchemist"
		author = "Paulo Coelho"
		price  = 248
	)

	err := utils.PrintTest(func() {
		abstractClasses(title, author, price)
	}, []string{
		fmt.Sprintf("Title: %s", title),
		fmt.Sprintf("Author: %s", author),
		fmt.Sprintf("Price: %d", price),
	})
	assert.NoError(t, err)
}
