package tutorial30daysofcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem")

	for _, tt := range []struct {
		given int32
		want  bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{7, true},
		{10, false},
		{12, false},
		{79, true},
		{131, true},
		{169, false},
		{173, true},
		{177, false},
	} {
		t.Run(fmt.Sprintf("%d", tt.given), func(t *testing.T) {
			got := isPrime(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
