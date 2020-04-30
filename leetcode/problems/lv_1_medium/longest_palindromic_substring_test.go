package lv_1_medium

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/longest-palindromic-substring/", func() {
	type testData struct {
		input    string
		expected string
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := longestPalindrome(data.input)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{"babad", "bab"}),
		Entry("1", testData{"cbbd", "bb"}),
		Entry("2", testData{"cbbc", "cbbc"}),
		Entry("3", testData{"abcd", "a"}),
		Entry("4", testData{"", ""}),
		Entry("5", testData{"a", "a"}),
		Entry("6", testData{"abcdefedcba", "abcdefedcba"}),
		Entry("7-0", testData{"aabcdefedcba", "abcdefedcba"}),
		Entry("7-1", testData{"abcdefedcbaa", "abcdefedcba"}),
		Entry("7-2", testData{"aaabcdefedcba", "abcdefedcba"}),
		Entry("7-3", testData{"abcdefedcbaaa", "abcdefedcba"}),
		Entry("8", testData{"bb", "bb"}),
		Entry("9", testData{"aba", "aba"}),
		Entry("10", testData{"abb", "bb"}),
		Entry("11", testData{"aab", "aa"}),
		Entry("12", testData{"abbb", "bbb"}),
		Entry("13", testData{"aabbb", "bbb"}),
	)
})
