package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zrma/going/utils"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-sorting/problem", func() {
	It("문제를 풀었다", func() {
		err := utils.PrintTest(func() {
			sorting([]int32{1, 2, 3})
			sorting([]int32{3, 2, 1})
		}, []string{
			"Array is sorted in 0 swaps.",
			"First Element: 1",
			"Last Element: 3",
			"Array is sorted in 3 swaps.",
			"First Element: 1",
			"Last Element: 3",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})
})
