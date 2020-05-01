package lv_1_medium

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/container-with-most-water/", func() {
	type testData struct {
		arr      []int
		expected int
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		actual := maxArea(data.arr)
		Expect(actual).Should(Equal(data.expected))
	},
		Entry("한 자리 올림", testData{
			arr:      []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		}),
	)
})
