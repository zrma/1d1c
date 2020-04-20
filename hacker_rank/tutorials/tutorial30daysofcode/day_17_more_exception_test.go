package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-more-exceptions/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			moreException(3, 5)
			moreException(2, 4)
			moreException(-1, -2)
			moreException(-1, 3)
		}, []string{
			"243",
			"16",
			"n and p should be non-negative",
			"n and p should be non-negative",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
