package strings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/morgan-and-a-string/problem", func() {
	type testData struct {
		jack, daniel string
		expected     string
	}

	//noinspection SpellCheckingInspection
	DescribeTable("문제를 풀었다",
		func(data testData) {
			actual := morganAndString(data.jack, data.daniel)
			Expect(actual).Should(Equal(data.expected))
		},
		Entry("test_0", testData{"ACA", "BCF", "ABCACF"}),
		Entry("test_1", testData{"JACK", "DANIEL", "DAJACKNIEL"}),
		Entry("test_2", testData{"ABACABA", "ABACABA", "AABABACABACABA"}),
		Entry("test_3", testData{"AAAAB", "AAAAC", "AAAAAAAABC"}),
		Entry("test_3", testData{"AAAAC", "AAAAB", "AAAAAAAABC"}),
	)
})
