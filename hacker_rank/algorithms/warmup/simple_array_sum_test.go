package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleArraySum(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/simple-array-sum/problem")

	given := []int32{1, 2, 3, 4, 10, 11}
	got := simpleArraySum(given)
	const want = 31
	assert.EqualValues(t, want, got)
}
