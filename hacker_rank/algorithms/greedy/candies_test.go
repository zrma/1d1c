package greedy

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/extensions/table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/candies/problem", func() {
	type testData struct {
		arr      []int32
		expected int64
	}

	DescribeTable("문제를 풀었다",
		func(data testData) {

			actual := candies(int32(len(data.arr)), data.arr)
			Expect(actual).Should(BeNumerically("==", data.expected))
		},
		Entry("test_0", testData{[]int32{1, 2, 2}, 4}),
		Entry("test_1", testData{[]int32{2, 4, 2, 6, 1, 7, 8, 9, 2, 1}, 19}),
		Entry("test_2", testData{[]int32{2, 4, 3, 5, 2, 6, 4, 5}, 12}),
		Entry("test_3", testData{[]int32{1, 3, 3, 3, 2, 1}, 10}),
		Entry("test_4", testData{[]int32{1, 3, 3, 3, 2, 2, 2, 3, 3, 3, 1}, 15}),
	)

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long array", func() {
			file, err := os.Open("./test_data/candies.csv")
			Expect(err).ShouldNot(HaveOccurred())
			defer file.Close()

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			Expect(err).ShouldNot(HaveOccurred())

			var arr []int32
			for _, col := range rows {
				num, err := strconv.ParseInt(col[0], 10, 32)
				Expect(err).ShouldNot(HaveOccurred())

				arr = append(arr, int32(num))
			}

			Expect(arr).Should(HaveLen(100000))

			actual := candies(int32(len(arr)), arr)
			Expect(actual).Should(BeNumerically("==", 160929))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 10), "시간 초과")
	}, 3)
})
