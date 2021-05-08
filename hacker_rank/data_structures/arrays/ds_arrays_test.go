package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseArray(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/arrays-ds/problem")

	given := []int32{1, 4, 3, 2}
	got := reverseArray(given)
	want := []int32{2, 3, 4, 1}
	assert.Equal(t, want, got)
}
