package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-linked-list-deletion/problem", func(t *testing.T) {
		t.Run("normal case", func(t *testing.T) {
			var list linkedList
			for _, num := range []int{1, 2, 2, 3, 3, 4} {
				list.head = list.insert(num)
			}

			list.head = removeDuplicates(list.head)
			err := utils.PrintTest(func() {
				list.display()
			}, []string{
				"1", "2", "3", "4",
			})
			assert.NoError(t, err)
		})

		t.Run("complex case", func(t *testing.T) {
			var list linkedList
			for _, num := range []int{1, 2, 2, 2, 3, 3, 3, 3, 4, 4} {
				list.head = list.insert(num)
			}

			list.head = removeDuplicates(list.head)
			err := utils.PrintTest(func() {
				list.display()
			}, []string{
				"1", "2", "3", "4",
			})
			assert.NoError(t, err)
		})
	})
}
