package lv_1_medium

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/zrma/1d1c/leetcode/problems/common"
)

var _ = Describe("https://leetcode.com/problems/add-two-numbers/", func() {
	build := func(arr []int) *ListNode {
		l := &ListNode{}
		cur := l
		for _, n := range arr {
			cur.Next = &ListNode{Val: n}
			cur = cur.Next
		}
		return l.GetNext()
	}

	type testData struct {
		l1, l2   []int
		expected []int
	}

	DescribeTable("문제를 풀었다.", func(data testData) {
		l1 := build(data.l1)
		l2 := build(data.l2)
		expected := build(data.expected)
		actual := addTwoNumbers(l1, l2)
		Expect(actual.Traversal()).Should(Equal(expected.Traversal()))
	},
		Entry("정상 동작", testData{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		}),
		Entry("한 자리 올림", testData{
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{0, 1},
		}),
		Entry("예외 처리.", testData{
			l1:       nil,
			l2:       []int{5, 6, 4},
			expected: []int{5, 6, 4},
		}),
	)
})
