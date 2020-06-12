package dp

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/play-game/problem", func() {
	Context("문제를 풀었다", func() {
		type testData struct {
			arr      []int32
			expected int64
		}

		DescribeTable("문제를 풀었다.", func(data testData) {
			actual := bricksGame(data.arr)
			Expect(actual).Should(Equal(data.expected))
		},
			Entry("0-0", testData{[]int32{1, 2, 3}, 6}),
			Entry("0-1", testData{[]int32{1, 2, 3, 4}, 6}),
			Entry("0-2", testData{[]int32{1, 2, 3, 4, 5}, 6}),
			Entry("1", testData{[]int32{999, 1, 1, 1, 0}, 1001}),
			Entry("2", testData{[]int32{0, 1, 1, 1, 999}, 999}),
			Entry("3", testData{[]int32{0, 1, 1, 1, 999, 999}, 1001}),
		)
	})

	Measure("성능 테스트", func(b Benchmarker) {
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

		runtime := b.Time("long string", func() {
			arr, err := readCSV("./test_data/bricks_game_0.csv")
			Expect(err).ShouldNot(HaveOccurred())

			arr2, err := readCSV("./test_data/bricks_game_1.csv")
			Expect(err).ShouldNot(HaveOccurred())

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

			Expect(len(arr)).Should(Equal(8))
			for i, arr := range arr {
				Expect(len(arr)).Should(BeNumerically("==", expected[i].l))
				actual := bricksGame(arr)
				Expect(actual).Should(BeNumerically("==", expected[i].e))
			}
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 5), "시간 초과")
	}, 3)
})
