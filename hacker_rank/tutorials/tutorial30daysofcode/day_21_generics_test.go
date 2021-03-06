package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
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
			got, err := utils.GetPrinted(func() {
				printArray(tt.given...)
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
