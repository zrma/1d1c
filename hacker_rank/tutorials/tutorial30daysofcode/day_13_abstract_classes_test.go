package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

//noinspection SpellCheckingInspection
var _ = Describe("https://www.hackerrank.com/challenges/30-abstract-classes/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			abstractClasses("The Alchemist", "Paulo Coelho", 248)
		}, []string{
			"Title: The Alchemist",
			"Author: Paulo Coelho",
			"Price: 248",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
