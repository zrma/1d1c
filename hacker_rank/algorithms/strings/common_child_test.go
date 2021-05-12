package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonChild(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/common-child/problem")

	//noinspection SpellCheckingInspection
	got := commonChild("ABCD", "ABDC")
	assert.EqualValues(t, 3, got)

	got = commonChild("HARRY", "SALLY")
	assert.EqualValues(t, 2, got)

	got = commonChild("AA", "BB")
	assert.EqualValues(t, 0, got)

	//noinspection SpellCheckingInspection
	got = commonChild("SHINCHAN", "NOHARAAA")
	assert.EqualValues(t, 3, got)

	//noinspection SpellCheckingInspection
	got = commonChild("ABCDEF", "FBDAMN")
	assert.EqualValues(t, 2, got)
}
