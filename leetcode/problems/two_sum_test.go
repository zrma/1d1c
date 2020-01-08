package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/two-sum/", func() {
	It("문제를 풀었다.", func() {
		actual := twoSum([]int{2, 7, 11, 15}, 9)
		Expect(actual).Should(Equal([]int{0, 1}))
	})
})
