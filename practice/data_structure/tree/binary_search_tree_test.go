package tree

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("BST 자료 구조가 삽입, 조회를 잘 한다.", func() {
	var bst *binarySearchTree
	BeforeEach(func() {
		bst = &binarySearchTree{}

		bst.insert(9)
		bst.insert(5)
		bst.insert(3)
		bst.insert(7)
		bst.insert(1)

		var result []int
		bst.traverse(func(data int) {
			result = append(result, data)
		})

		Expect(result).Should(Equal([]int{1, 3, 5, 7, 9}))
	})

	It("정확한 위치에 잘 들어 간다.", func() {
		//           9
		//       5
		//    3    7
		// 1
		Expect(bst.top.data).Should(Equal(9), "9")

		Expect(bst.top.left).ShouldNot(BeNil(), "5")
		Expect(bst.top.left.data).Should(Equal(5), "5")

		Expect(bst.top.right).Should(BeNil())

		Expect(bst.top.left.left).ShouldNot(BeNil(), "3")
		Expect(bst.top.left.left.data).Should(Equal(3), "3")

		Expect(bst.top.left.right).ShouldNot(BeNil(), "7")
		Expect(bst.top.left.right.data).Should(Equal(7), "7")

		Expect(bst.top.left.left.left).ShouldNot(BeNil(), "1")
		Expect(bst.top.left.left.left.data).Should(Equal(1), "1")

		Expect(bst.top.left.left.right).Should(BeNil(), "nil")
	})

	It("중복 입력 해도 한 번만 들어 간다.", func() {
		bst.insert(3)
		bst.insert(1)
		bst.insert(2)
		bst.insert(1)
		bst.insert(2)
		bst.insert(3)
		bst.insert(3)

		var result []int
		bst.traverse(func(data int) {
			result = append(result, data)
		})

		Expect(result).Should(Equal([]int{1, 2, 3, 5, 7, 9}))
	})

	Context("exist 함수 검증", func() {
		DescribeTable("원소가 있는지 잘 확인한다.", func(target int, expected bool) {
			Expect(bst.exist(target)).Should(Equal(expected))
		},
			Entry("1", 1, true),
			Entry("2", 2, false),
			Entry("3", 3, true),
			Entry("4", 4, false),
		)

		It("비어 있는 BST 자료구조에서 panic 하지 않음", func(done Done) {
			defer close(done)

			b := &binarySearchTree{}
			Expect(b.exist(100)).Should(BeFalse())
		})
	})
})
