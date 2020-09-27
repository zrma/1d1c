package dictionariesandhashmaps

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/count-triplets-1/problem", func() {
	type testData struct {
		arr      []int64
		ratio    int64
		expected int64
	}

	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := countTriplets(data.arr, data.ratio)
			Expect(actual).Should(BeNumerically("==", data.expected))
		},
		Entry("test_0", testData{[]int64{1, 2, 2, 4}, 2, 2}),
		Entry("test_1", testData{[]int64{1, 3, 9, 9, 27, 81}, 3, 6}),
		Entry("test_2", testData{[]int64{1, 5, 5, 25, 125}, 5, 4}),
		Entry("test_3", testData{[]int64{1, 2, 4, 2, 4, 8}, 2, 6}),
		Entry("test_4", testData{[]int64{1, 2, 4, 2, 4, 8, 2, 4, 2, 8}, 2, 15}),
		Entry("test_5", testData{[]int64{1, 2, 1, 2, 1, 2}, 1, 2}),
	)

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long array", func() {
			file, err := os.Open("./test_data/count_triplets.csv")
			Expect(err).ShouldNot(HaveOccurred())
			defer file.Close()

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			Expect(err).ShouldNot(HaveOccurred())

			row := rows[0]
			var ratio int64
			ratio, err = strconv.ParseInt(row[0], 10, 64)
			Expect(err).ShouldNot(HaveOccurred())

			var arr []int64
			for _, col := range row[1:] {
				num, err := strconv.ParseInt(col, 10, 64)
				Expect(err).ShouldNot(HaveOccurred())

				arr = append(arr, num)
			}

			Expect(ratio).Should(BeNumerically("==", 3))
			Expect(arr).Should(HaveLen(100000))
			actual := countTriplets(arr, ratio)
			Expect(actual).Should(BeNumerically("==", 2325652489))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 10), "시간 초과")
	}, 3)
})
