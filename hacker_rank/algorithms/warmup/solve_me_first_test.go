package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMeFirst(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/solve-me-first/problem")

	var a uint32 = 2
	var b uint32 = 3
	got := solveMeFirst(a, b)
	const want = 5
	assert.EqualValues(t, want, got)
}
