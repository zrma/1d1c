package strings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem", func() {
	It("문제를 풀었다", func() {
		Expect(isValid("aabbcd")).Should(Equal("NO"))
		Expect(isValid("aabbccddeefghi")).Should(Equal("NO"))
		Expect(isValid("abbccc")).Should(Equal("NO"))
		Expect(isValid("aaaabbcc")).Should(Equal("NO"))
		Expect(isValid("aaaaabc")).Should(Equal("NO"))
		Expect(isValid("aabbc")).Should(Equal("YES"))
		Expect(isValid("aabbcc")).Should(Equal("YES"))
		Expect(isValid("abcdefghhgfedecba")).Should(Equal("YES"))
	})
})
