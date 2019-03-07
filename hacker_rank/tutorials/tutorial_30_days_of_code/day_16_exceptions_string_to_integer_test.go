package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			stringToInteger("3")
			stringToInteger("za")
			return nil
		}, []string{
			"3",
			"Bad String",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
