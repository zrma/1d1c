package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/zigzag-conversion/", func() {
	It("문제를 풀었다.", func() {
		//noinspection SpellCheckingInspection
		const inputString = "PAYPALISHIRING"

		By("case 1", func() {
			//noinspection SpellCheckingInspection
			const expected = "PAHNAPLSIIGYIR"

			actual := convert(inputString, 3)
			Expect(actual).Should(Equal(expected))
		})

		By("case 2", func() {
			//noinspection SpellCheckingInspection
			const expected = "PINALSIGYAHRPI"

			actual := convert(inputString, 4)
			Expect(actual).Should(Equal(expected))
		})
	})

	It("it should not be panic", func(done Done) {
		defer close(done)

		//noinspection SpellCheckingInspection
		const expected = "AB"

		actual := convert("AB", 1)
		Expect(actual).Should(Equal(expected))
	})
})
