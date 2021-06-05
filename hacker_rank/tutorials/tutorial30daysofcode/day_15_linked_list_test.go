package tutorial30daysofcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestDisplayLinkedList(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-linked-list/problem")

	var given = []int{2, 3, 4, 1}
	want := make([]string, len(given))
	for i, n := range given {
		want[i] = strconv.Itoa(n)
	}
	got, err := utils.GetPrinted(func() {
		displayLinkedList(given)
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
