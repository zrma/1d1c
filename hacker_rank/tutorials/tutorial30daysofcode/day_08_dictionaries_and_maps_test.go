package tutorial30daysofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestDictionariesAndMaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem")

	want := []string{
		"sam=99912222",
		"Not found",
		"harry=12299933",
	}
	got, err := utils.GetPrinted(func() {
		dictionariesAndMaps(3, []string{
			"sam 99912222",
			"tom 11122222",
			"harry 12299933",
			"sam",
			"edward",
			"harry",
		})
	})
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
