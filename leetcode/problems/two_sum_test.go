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

var _ = Describe("https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/", func() {
	It("문제를 풀었다.", func() {
		actual := twoSum2InputArrayIsSorted([]int{2, 7, 11, 15}, 9)
		Expect(actual).Should(Equal([]int{1, 2}))
	})
})
