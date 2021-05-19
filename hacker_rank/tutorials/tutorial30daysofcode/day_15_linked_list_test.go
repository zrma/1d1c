package tutorial30daysofcode

import (
	"strconv"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestDisplayLinkedList(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-linked-list/problem")

	var (
		given = []int{2, 3, 4, 1}
		want  = make([]string, len(given))
	)

	for i, n := range given {
		want[i] = strconv.Itoa(n)
	}

	err := utils.PrintTest(func() {
		displayLinkedList(given)
	}, want)
	assert.NoError(t, err)
}
