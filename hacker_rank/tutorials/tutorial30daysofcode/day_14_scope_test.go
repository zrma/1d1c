package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-scope/problem", func() {
	It("문제를 풀었다", func() {
		difference := newDifference([]int{1, 2, 5})
		Expect(difference.computeDifference()).Should(BeNumerically("==", 4))
	})
})
