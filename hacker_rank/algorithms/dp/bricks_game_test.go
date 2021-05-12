package dp

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBricksGame(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/play-game/problem")

	for _, tc := range []struct {
		description string
		given       []int32
		want        int64
	}{
		{
			description: "1~3",
			given:       []int32{1, 2, 3},
			want:        6,
		},
		{
			description: "1~4",
			given:       []int32{1, 2, 3, 4},
			want:        6,
		},
		{
			description: "1~5",
			given:       []int32{1, 2, 3, 4, 5},
			want:        6,
		},
		{
			description: "앞에 큰 수, 뒤에 0",
			given:       []int32{999, 1, 1, 1, 0},
			want:        1001,
		},
		{
			description: "앞에 0, 뒤에 큰 수",
			given:       []int32{0, 1, 1, 1, 999},
			want:        999,
		},
		{
			description: "여러 중복 값의 처리",
			given:       []int32{0, 1, 1, 1, 999, 999},
			want:        1001,
		},
	} {
		t.Run(tc.description, func(t *testing.T) {
			got := bricksGame(tc.given)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBricksGamePerformance(t *testing.T) {
	readCSV := func(fileName string) ([][]int32, error) {
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer func() {
			err := file.Close()
			assert.NoError(t, err)
		}()

		r := csv.NewReader(bufio.NewReader(file))
		rows, err := r.ReadAll()
		if err != nil {
			return nil, err
		}

		var arr [][]int32
		for _, row := range rows {
			var cols []int32
			for _, col := range row {
				num, err := strconv.ParseInt(strings.TrimSpace(col), 10, 32)
				if err != nil {
					return nil, err

				}
				cols = append(cols, int32(num))
			}
			arr = append(arr, cols)
		}
		return arr, nil
	}

	assert.Eventually(t, func() bool {
		arr, err := readCSV("./test_data/bricks_game_0.csv")
		assert.NoError(t, err)

		arr2, err := readCSV("./test_data/bricks_game_1.csv")
		assert.NoError(t, err)

		arr = append(arr, arr2...)

		want := []struct {
			l int
			e int64
		}{
			{1000, 249147},
			{1000, 251633},
			{1000, 249302},
			{1000, 247737},
			{1000, 253105},
			{100000, 249791261588},
			{100000, 249894676936},
			{100000, 250224672758},
		}
		assert.Len(t, arr, 8)

		for i, arr := range arr {
			assert.Len(t, arr, want[i].l)

			got := bricksGame(arr)
			assert.EqualValues(t, want[i].e, got)
		}
		return true
	}, time.Second, time.Millisecond*100, "시간 초과")
}
