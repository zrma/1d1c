package sorting

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"1d1go/utils"
)

func TestCountSwaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/ctci-bubble-sort/problem")

	for i, tt := range []struct {
		given []int32
		want  []string
	}{
		{
			given: []int32{6, 4, 1},
			want: []string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 6",
			},
		},
		{
			given: []int32{1, 2, 3},
			want: []string{
				"Array is sorted in 0 swaps.",
				"First Element: 1",
				"Last Element: 3",
			},
		},
		{
			given: []int32{3, 2, 1},
			want: []string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 3",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, err := utils.GetPrinted(func() {
				countSwaps(tt.given)
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountSwapsPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		file, err := os.Open("./test_data/bubble_sort.csv")
		assert.NoError(t, err)
		defer func() {
			err := file.Close()
			assert.NoError(t, err)
		}()

		r := csv.NewReader(bufio.NewReader(file))
		rows, err := r.ReadAll()
		assert.NoError(t, err)

		var given []int32
		for _, row := range rows {
			num, err := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 32)
			assert.NoError(t, err)

			given = append(given, int32(num))
		}

		assert.Len(t, given, 528)

		want := []string{
			"Array is sorted in 68472 swaps.",
			"First Element: 4842",
			"Last Element: 1994569",
		}
		got, err := utils.GetPrinted(func() {
			countSwaps(given)
		})
		assert.NoError(t, err)
		return assert.Equal(t, want, got)
	}, time.Second, time.Millisecond*100, "시간 초과")
}
