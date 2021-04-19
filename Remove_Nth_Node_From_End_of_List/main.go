package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptrs := []*ListNode{head}
	nxt := head.Next
	for nxt != nil {
		ptrs = append(ptrs, nxt)
		nxt = nxt.Next
	}
	if len(ptrs) < n+1 {
		head = head.Next
	} else {
		ptrs[len(ptrs)-n-1].Next = ptrs[len(ptrs)-n].Next
	}
	return head
}
