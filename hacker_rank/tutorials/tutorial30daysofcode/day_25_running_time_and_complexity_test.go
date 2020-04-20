package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem", func() {
	It("문제를 풀었다", func() {
		Expect(isPrime(1)).Should(BeFalse())
		Expect(isPrime(2)).Should(BeTrue())
		Expect(isPrime(3)).Should(BeTrue())
		Expect(isPrime(4)).Should(BeFalse())
		Expect(isPrime(5)).Should(BeTrue())
		Expect(isPrime(7)).Should(BeTrue())
		Expect(isPrime(10)).Should(BeFalse())
		Expect(isPrime(12)).Should(BeFalse())
		Expect(isPrime(79)).Should(BeTrue())
		Expect(isPrime(131)).Should(BeTrue())
		Expect(isPrime(169)).Should(BeFalse())
		Expect(isPrime(173)).Should(BeTrue())
		Expect(isPrime(177)).Should(BeFalse())
	})
})
