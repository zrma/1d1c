package problems

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("https://leetcode.com/problems/add-two-numbers/", func() {
	build := func(arr []int) *ListNode {
		l := &ListNode{}
		cur := l
		for _, n := range arr {
			cur.Next = &ListNode{Val: n}
			cur = cur.Next
		}
		return getNext(l)
	}

	It("문제를 풀었다.", func() {
		l1 := build([]int{2, 4, 3})
		l2 := build([]int{5, 6, 4})
		expected := build([]int{7, 0, 8})
		actual := addTwoNumbers(l1, l2)
		Expect(actual.traversal()).Should(Equal(expected.traversal()))
	})

	It("한 자리 올림", func() {
		l1 := build([]int{5})
		l2 := build([]int{5})
		expected := build([]int{0, 1})
		actual := addTwoNumbers(l1, l2)
		Expect(actual.traversal()).Should(Equal(expected.traversal()))
	})

	It("예외 처리.", func() {
		l1 := build(nil)
		l2 := build([]int{5, 6, 4})
		expected := build([]int{5, 6, 4})
		actual := addTwoNumbers(l1, l2)
		Expect(actual.traversal()).Should(Equal(expected.traversal()))
	})
})
