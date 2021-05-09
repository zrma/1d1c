package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDifference(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-scope/problem")

	diff := newDifference([]int{1, 2, 5})
	const want = 4

	got := diff.computeDifference()
	assert.Equal(t, want, got)
}
