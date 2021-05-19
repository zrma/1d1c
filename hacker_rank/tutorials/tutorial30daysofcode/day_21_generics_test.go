package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPrintArray(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem")

	for i, tt := range []struct {
		given []interface{}
		want  []string
	}{
		{
			given: []interface{}{1, 2, 3},
			want:  []string{"1", "2", "3"},
		},
		{
			given: []interface{}{"Hello", "World"},
			want:  []string{"Hello", "World"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := utils.PrintTest(func() {
				printArray(tt.given...)
			}, tt.want)
			assert.NoError(t, err)
		})
	}
}
