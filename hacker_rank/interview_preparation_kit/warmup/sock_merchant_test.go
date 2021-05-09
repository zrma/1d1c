package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSockMerchant(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/sock-merchant/problem")

	given := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	const want = 3

	got := sockMerchant(given)
	assert.EqualValues(t, want, got)
}
