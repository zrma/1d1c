package warmup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/time-conversion/problem")

	for _, tt := range []struct {
		given string
		want  string
	}{
		{"07:05:45PM", "19:05:45"},
		{"AB:12:34AM", "00:12:34"},
	} {
		t.Run(tt.given, func(t *testing.T) {
			assert.NotPanics(t, func() {
				got := timeConversion(tt.given)
				assert.Equal(t, tt.want, got)
			})
		})
	}
}
