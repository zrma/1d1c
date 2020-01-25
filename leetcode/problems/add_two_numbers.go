package problems

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) traversal() string {
	res := strconv.Itoa(getVal(l))
	if getNext(l) == nil {
		return res
	}

	return getNext(l).traversal() + res
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func getNext(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func add(l *ListNode, val int) {
	prev := l
	tail := getNext(l)
	for tail != nil {
		prev = tail
		tail = getNext(tail)
	}
	prev.Next = &ListNode{Val: val}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	res := &ListNode{}
	for {
		val := getVal(l1) + getVal(l2) + carry
		if val >= 10 {
			val -= 10
			carry = 1
		} else {
			carry = 0
		}
		add(res, val)

		if getNext(l1) == nil && getNext(l2) == nil {
			break
		}
		l1, l2 = getNext(l1), getNext(l2)
	}
	if carry > 0 {
		add(res, carry)
	}
	return getNext(res)
}
