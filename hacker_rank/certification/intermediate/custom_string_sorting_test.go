package intermediate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomSorting(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	for i, tt := range []struct {
		given []string
		want  []string
	}{
		{
			[]string{"abcd", "abc", "cd", "ab", "a", "cde", "b", "abcde"},
			[]string{"ab", "cd", "abcd", "abcde", "cde", "abc", "b", "a"},
		},
		{
			[]string{"abcd", "bcde"},
			[]string{"abcd", "bcde"},
		},
		{
			[]string{"abc", "bcd"},
			[]string{"bcd", "abc"},
		},
		{
			[]string{"abc", "bcde"},
			[]string{"bcde", "abc"},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := customSorting(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}
