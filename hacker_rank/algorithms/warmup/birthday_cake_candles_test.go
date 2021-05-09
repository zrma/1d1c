package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthdayCakeCandles(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/birthday-cake-candles/problem")

	given := []int32{3, 2, 1, 3}
	const want = 2

	got := birthdayCakeCandles(given)
	assert.EqualValues(t, want, got)
}
