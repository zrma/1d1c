package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-loops/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			loop(2)
		}, []string{
			"2 x 1 = 2",
			"2 x 2 = 4",
			"2 x 3 = 6",
			"2 x 4 = 8",
			"2 x 5 = 10",
			"2 x 6 = 12",
			"2 x 7 = 14",
			"2 x 8 = 16",
			"2 x 9 = 18",
			"2 x 10 = 20",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
