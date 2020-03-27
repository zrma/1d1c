package tutorial30daysofcode

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNestedLogic(t *testing.T) {
	t.Run("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem", func(t *testing.T) {
		t.Run("whole logic", func(t *testing.T) {
			const layout = "2 1 2006"

			for i, data := range []struct {
				from, to string
				result   int
			}{
				{"1 1 1", "8 8 8", 0},
				{"6 6 2015", "9 6 2015", 0},
				{"9 6 2015", "6 6 2015", 45},
				{"6 7 2015", "6 6 2015", 500},
				{"6 6 2016", "6 6 2015", 10000},
			} {
				from, err := time.Parse(layout, convertCalendar(data.from))
				assert.NoError(t, err, i)
				to, err := time.Parse(layout, convertCalendar(data.to))
				assert.NoError(t, err, i)

				actual := nestedLogic(from, to)
				assert.Equal(t, data.result, actual, i)
			}
		})

		t.Run("convertCalendar", func(t *testing.T) {
			for i, data := range []struct {
				input, expected string
			}{
				{"1 1 1", "1 1 0001"},
				{"8 8 8", "8 8 0008"},
				{"31 1 2009", "31 1 2009"},
				{"12 3 456", "12 3 0456"},
			} {
				actual := convertCalendar(data.input)
				assert.Equal(t, data.expected, actual, i)
			}
		})
	})
}
