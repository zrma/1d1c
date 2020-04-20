package strings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/common-child/problem", func() {
	It("문제를 풀었다", func() {
		actual := commonChild("ABCD", "ABDC")
		Expect(actual).Should(BeNumerically("==", 3))
		actual = commonChild("HARRY", "SALLY")
		Expect(actual).Should(BeNumerically("==", 2))
		actual = commonChild("AA", "BB")
		Expect(actual).Should(BeNumerically("==", 0))
		actual = commonChild("SHINCHAN", "NOHARAAA")
		Expect(actual).Should(BeNumerically("==", 3))
		actual = commonChild("ABCDEF", "FBDAMN")
		Expect(actual).Should(BeNumerically("==", 2))
	})
})
