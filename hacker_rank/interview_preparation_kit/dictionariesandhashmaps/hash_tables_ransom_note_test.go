package dictionariesandhashmaps

import (
	"fmt"
	"testing"

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestCheckMagazine(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/ctci-ransom-note/problem")

	for i, tt := range []struct {
		magazine, note []string
		want           string
	}{
		{
			magazine: []string{"give", "me", "one", "grand", "today", "night"},
			note:     []string{"give", "one", "grand", "today"},
			want:     "Yes",
		},
		{
			magazine: []string{"two", "times", "three", "is", "not", "four"},
			note:     []string{"two", "times", "two", "is", "four"},
			want:     "No",
		},
		{
			magazine: []string{"ive", "got", "a", "lovely", "brunch", "of", "coconuts"},
			note:     []string{"ive", "got", "some", "coconuts"},
			want:     "No",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := utils.PrintTest(func() {
				checkMagazine(tt.magazine, tt.note)
			}, []string{tt.want})
			assert.NoError(t, err)
		})
	}
}
