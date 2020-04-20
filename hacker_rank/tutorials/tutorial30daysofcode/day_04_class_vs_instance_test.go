package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-class-vs-instance/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			classAndInstance(-1)
			classAndInstance(10)
			classAndInstance(16)
			classAndInstance(18)
		}, []string{
			"Age is not valid, setting age to 0.",
			"You are young.",
			"You are young.",
			"",
			"You are young.",
			"You are a teenager.",
			"",
			"You are a teenager.",
			"You are old.",
			"",
			"You are old.",
			"You are old.",
			"",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
