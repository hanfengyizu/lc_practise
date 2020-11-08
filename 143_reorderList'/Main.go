package main

import (
	"fmt"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	reorderList1(head)
	fmt.Printf("%+v\n", *head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	l := head.Next
	if l == nil || l.Next == nil {
		return
	}
	h := head
	nums := make([]int, 0)
	length := 0
	for ; head.Next != nil; {
		nums = append(nums, head.Val)
		l = head.Next
		head = l
		length++
	}
	nums = append(nums, head.Val)
	head = h
	index := 0
	length++
	preIndex := 0
	postIndex := 0
	for ; head != nil; index++ {
		preIndex = index
		postIndex = length - index - 1
		head.Val = nums[preIndex]
		l = head.Next
		head = l
		if head == nil {
			continue
		}
		head.Val = nums[postIndex]
		l = head.Next
		head = l
	}
	if head == nil {
		return
	}

	head.Val = nums[index]
}

func reorderList1(head *ListNode) {
	if head == nil {
		return
	}
	l := head.Next
	if l == nil || l.Next == nil {
		return
	}

	popLastNode := func(h *ListNode) *ListNode {
		if h == nil {
			return nil
		}
		temp := h
		for ; h.Next != nil && h.Next.Next != nil; {
			temp = h.Next
			h = temp
		}
		tmp := temp.Next
		temp.Next = nil
		return tmp
	}
	secondNode := &ListNode{
		Val:  0,
		Next: nil,
	}
	for ; head != nil; {
		secondNode = head.Next
		if secondNode == nil {
			break
		}
		last := popLastNode(secondNode)
		if last != nil {
			head.Next = last
			head = head.Next
		}
		head.Next = secondNode
		head = head.Next
	}
}
