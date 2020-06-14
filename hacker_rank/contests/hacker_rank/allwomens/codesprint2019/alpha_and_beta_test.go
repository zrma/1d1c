package codesprint2019

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/contests/hackerrank-all-womens-codesprint-2019/challenges/alpha-and-beta", func() {
	type testData struct {
		pile     []int32
		expected int32
	}
	DescribeTable("문제를 풀었다", func(data testData) {
		Expect(alphaBeta(data.pile)).Should(Equal(data.expected))
	},
		Entry("0", testData{[]int32{1, 2, 3, 3, 2, 1}, 2}),
		Entry("1", testData{[]int32{1, 2, 3, 4, 5}, 4}),
		Entry("2", testData{[]int32{1, 2, 3, 2, 1}, 2}),
		Entry("3", testData{[]int32{1, 2, 4, 3, 2, 1}, 3}),
	)
})
