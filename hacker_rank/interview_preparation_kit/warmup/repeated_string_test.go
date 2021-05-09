package warmup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedString(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/repeated-string/problem")

	for _, tt := range []struct {
		s    string
		n    int64
		want int64
	}{
		{"aba", 10, 7},
		{"a", 1000000000000, 1000000000000},
		{"afternoon after", 3, 1},
		{"school hometown", 5, 0},
		{"", 5, 0},
	} {
		t.Run(fmt.Sprintf("%s %d", tt.s, tt.n), func(t *testing.T) {
			got := repeatedString(tt.s, tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
