package strings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/reduced-string/problem", func() {
	It("문제를 풀었다", func() {
		actual := superReducedString("aaabccddd")
		Expect(actual).Should(Equal("abd"))
		actual = superReducedString("aa")
		Expect(actual).Should(Equal("Empty String"))
		actual = superReducedString("baab")
		Expect(actual).Should(Equal("Empty String"))
		actual = superReducedString("abab")
		Expect(actual).Should(Equal("abab"))
		actual = superReducedString("abccba")
		Expect(actual).Should(Equal("Empty String"))
		actual = superReducedString("")
		Expect(actual).Should(Equal("Empty String"))
	})
})
