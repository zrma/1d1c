package strings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/string-construction/problem", func() {
	It("문제를 풀었다", func() {
		actual := stringConstruction("abcd")
		Expect(actual).Should(BeNumerically("==", 4))
		actual = stringConstruction("abab")
		Expect(actual).Should(BeNumerically("==", 2))
		actual = stringConstruction("abcabc")
		Expect(actual).Should(BeNumerically("==", 3))
	})
})
