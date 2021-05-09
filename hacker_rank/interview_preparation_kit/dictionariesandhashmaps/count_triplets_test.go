package dictionariesandhashmaps

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

func TestCountTriplets(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/count-triplets-1/problem")

	for i, tt := range []struct {
		arr   []int64
		ratio int64
		want  int64
	}{
		{[]int64{1, 2, 2, 4}, 2, 2},
		{[]int64{1, 3, 9, 9, 27, 81}, 3, 6},
		{[]int64{1, 5, 5, 25, 125}, 5, 4},
		{[]int64{1, 2, 4, 2, 4, 8}, 2, 6},
		{[]int64{1, 2, 4, 2, 4, 8, 2, 4, 2, 8}, 2, 15},
		{[]int64{1, 2, 1, 2, 1, 2}, 1, 2},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := countTriplets(tt.arr, tt.ratio)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountTripletsPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		file, err := os.Open("./test_data/count_triplets.csv")
		assert.NoError(t, err)
		defer func() {
			err := file.Close()
			assert.NoError(t, err)
		}()

		r := csv.NewReader(bufio.NewReader(file))
		rows, err := r.ReadAll()
		assert.NoError(t, err)

		row := rows[0]
		var ratio int64
		ratio, err = strconv.ParseInt(strings.TrimSpace(row[0]), 10, 64)
		assert.NoError(t, err)

		const wantRatio = 3
		assert.EqualValues(t, wantRatio, ratio)

		var arr []int64
		for _, col := range row[1:] {
			num, err := strconv.ParseInt(strings.TrimSpace(col), 10, 64)
			assert.NoError(t, err)

			arr = append(arr, num)
		}
		assert.Len(t, arr, 100000)
		const want = 2325652489

		got := countTriplets(arr, ratio)
		return assert.EqualValues(t, want, got)
	}, time.Second, time.Millisecond*100, "시간 초과")
}
