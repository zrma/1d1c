package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem", func(t *testing.T) {
		assert.False(t, isPrime(1))
		assert.True(t, isPrime(2))
		assert.True(t, isPrime(3))
		assert.False(t, isPrime(4))
		assert.True(t, isPrime(5))
		assert.True(t, isPrime(7))
		assert.False(t, isPrime(10))
		assert.False(t, isPrime(12))
		assert.True(t, isPrime(79))
		assert.True(t, isPrime(131))
		assert.False(t, isPrime(169))
		assert.True(t, isPrime(173))
		assert.False(t, isPrime(177))
	})
}
