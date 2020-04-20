package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-recursion/problem", func() {
	It("문제를 풀었다", func() {
		actual := factorial(3)
		Expect(actual).Should(BeNumerically("==", 6))

		actual = factorial(10)
		Expect(actual).Should(BeNumerically("==", 3628800))
	})
})
