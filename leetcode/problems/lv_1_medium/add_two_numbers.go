package lv_1_medium

import (
	. "github.com/zrma/1d1c/leetcode/problems/common"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	res := &ListNode{}
	cur := res
	for {
		val := l1.GetVal() + l2.GetVal() + carry
		if val >= 10 {
			val -= 10
			carry = 1
		} else {
			carry = 0
		}

		cur.Next = &ListNode{Val: val}
		cur = cur.Next

		if l1.GetNext() == nil && l2.GetNext() == nil {
			break
		}
		l1, l2 = l1.GetNext(), l2.GetNext()
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return res.GetNext()
}
