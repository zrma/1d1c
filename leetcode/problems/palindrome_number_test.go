package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/palindrome-number/", func() {
	type testData struct {
		input    int
		expected bool
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := isPalindrome(data.input)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{121, true}),
		Entry("1", testData{-121, false}),
		Entry("2", testData{10, false}),
		Entry("3", testData{1234, false}),
		Entry("4", testData{1234321, true}),
	)
})
