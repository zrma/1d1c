package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMeFirst(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/solve-me-first/problem")

	var a uint32 = 2
	var b uint32 = 3
	const want = 5

	got := solveMeFirst(a, b)
	assert.EqualValues(t, want, got)
}
