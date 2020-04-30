package common

import "strconv"

// ListNode struct implements single linked list data structure.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traversal() string {
	res := strconv.Itoa(l.GetVal())
	if l.GetNext() == nil {
		return res
	}
	return l.GetNext().Traversal() + res
}

func (l *ListNode) GetVal() int {
	if l == nil {
		return 0
	}
	return l.Val
}

func (l *ListNode) GetNext() *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}
