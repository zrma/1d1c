package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListNode", func() {
	It("동작 검증", func() {
		node := ListNode{
			Val: 123,
			Next: &ListNode{
				Val: 456,
			},
		}

		By("GetVal 함수 검증", func() {
			actual := node.GetVal()
			Expect(actual).Should(Equal(123))
		})

		By("GetNext 함수 검증", func() {
			next := node.GetNext()
			Expect(next).ShouldNot(BeNil())

			actual := next.GetVal()
			Expect(actual).Should(Equal(456))
		})

		By("Traversal 함수 검증", func() {
			actual := node.Traversal()
			Expect(actual).Should(Equal("456123"))
		})
	})
})
