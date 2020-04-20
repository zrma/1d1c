package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-interfaces/problem", func() {
	It("문제를 풀었다", func() {
		c := &calculator{}

		Expect(divisorSum(c, 1)).Should(BeNumerically("==", 1))
		Expect(divisorSum(c, 6)).Should(BeNumerically("==", 12))
		Expect(divisorSum(c, 20)).Should(BeNumerically("==", 42))
	})
})
