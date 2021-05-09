package tutorial30daysofcode

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNestedLogic(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem")

	for _, tt := range []struct {
		from, to string
		want     int
	}{
		{"1 1 1", "8 8 8", 0},
		{"6 6 2015", "9 6 2015", 0},
		{"9 6 2015", "6 6 2015", 45},
		{"6 7 2015", "6 6 2015", 500},
		{"6 6 2016", "6 6 2015", 10000},
	} {
		t.Run(fmt.Sprintf("%s %s", tt.from, tt.to), func(t *testing.T) {
			const layout = "2 1 2006"

			from, err := time.Parse(layout, convertCalendar(tt.from))
			assert.NoError(t, err)

			to, err := time.Parse(layout, convertCalendar(tt.to))
			assert.NoError(t, err)

			got := nestedLogic(from, to)
			assert.Equal(t, tt.want, got)
		})
	}
}
