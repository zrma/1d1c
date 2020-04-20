package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-binary-numbers/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			binaryNumbers(5)
			binaryNumbers(13)
		}, []string{
			"1",
			"2",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
