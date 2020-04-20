package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-conditional-statements/problem", func() {
	It("문제를 풀었다", func() {
		actual := cond(3)
		Expect(actual).Should(Equal(weird))

		actual = cond(24)
		Expect(actual).Should(Equal(notWeird))
	})
})
