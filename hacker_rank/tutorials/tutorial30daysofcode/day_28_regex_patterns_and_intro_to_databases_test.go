package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-regex-patterns/problem", func() {
	It("문제를 풀었다", func() {
		//noinspection SpellCheckingInspection
		err := utils.PrintTest(func() {
			filter([][]string{
				{"riya", "riya@gmail.com"},
				{"julia", "julia@julia.me"},
				{"julia", "sjulia@gmail.com"},
				{"julia", "julia@gmail.com"},
				{"samantha", "samantha@gmail.com"},
				{"tanya", "tanya@gmail.com"},
			})
		}, []string{
			"julia",
			"julia",
			"riya",
			"samantha",
			"tanya",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
