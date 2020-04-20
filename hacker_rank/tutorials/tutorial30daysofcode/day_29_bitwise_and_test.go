package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-bitwise-and/problem", func() {
	type testData struct {
		n, k     int32
		expected int32
	}
	DescribeTable("문제를 풀었다.",
		func(data testData) {
			actual := bitwiseAND(data.n, data.k)
			Expect(actual).Should(BeNumerically("==", data.expected))
		},
		Entry("test_0", testData{5, 2, 1}),
		Entry("test_1", testData{8, 5, 4}),
		Entry("test_2", testData{2, 2, 0}),
		Entry("test_2", testData{955, 236, 235}),
	)
})
