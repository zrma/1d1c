package dictionariesandhashmaps

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/frequency-queries/problem", func() {
	It("문제를 풀었다", func() {
		Expect(freqQuery([][]int32{
			{1, 5},
			{1, 6},
			{3, 2},
			{1, 10},
			{1, 10},
			{1, 6},
			{2, 5},
			{3, 2},
		})).Should(Equal([]int32{0, 1}))
		Expect(freqQuery([][]int32{
			{3, 4},
			{2, 1003},
			{1, 16},
			{3, 1},
		})).Should(Equal([]int32{0, 1}))
		Expect(freqQuery([][]int32{
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
		})).Should(Equal([]int32{0, 1, 1}))
	})

	Measure("성능 테스트", func(b Benchmarker) {
		runtime := b.Time("long array", func() {
			arr, err := readInputCSV("./test_data/frequency_queries_0.csv")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(arr).Should(HaveLen(1000000))

			result, err := readResultCSV("./test_data/frequency_queries_result_0.csv")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(HaveLen(332855))

			diff := deep.Equal(freqQuery(arr), result)
			Expect(diff).Should(BeNil())

			arr, err = readInputCSV("./test_data/frequency_queries_1.csv")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(arr).Should(HaveLen(100000))

			result, err = readResultCSV("./test_data/frequency_queries_result_1.csv")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(HaveLen(33246))

			diff = deep.Equal(freqQuery(arr), result)
			Expect(diff).Should(BeNil())
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 5), "시간 초과")
	}, 1)
})

func readInputCSV(fileName string) ([][]int32, error) {
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

func readResultCSV(fileName string) ([]int32, error) {
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

	var arr []int32
	for _, row := range rows {
		num, err := strconv.ParseInt(row[0], 10, 32)
		if err != nil {
			return nil, err
		}
		arr = append(arr, int32(num))
	}
	return arr, nil
}
