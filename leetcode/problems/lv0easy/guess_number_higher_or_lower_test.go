package lv0easy

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/guess-number-higher-or-lower/", func() {
	type testData struct {
		n        int
		expected int
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual, callCount := solveGuessNumber(data.n, data.expected)
			Expect(actual).Should(Equal(data.expected))

			// 3진 탐색, 1회 탐색에 2번씩 호출
			expectedCount := ((math.Log(float64(data.n)))/math.Log(3))*2 + 1
			Expect(callCount).Should(BeNumerically("<=", math.Ceil(expectedCount)))
		},
		Entry("0", testData{5, 4}),
		Entry("1-0", testData{10, 3}),
		Entry("1-1", testData{10, 6}),
		Entry("2-0", testData{100, 37}),
		Entry("2-1", testData{100, 38}),
		Entry("2-2", testData{100, 49}),
		Entry("2-3", testData{100, 50}),
		Entry("2-4", testData{100, 51}),
		Entry("3-0", testData{1024, 511}),
		Entry("3-1", testData{1024, 512}),
		Entry("3-2", testData{1024, 513}),
		Entry("3-3", testData{1024, 1024}),
	)
})
