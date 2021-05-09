package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorSum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-interfaces/problem")

	for _, tt := range []struct {
		given int
		want  int
	}{
		{1, 1},
		{6, 12},
		{20, 42},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			got := divisorSum(&calculator{}, tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
