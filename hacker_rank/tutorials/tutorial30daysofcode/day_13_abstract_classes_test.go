package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestAbstractClasses(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-abstract-classes/problem")

	//noinspection SpellCheckingInspection
	const (
		title  = "The Alchemist"
		author = "Paulo Coelho"
		price  = 248
	)
	want := []string{
		fmt.Sprintf("Title: %s", title),
		fmt.Sprintf("Author: %s", author),
		fmt.Sprintf("Price: %d", price),
	}
	got, err := utils.GetPrinted(func() {
		abstractClasses(title, author, price)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
