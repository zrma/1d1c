package lv1medium

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/longest-substring-without-repeating-characters/", func() {
	type testData struct {
		input    string
		expected int
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := lengthOfLongestSubstring(data.input)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{"abcabcbb", 3}),
		Entry("1", testData{"bbbbb", 1}),
		Entry("2", testData{"pwwkew", 3}),
		Entry("3", testData{"", 0}),
		Entry("4", testData{"au", 2}),
		Entry("5", testData{"abba", 2}),
	)
})
