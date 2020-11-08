package main

import (
	"unsafe"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func hasCycle_v2(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeInfoMap := make(map[*ListNode]bool)
	for ; head.Next != nil; {
		if _, ok := nodeInfoMap[head]; ok {
			return true
		}
		nodeInfoMap[head] = true
		head = head.Next
	}
	return false
}
func hasCycle_v3(head *ListNode) bool {
	if head == nil {
		return false
	}
	nodeInfoMap := make(map[*ListNode]struct{})
	for ; head.Next != nil; {
		if _, ok := nodeInfoMap[head]; ok {
			return true
		}
		nodeInfoMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

func hasCycle_v1(head *ListNode) bool {
	nodeInfoMap := make(map[unsafe.Pointer]bool)
	if head == nil {
		return false
	}
	for ; head.Next != nil; {
		if _, ok := nodeInfoMap[unsafe.Pointer(head)]; ok {
			return true
		}
		nodeInfoMap[unsafe.Pointer(head)] = true
		head = head.Next
	}
	return false
}
