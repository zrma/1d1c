package tutorial30daysofcode

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://www.hackerrank.com/challenges/30-binary-trees/problem", func() {
	It("문제를 풀었다", func() {
		var root *treeNode
		for _, num := range []int{3, 5, 4, 7, 2, 1} {
			root = insert(root, num)
		}
		diff := deep.Equal(levelOrder(root), []int{3, 2, 5, 1, 4, 7})
		Expect(diff).Should(BeNil())
	})
})
