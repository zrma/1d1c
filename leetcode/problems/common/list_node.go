package common

import "strconv"

// ListNode struct implements single linked list data structure.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Traversal method concat all values of linked list as string
func (l *ListNode) Traversal() string {
	res := strconv.Itoa(l.GetVal())
	if l.GetNext() == nil {
		return res
	}
	return l.GetNext().Traversal() + res
}

// GetVal method return value of node
func (l *ListNode) GetVal() int {
	if l == nil {
		return 0
	}
	return l.Val
}

// GetNext method return pointer of next node
func (l *ListNode) GetNext() *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}
