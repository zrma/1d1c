package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-loops/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			letsReview("Hacker")
			letsReview("Rank")
		}, []string{
			"Hce akr",
			"Rn ak",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
