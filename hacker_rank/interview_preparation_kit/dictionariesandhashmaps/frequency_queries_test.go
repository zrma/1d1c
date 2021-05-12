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

func TestFreqQuery(t *testing.T) {
	t.Log("https://www.hackerrank.com/challenges/frequency-queries/problem")

	for i, tt := range []struct {
		given [][]int32
		want  []int32
	}{
		{
			[][]int32{
				{1, 5},
				{1, 6},
				{3, 2},
				{1, 10},
				{1, 10},
				{1, 6},
				{2, 5},
				{3, 2},
			},
			[]int32{0, 1},
		},
		{
			[][]int32{
				{3, 4},
				{2, 1003},
				{1, 16},
				{3, 1},
			},
			[]int32{0, 1},
		},
		{
			[][]int32{
				{1, 3},
				{2, 3},
				{3, 2},
				{1, 4},
				{1, 5},
				{1, 5},
				{1, 4},
				{3, 2},
				{2, 4},
				{3, 2},
			},
			[]int32{0, 1, 1},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := freqQuery(tt.given)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFreqQueryPerformance(t *testing.T) {
	assert.Eventually(t, func() bool {
		for _, tt := range []struct {
			givenFile string
			givenLen  int
			wantFile  string
			wantLen   int
		}{
			{
				givenFile: "./test_data/frequency_queries_0.csv",
				givenLen:  1000000,
				wantFile:  "./test_data/frequency_queries_result_0.csv",
				wantLen:   332855,
			},
			{
				givenFile: "./test_data/frequency_queries_1.csv",
				givenLen:  100000,
				wantFile:  "./test_data/frequency_queries_result_1.csv",
				wantLen:   33246,
			},
		} {
			t.Run(tt.givenFile, func(t *testing.T) {
				given, err := readInputCSV(tt.givenFile)
				assert.NoError(t, err)
				assert.Len(t, given, tt.givenLen)

				want, err := readResultCSV(tt.wantFile)
				assert.NoError(t, err)
				assert.Len(t, want, tt.wantLen)

				got := freqQuery(given)
				assert.Equal(t, want, got)
			})
		}
		return true
	}, time.Second*10, time.Millisecond*100, "시간 초과")
}

func readInputCSV(fileName string) (arr [][]int32, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

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

func readResultCSV(fileName string) (arr []int32, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		num, err := strconv.ParseInt(strings.TrimSpace(row[0]), 10, 32)
		if err != nil {
			return nil, err
		}
		arr = append(arr, int32(num))
	}
	return arr, nil
}
