package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-hello-world/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			helloWorld("Welcome to 30 Days of Code!")
		}, []string{
			"Hello, World.",
			"Welcome to 30 Days of Code!",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
