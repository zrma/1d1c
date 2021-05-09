package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingValleys(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/counting-valleys/problem")

	for _, tt := range []struct {
		given string
		want  int32
	}{
		{"UDDDUDUU", 1},
		{"", 0},
		{"ABC", 0},
	} {
		t.Run(tt.given, func(t *testing.T) {
			got := countingValleys(int32(len(tt.given)), tt.given)
			assert.Equal(t, tt.want, got)
		})
	}

	assert.NotPanics(t, func() {
		const (
			given = "ABC"
			want  = 0
		)
		got := countingValleys(int32(len(given))+123, given)
		assert.EqualValues(t, want, got)
	})
}
