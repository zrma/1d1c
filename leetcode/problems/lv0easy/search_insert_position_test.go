package lv0easy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/search-insert-position/", func() {
	type testData struct {
		nums     []int
		target   int
		expected int
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := searchInsert(data.nums, data.target)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{[]int{1, 3, 5, 6}, 5, 2}),
		Entry("1", testData{[]int{1, 3, 5, 6}, 2, 1}),
		Entry("2", testData{[]int{1, 3, 5, 6}, 7, 4}),
		Entry("3", testData{[]int{1, 3, 5, 6}, 0, 0}),
	)
})
