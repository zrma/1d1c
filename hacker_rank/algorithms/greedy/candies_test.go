package greedy

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Done chan struct{}

func TestCandies(t *testing.T) {
	t.Parallel()

	t.Run("https://www.hackerrank.com/challenges/candies/problem", func(t *testing.T) {
		for _, data := range []struct {
			arr      []int32
			expected int64
		}{
			{[]int32{1, 2, 2}, 4},
			{[]int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}, 19},
			{[]int32{2, 4, 3, 5, 2, 6, 4, 5}, 12},
			{[]int32{1, 3, 3, 3, 2, 1}, 10},
			{[]int32{1, 3, 3, 3, 2, 2, 2, 3, 3, 3, 1}, 15},
		} {
			actual := candies(int32(len(data.arr)), data.arr)
			assert.Equal(t, actual, data.expected)
		}
	})

	t.Run("", func(t *testing.T) {
		timeout := time.After(3 * time.Second)
		waitCh := make(chan struct{})
		go func(done Done) {
			file, err := os.Open("./test_data/candies.csv")
			assert.NoError(t, err)
			defer file.Close()

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			assert.NoError(t, err)

			var arr []int32
			for _, col := range rows {
				num, err := strconv.ParseInt(col[0], 10, 32)
				assert.NoError(t, err)

				arr = append(arr, int32(num))
			}

			assert.Equal(t, len(arr), 100000)

			actual := candies(int32(len(arr)), arr)
			var expected int64 = 160929
			assert.Equal(t, actual, expected)

			close(done)
		}(waitCh)

		select {
		case <-timeout:
			t.Fatal("timeout")
		case <-waitCh:
		}
	})
}
