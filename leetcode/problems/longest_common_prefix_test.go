package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/longest-common-prefix/", func() {
	type testData struct {
		input    []string
		expected string
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := longestCommonPrefix(data.input)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{[]string{"flower", "flow", "flight"}, "fl"}),
		Entry("1", testData{[]string{"dog", "racecar", "car"}, ""}),
		Entry("2", testData{[]string{"flower", "flower", "flower"}, "flower"}),
		Entry("3", testData{[]string{"flowers", "flower", "flower"}, "flower"}),
		Entry("4", testData{[]string{}, ""}),
		Entry("5", testData{nil, ""}),
		Entry("6", testData{[]string{"", "flower", "flower"}, ""}),
		Entry("7", testData{[]string{"c", "c"}, "c"}),
	)
})
