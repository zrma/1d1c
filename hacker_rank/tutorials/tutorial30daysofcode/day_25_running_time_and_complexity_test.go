package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem", func(t *testing.T) {
		for _, data := range []struct {
			num int32
			ok  bool
		}{
			{num: 1, ok: false},
			{num: 2, ok: true},
			{num: 3, ok: true},
			{num: 4, ok: false},
			{num: 5, ok: true},
			{num: 7, ok: true},
			{num: 10, ok: false},
			{num: 12, ok: false},
			{num: 79, ok: true},
			{num: 131, ok: true},
			{num: 169, ok: false},
			{num: 173, ok: true},
			{num: 177, ok: false},
		} {
			if data.ok {
				assert.True(t, isPrime(data.num))
			} else {
				assert.False(t, isPrime(data.num))
			}
		}
	})
}
