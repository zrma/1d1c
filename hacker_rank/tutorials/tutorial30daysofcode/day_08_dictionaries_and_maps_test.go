package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zrma/going/utils"
)

func TestDictionariesAndMaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem")

	err := utils.PrintTest(func() {
		dictionariesAndMaps(3, []string{
			"sam 99912222",
			"tom 11122222",
			"harry 12299933",
			"sam",
			"edward",
			"harry",
		})
	}, []string{
		"sam=99912222",
		"Not found",
		"harry=12299933",
	})
	assert.NoError(t, err)
}
