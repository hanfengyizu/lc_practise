package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil && head.Val == val {
		return nil
	}

	if head.Next == nil && head.Val != val {
		return head
	}

	pre := &ListNode{
		Val: 0, Next: head,
	}
	H := pre

	for pre.Next.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}

	if pre.Next.Val == val {
		pre.Next = nil
	}
	return H.Next
}
