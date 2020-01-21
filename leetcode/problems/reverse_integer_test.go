package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/reverse-integer/", func() {
	type testData struct {
		input, expected int
	}

	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := reverse(data.input)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{123, 321}),
		Entry("1", testData{-123, -321}),
		Entry("2", testData{120, 21}),
		Entry("3", testData{1534236469, 0}),
	)
})
