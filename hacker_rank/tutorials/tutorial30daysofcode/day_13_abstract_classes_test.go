package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

//noinspection SpellCheckingInspection
func TestAbstractClasses(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-abstract-classes/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			abstractClasses("The Alchemist", "Paulo Coelho", 248)
		}, []string{
			"Title: The Alchemist",
			"Author: Paulo Coelho",
			"Price: 248",
		})
		assert.NoError(t, err)
	})
}
