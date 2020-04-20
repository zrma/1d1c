package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-2d-arrays/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			arr := [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			}
			hourGlassSum(arr)
		}, []string{
			"19",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
