package implementation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/between-two-sets/problem", func() {
	It("문제를 풀었다.", func() {
		expected := getTotalX([]int32{2, 4}, []int32{16, 32, 96})
		Expect(expected).Should(BeNumerically("==", 3))

		expected = getTotalX([]int32{3, 4}, []int32{24, 48})
		Expect(expected).Should(BeNumerically("==", 2))
	})
})
