package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			stringToInteger("3")
			stringToInteger("za")
		}, []string{
			"3",
			"Bad String",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
