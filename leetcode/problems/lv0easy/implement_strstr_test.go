package lv0easy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/implement-strstr/", func() {
	type testData struct {
		haystack string
		needle   string
		expected int
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := strStr(data.haystack, data.needle)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{"hello", "ll", 2}),
		Entry("1", testData{"aaaaa", "bba", -1}),
		Entry("2", testData{"aaaaa", "", 0}),
		Entry("3", testData{"", "", 0}),
		Entry("4", testData{"", "abc", -1}),
	)
})
