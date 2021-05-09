package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-recursion/problem")

	for _, tt := range []struct {
		given int32
		want  int32
	}{
		{3, 6},
		{10, 3628800},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			got := factorial(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
