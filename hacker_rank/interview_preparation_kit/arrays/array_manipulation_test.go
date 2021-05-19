package arrays

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
)

func TestArrayManipulation(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/crush/problem")

	for i, tt := range []struct {
		n       int32
		queries [][]int32
		want    int64
	}{
		{
			n:       10,
			queries: [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}},
			want:    10,
		},
		{
			n:       5,
			queries: [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}},
			want:    200,
		},
		{
			n:       10,
			queries: [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}},
			want:    31,
		},
		{
			n:       10,
			queries: [][]int32{{3, 5, 7}, {2, 6, 8}, {1, 8, 1}, {5, 9, 15}},
			want:    31,
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := arrayManipulation(tt.n, tt.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestArrayManipulationPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		file, err := os.Open("./test_data/array_manipulation.csv")
		assert.NoError(t, err)
		defer func() {
			err := file.Close()
			assert.NoError(t, err)
		}()

		r := csv.NewReader(bufio.NewReader(file))
		rows, err := r.ReadAll()
		assert.NoError(t, err)

		var arr [][]int32
		for _, row := range rows {
			begin, err := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 64)
			assert.NoError(t, err)

			end, err := strconv.ParseInt(strings.TrimSpace(row[1]), 10, 64)
			assert.NoError(t, err)

			value, err := strconv.ParseInt(strings.TrimSpace(row[2]), 10, 64)
			assert.NoError(t, err)

			arr = append(arr, []int32{int32(begin), int32(end), int32(value)})
		}

		assert.Len(t, arr, 100000)
		const want = 2506721627

		got := arrayManipulation(10000000, arr)
		return assert.EqualValues(t, want, got)
	}, time.Second*5, time.Millisecond*100, "시간 초과")
}
