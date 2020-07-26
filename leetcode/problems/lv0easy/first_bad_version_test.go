package lv0easy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/first-bad-version/", func() {
	type testData struct {
		n        int
		expected int
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := solveFirstBadVersion(data.n, data.expected)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("0", testData{5, 4}),
		Entry("1", testData{10, 3}),
		Entry("2-0", testData{100, 37}),
		Entry("2-1", testData{100, 38}),
		Entry("2-2", testData{100, 49}),
		Entry("2-3", testData{100, 50}),
		Entry("2-4", testData{100, 51}),
		Entry("3-0", testData{1024, 511}),
		Entry("3-1", testData{1024, 512}),
		Entry("3-2", testData{1024, 513}),
	)
})
