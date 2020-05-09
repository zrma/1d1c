package sorting

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/mark-and-toys/problem", func() {
	type testData struct {
		k        int32
		prices   []int32
		expected int32
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := maximumToys(data.prices, data.k)
		expected := data.expected
		Expect(actual).Should(BeNumerically("==", expected))
	},
		Entry("case 0", testData{
			k:        7,
			prices:   []int32{1, 2, 3, 4},
			expected: 3,
		}),
		Entry("case 1-0", testData{
			k:        50,
			prices:   []int32{1, 12, 5, 111, 200, 1000, 10},
			expected: 4,
		}),
		Entry("case 1-1", testData{
			k:        15,
			prices:   []int32{3, 7, 2, 9, 4},
			expected: 3,
		}),
		Entry("case 2", testData{
			k:        1,
			prices:   []int32{1},
			expected: 1,
		}),
		Entry("case 3", testData{
			k:        0,
			prices:   []int32{1, 2, 3},
			expected: 0,
		}),
		Entry("case 3", testData{
			k:        3,
			prices:   []int32{1, 2, 1, 2},
			expected: 2,
		}),
		Entry("case 4", testData{
			k:        4,
			prices:   []int32{1, 2, 1, 2},
			expected: 3,
		}),
		Entry("case 5", testData{
			k:        100,
			prices:   []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 10,
		}),
	)
})
