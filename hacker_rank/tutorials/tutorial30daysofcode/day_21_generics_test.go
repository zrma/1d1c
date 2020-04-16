package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestPrintArray(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem", func(t *testing.T) {
		err := utils.PrintTest(func() {
			var arr []interface{}
			for _, data := range []int{1, 2, 3} {
				arr = append(arr, data)
			}
			printArray(arr...)

			arr = arr[:0]
			for _, data := range []string{"Hello", "World"} {
				arr = append(arr, data)
			}
			printArray(arr...)
		}, []string{
			"1",
			"2",
			"3",
			"Hello",
			"World",
		})
		assert.NoError(t, err)
	})
}
