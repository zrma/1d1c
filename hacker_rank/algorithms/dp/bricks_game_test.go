package dp

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/zrma/going/utils"
)

func TestBricksGame(t *testing.T) {
	t.Parallel()

	t.Run("https://www.hackerrank.com/challenges/play-game/problem", func(t *testing.T) {
		for _, data := range []struct {
			arr      []int32
			expected int64
		}{
			{arr: []int32{1, 2, 3}, expected: 6},
			{arr: []int32{1, 2, 3, 4}, expected: 6},
			{arr: []int32{1, 2, 3, 4, 5}, expected: 6},
			{arr: []int32{999, 1, 1, 1, 0}, expected: 1001},
			{arr: []int32{0, 1, 1, 1, 999}, expected: 999},
			{arr: []int32{0, 1, 1, 1, 999, 999}, expected: 1001},
		} {
			actual := bricksGame(data.arr)
			assert.Equal(t, actual, data.expected)
		}
	})

	t.Run("performance measure", func(t *testing.T) {
		RunUntil(t, func(done Done) {
			defer close(done)

			readCSV := func(fileName string) ([][]int32, error) {
				file, err := os.Open(fileName)
				if err != nil {
					return nil, err
				}
				defer file.Close()

				r := csv.NewReader(bufio.NewReader(file))
				rows, err := r.ReadAll()
				if err != nil {
					return nil, err
				}

				var arr [][]int32
				for _, row := range rows {
					var cols []int32
					for _, col := range row {
						num, err := strconv.ParseInt(col, 10, 32)
						if err != nil {
							return nil, err
						}
						cols = append(cols, int32(num))
					}
					arr = append(arr, cols)
				}
				return arr, nil
			}

			arr, err := readCSV("./test_data/bricks_game_0.csv")
			assert.NoError(t, err)

			arr2, err := readCSV("./test_data/bricks_game_1.csv")
			assert.NoError(t, err)

			arr = append(arr, arr2...)

			expected := []struct{ l, e int64 }{
				{1000, 249147},
				{1000, 251633},
				{1000, 249302},
				{1000, 247737},
				{1000, 253105},
				{100000, 249791261588},
				{100000, 249894676936},
				{100000, 250224672758},
			}

			assert.Equal(t, len(arr), 8)
			for i, arr := range arr {
				assert.Equal(t, int64(len(arr)), expected[i].l)
				actual := bricksGame(arr)
				assert.Equal(t, actual, expected[i].e)
			}
		}, 3)
	})
}
