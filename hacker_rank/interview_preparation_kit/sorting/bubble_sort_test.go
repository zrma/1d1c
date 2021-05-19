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

	"1d1go/utils"
	"github.com/stretchr/testify/assert"
)

func TestCountSwaps(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/ctci-bubble-sort/problem")

	for i, tt := range []struct {
		given []int32
		want  []string
	}{
		{
			[]int32{6, 4, 1},
			[]string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 6",
			},
		},
		{
			[]int32{1, 2, 3},
			[]string{
				"Array is sorted in 0 swaps.",
				"First Element: 1",
				"Last Element: 3",
			},
		},
		{
			[]int32{3, 2, 1},
			[]string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 3",
			},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			err := utils.PrintTest(func() {
				countSwaps(tt.given)
			}, tt.want)
			assert.NoError(t, err)
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

		err = utils.PrintTest(func() {
			countSwaps(given)
		}, []string{
			"Array is sorted in 68472 swaps.",
			"First Element: 4842",
			"Last Element: 1994569",
		})
		return assert.NoError(t, err)
	}, time.Second, time.Millisecond*100, "시간 초과")
}
