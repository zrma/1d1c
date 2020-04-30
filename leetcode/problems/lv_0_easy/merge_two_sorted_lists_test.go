package lv_0_easy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/zrma/1d1c/leetcode/problems/common"
)

var _ = Describe("https://leetcode.com/problems/merge-two-sorted-lists/", func() {
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
		actual := mergeTwoLists(l1, l2)
		Expect(actual.Traversal()).Should(Equal(expected.Traversal()))
	},
		Entry("정상 동작", testData{
			l1:       []int{1, 2, 4},
			l2:       []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		}),
		Entry("둘 다 nil", testData{
			l1:       []int(nil),
			l2:       []int(nil),
			expected: []int(nil),
		}),
		Entry("앞 nil", testData{
			l1:       []int(nil),
			l2:       []int{1, 3, 4},
			expected: []int{1, 3, 4},
		}),
		Entry("뒤 nil", testData{
			l1:       []int{1, 2, 4},
			l2:       []int(nil),
			expected: []int{1, 2, 4},
		}),
	)
})