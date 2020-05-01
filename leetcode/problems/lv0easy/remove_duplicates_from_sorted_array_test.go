package lv0easy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/remove-duplicates-from-sorted-array/", func() {
	type testData struct {
		nums     []int
		expected []int
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := removeDuplicates(data.nums)
		length := len(data.expected)

		Expect(actual).Should(Equal(length))
		Expect(data.nums[:length]).Should(Equal(data.expected))
	},
		Entry("case 1", testData{
			nums:     []int{1, 1, 2},
			expected: []int{1, 2},
		}),
		Entry("case 2", testData{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		}),
		Entry("case 3", testData{
			nums:     []int{0, 0, 3, 3, 3, 3, 3, 4, 5, 5},
			expected: []int{0, 3, 4, 5},
		}),
		Entry("nil 방어", testData{
			nums:     nil,
			expected: nil,
		}),
	)
})
