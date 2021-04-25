package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalX(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/between-two-sets/problem")

	got := getTotalX([]int32{2, 4}, []int32{16, 32, 96})
	assert.EqualValues(t, 3, got)

	got = getTotalX([]int32{3, 4}, []int32{24, 48})
	assert.EqualValues(t, 2, got)
}
