package dictionariesandhashmaps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoStrings(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/two-strings/problem")

	for _, tt := range []struct {
		s1, s2 string
		want   string
	}{
		{
			s1: "hello", s2: "world",
			want: "YES",
		},
		{
			s1: "hi", s2: "world",
			want: "NO",
		},
	} {
		t.Run(fmt.Sprintf("%s, %s", tt.s1, tt.s2), func(t *testing.T) {
			got := twoStrings(tt.s1, tt.s2)
			assert.Equal(t, tt.want, got)
		})
	}
}
