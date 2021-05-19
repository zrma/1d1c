package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestBinaryNumbers(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-binary-numbers/problem")

	for _, tt := range []struct {
		given int32
		want  string
	}{
		{5, "1"},
		{13, "2"},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			err := utils.PrintTest(func() {
				binaryNumbers(tt.given)
			}, []string{tt.want})
			assert.NoError(t, err)
		})
	}
}
