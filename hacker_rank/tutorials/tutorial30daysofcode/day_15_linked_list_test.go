package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestDisplayLinkedList(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-linked-list/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			displayLinkedList([]int{2, 3, 4, 1})
		}, []string{
			"2",
			"3",
			"4",
			"1",
		})
		assert.NoError(t, err)
	})
}
