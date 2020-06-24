package warmup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/solve-me-first/problem", func() {
	It("문제를 풀었다", func() {
		var a uint32 = 2
		var b uint32 = 3
		actual := solveMeFirst(a, b)
		Expect(actual).Should(BeNumerically("==", 5))
	})
})
