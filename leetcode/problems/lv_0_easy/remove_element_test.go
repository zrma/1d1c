package lv_0_easy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/remove-element/", func() {
	type testData struct {
		nums     []int
		val      int
		expected []int
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := removeElement(data.nums, data.val)
		length := len(data.expected)

		Expect(actual).Should(Equal(length))
		Expect(data.nums[:length]).Should(Equal(data.expected))
	},
		Entry("case 1", testData{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: []int{2, 2},
		}),
		Entry("case 2", testData{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: []int{0, 1, 3, 0, 4},
		}),
		Entry("case 3", testData{
			nums:     []int{0, 0, 3, 3, 3, 3, 3, 4, 5, 5},
			val:      3,
			expected: []int{0, 0, 4, 5, 5},
		}),
		Entry("nil 방어", testData{
			nums:     nil,
			val:      3,
			expected: nil,
		}),
	)
})
