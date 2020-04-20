package tutorial30daysofcode

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-binary-search-trees/problem", func() {
	It("문제를 풀었다", func() {
		var root *treeNode
		for _, num := range []int{3, 5, 2, 1, 4, 6, 7} {
			root = insert(root, num)
		}
		Expect(getHeight(root)).Should(BeNumerically("==", 3))
	})
})
