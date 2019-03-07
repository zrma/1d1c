package tutorial_30_days_of_code

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/1d1c/hacker_rank/common/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-2d-arrays/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() error {
			arr := [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			}
			hourGlassSum(arr)
			return nil
		}, []string{
			"19",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
