package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := TreeNode{Val: 5}
	node.Left = &TreeNode{Val: 2}
	node.Right = &TreeNode{Val: 13}
	resNode := convertBST(&node)
	fmt.Printf("Res: %+v", *resNode)
}

func convertBST(root *TreeNode) *TreeNode {
	num2Map := make(map[int]int)
	walkTree(root, num2Map)
	modifyTree(root, num2Map)
	return root
}

func modifyTree(root *TreeNode, num2Map map[int]int) {
	if root == nil {
		return
	}
	(*root).Val = num2Map[(*root).Val]
	modifyTree(root.Left, num2Map)
	modifyTree(root.Right, num2Map)
}

func walkTree(root *TreeNode, num2Map map[int]int) {
	if root == nil {
		return
	}
	key := (*root).Val
	_, ok := num2Map[key]
	if !ok {
		num2Map[key] = key
		for k, v := range num2Map {
			if key > k {
				num2Map[k] = v + key
			}
			if key < k {
				num2Map[key] = k + num2Map[key]
			}
		}
	}
	walkTree(root.Left, num2Map)
	walkTree(root.Right, num2Map)
}
