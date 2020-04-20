package implementation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/apple-and-orange/problem", func() {
	It("문제를 풀었다.", func() {
		err := utils.PrintTest(func() {
			countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
		}, []string{
			"1",
			"1",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
