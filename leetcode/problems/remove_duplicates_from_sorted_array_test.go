package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/remove-duplicates-from-sorted-array/", func() {
	type testData struct {
		nums         []int
		expectedNums []int
		expectedLen  int
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := removeDuplicates(data.nums)
		Expect(actual).Should(Equal(data.expectedLen))
		Expect(data.nums[:data.expectedLen]).Should(Equal(data.expectedNums))
	},
		Entry("case 1", testData{
			nums:         []int{1, 1, 2},
			expectedNums: []int{1, 2},
			expectedLen:  2,
		}),
		Entry("case 2", testData{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
			expectedLen:  5,
		}),
		Entry("case 3", testData{
			nums:         []int{0, 0, 3, 3, 3, 3, 3, 4, 5, 5},
			expectedNums: []int{0, 3, 4, 5},
			expectedLen:  4,
		}),
	)
})
