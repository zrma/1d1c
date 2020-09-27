package sorting

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/ctci-bubble-sort/problem", func() {
	It("문제를 풀었다.", func() {
		{
			err := utils.PrintTest(func() {
				countSwaps([]int32{6, 4, 1})
			}, []string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 6",
			})
			Expect(err).ShouldNot(HaveOccurred())
		}

		{
			err := utils.PrintTest(func() {
				countSwaps([]int32{1, 2, 3})
			}, []string{
				"Array is sorted in 0 swaps.",
				"First Element: 1",
				"Last Element: 3",
			})
			Expect(err).ShouldNot(HaveOccurred())
		}

		{
			err := utils.PrintTest(func() {
				countSwaps([]int32{3, 2, 1})
			}, []string{
				"Array is sorted in 3 swaps.",
				"First Element: 1",
				"Last Element: 3",
			})
			Expect(err).ShouldNot(HaveOccurred())
		}
	})

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long array", func() {
			file, err := os.Open("./test_data/bubble_sort.csv")
			Expect(err).ShouldNot(HaveOccurred())
			defer file.Close()

			r := csv.NewReader(bufio.NewReader(file))
			rows, err := r.ReadAll()
			Expect(err).ShouldNot(HaveOccurred())

			var arr []int32
			var num int64
			for _, row := range rows {
				num, err = strconv.ParseInt(row[0], 10, 32)
				Expect(err).ShouldNot(HaveOccurred())

				arr = append(arr, int32(num))
			}

			Expect(arr).Should(HaveLen(528))
			err = utils.PrintTest(func() {
				countSwaps(arr)
			}, []string{
				"Array is sorted in 68472 swaps.",
				"First Element: 4842",
				"Last Element: 1994569",
			})
			Expect(err).ShouldNot(HaveOccurred())
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 10), "시간 초과")
	}, 3)
})
